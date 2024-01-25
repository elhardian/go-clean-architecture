package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv                     string `mapstructure:"APP_ENV"`
	AppTz                      string `mapstructure:"APP_TZ"`
	AppIsDev                   bool
	AppPort                    string `mapstructure:"APP_PORT"`
	DatabaseDriver             string `mapstructure:"DATABASE_DRIVER"`
	DatabaseConnectionString   string `mapstructure:"DATABASE_CONNECTION_STRING"`
	DatabaseMaxOpenConnections int    `mapstructure:"DATABASE_MAX_OPEN_CONNECTIONS"`
	DatabaseMaxIdleConnections int    `mapstructure:"DATABASE_MAX_IDLE_CONNECTIONS"`
	HttpTimeout                int    `mapstructure:"HTTP_TIMEOUT"`
}

func NewConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("libs/config")
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Failed To Read Config")
		return nil, err
	}

	cfg := &Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed To Unmarshal Config")
		return nil, err
	}

	cfg.AppIsDev = env == "local" || env == "dev"

	return cfg, nil
}
