package main

import (
	"github.com/isteshkov/highload-social-network/internal/app/stores"
	"log"
	"os"

	"github.com/isteshkov/highload-social-network/configs"
	"github.com/isteshkov/highload-social-network/internal/app/business"
	application "github.com/isteshkov/highload-social-network/internal/app/server"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database/migrations"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/services"
	external "github.com/isteshkov/highload-social-network/third_party/metrics"
)

// @title Socnet Manage API
// @version 1.0
// @securityDefinitions.apikey AccessToken
// @in header
// @name X-Access-Token

func main() {
	file := ""
	args := os.Args[1:]
	if len(args) > 0 {
		file = args[0]
	}

	cfg, err := configs.LoadConfig(file)
	if err != nil {
		log.Println("Error loading config, using default")
		panic(err)
	}

	if len(args) > 1 && args[1] == "migrate" {
		err = migrations.MigrateUp(cfg.DatabaseDSN)
		if err != nil {
			panic(err)
		}
		return
	}

	logger, err := logging.NewLogger(&logging.Config{
		LogLvl:    cfg.LogLevel,
		Version:   cfg.Version,
		CommitSha: cfg.CommitSha,
		AppName:   cfg.AppName,
	})
	if err != nil {
		panic(err)
	}

	db, err := database.GetDatabase(database.Config{ConnectionDSN: cfg.DatabaseDSN}, logger)
	if err != nil {
		panic(err)
	}

	authStore := stores.NewSessionsStore(db, logger)
	authService := business.NewAuthService(logger, authStore)

	usersStore := stores.NewUsersStore(db, logger)
	usersService := business.NewUsersService(logger, usersStore)

	profilesStore := stores.NewProfilesStore(db, logger)
	profilesService := business.NewProfilesService(logger, profilesStore)

	friendshipsStore := stores.NewFriendshipsStore(db, logger)
	friendshipsService := business.NewFriendshipsService(logger, friendshipsStore)

	businessServices := services.NewServices(
		logger, authService, usersService, profilesService, friendshipsService)

	metrics, err := external.NewMetrics()
	if err != nil {
		panic(err)
	}

	app := application.BuildApplication(logger, businessServices, metrics, application.Config{
		ProfilingApiPort: cfg.ProfilingAPIPort,
		PublicApiPort:    cfg.PublicAPIPort,
		MetricApiPort:    cfg.MetricAPIPort,
		TimeOutSecond:    cfg.Timeout,
	})

	app.ListenAndServe()
}
