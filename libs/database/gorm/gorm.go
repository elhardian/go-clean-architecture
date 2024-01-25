package gormDatabase

import (
	"fmt"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/elhardian/go-clean-architecture/libs/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm interface {
	Connect() (*gorm.DB, error)
}

type Options struct {
	driver           string
	connectionString string
	maxOpen          int
	maxIdle          int
}

const (
	MySQL      = "mysql"
	PostgreSQL = "postgres"
)

func NewGorm(cfg *config.Config) Gorm {
	opt := new(Options)
	opt.driver = cfg.DatabaseDriver
	opt.connectionString = cfg.DatabaseConnectionString
	opt.maxOpen = cfg.DatabaseMaxOpenConnections
	opt.maxIdle = cfg.DatabaseMaxIdleConnections

	return opt
}

func (o *Options) Connect() (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	switch o.driver {
	case MySQL:
		db, err = gorm.Open(mysql.Open(o.connectionString), &gorm.Config{})
	case PostgreSQL:
		db, err = gorm.Open(postgres.Open(o.connectionString), &gorm.Config{})
	default:
		log.Error().Msgf("Unsupported database driver: %s", o.driver)
		return nil, fmt.Errorf("unsupported database driver: %s", o.driver)
	}

	if err != nil {
		log.Error().Err(err).Msg("Failed To Connect Gorm")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to connect to %s database", o.driver)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(o.maxOpen)
	sqlDB.SetMaxIdleConns(o.maxIdle)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := backoff.Retry(func() error {
		if err := sqlDB.Ping(); err != nil {
			log.Error().Err(err).Msg("Failed To Ping Gorm Database")
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		log.Error().Err(err).Msg("Failed To Retry Ping Gorm Database")
		return nil, err
	}

	return db, nil
}
