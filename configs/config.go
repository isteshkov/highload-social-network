package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	DefaultServiceName   = "brute-force-protection-server"
	DefaultProfilingPort = ":6060"
	DefaultMetricsPort   = ":8001"
)

type Config struct {
	AppName   string `env:"APP_NAME"`
	Version   string `env:"VERSION"`
	CommitSha string `env:"COMMIT_SHA"`

	ProfilingAPIPort string `env:"PROFILING_API_PORT"`
	PublicAPIPort    string `env:"PUBLIC_API_PORT,required"`
	MetricAPIPort    string `env:"METRIC_API_PORT"`
	Timeout          int    `env:"TIME_OUT,required"`

	LogLevel    string `env:"LOG_LEVEL"`
	DatabaseDSN string `env:"DATABASE_DSN,required"`
}

func LoadConfig(fileName string) (cfg *Config, err error) {
	if len(fileName) > 0 {
		_ = loadEnvFromFile(fileName)
	}
	cfg = &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return
	}

	fillDefault(cfg)

	return
}

func loadEnvFromFile(filename string) error {
	if _, err := os.Stat(filename); err == nil {
		if err := godotenv.Load(filename); err != nil {
			return fmt.Errorf("error loading file: %s, %w", filename, err)
		}
		log.Printf("Config file %s is using.\n", filename)
	} else {
		log.Printf("Missed config %s. Using env only.\n", filename)
	}

	return nil
}

func fillDefault(cfg *Config) {
	if cfg.AppName == "" {
		cfg.AppName = DefaultServiceName
	}
	if cfg.ProfilingAPIPort == "" {
		cfg.ProfilingAPIPort = DefaultProfilingPort
	}
	if cfg.MetricAPIPort == "" {
		cfg.MetricAPIPort = DefaultMetricsPort
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = logging.LevelDebug
	}
}
