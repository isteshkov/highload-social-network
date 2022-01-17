package migrations

import (
	"fmt"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"

	"github.com/gobuffalo/packr"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
)

func MigrateUp(dbUrl string) (err error) {
	l, err := logging.NewLogger(&logging.Config{LogLvl: logging.LevelError})
	if err != nil {
		return
	}

	db, err := database.GetDatabase(database.Config{
		ConnectionDSN: dbUrl,
	}, l)
	if err != nil {
		return
	}

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./scripts"),
	}

	n, err := migrate.Exec(db.Client(), "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return nil
}

func MigrateDown(dbUrl string) (err error) {
	l, err := logging.NewLogger(&logging.Config{LogLvl: logging.LevelError})
	if err != nil {
		return
	}

	db, err := database.GetDatabase(database.Config{
		ConnectionDSN: dbUrl,
	}, l)
	if err != nil {
		return
	}

	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./scripts"),
	}

	n, err := migrate.Exec(db.Client(), "postgres", migrations, migrate.Down)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Downgraded %d migrations!\n", n)

	return nil
}
