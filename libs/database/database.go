// package/database/database.go

package database

import (
	"database/sql"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/elhardian/go-clean-architecture/libs/config"
	"github.com/rs/zerolog/log"
)

type Database interface {
	Connect() (*sql.DB, error)
}

type Options struct {
	Driver           string
	ConnectionString string
	MaxOpen          int
	MaxIdle          int
}

func NewDatabase(cfg *config.Config) Database {
	opt := &Options{
		Driver:           cfg.DatabaseDriver,
		ConnectionString: cfg.DatabaseConnectionString,
		MaxOpen:          cfg.DatabaseMaxOpenConnections,
		MaxIdle:          cfg.DatabaseMaxIdleConnections,
	}
	return opt
}

func (o *Options) Connect() (*sql.DB, error) {
	db, err := sql.Open(o.Driver, o.ConnectionString)
	if err != nil {
		log.Error().Err(err).Msgf("Failed To Connect %s", o.Driver)
		return nil, err
	}

	db.SetMaxOpenConns(o.MaxOpen)
	db.SetMaxIdleConns(o.MaxIdle)
	db.SetConnMaxLifetime(time.Hour)

	if err := backoff.Retry(func() error {
		if err := db.Ping(); err != nil {
			log.Error().Err(err).Msgf("Failed To Ping %s", o.Driver)
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		log.Error().Err(err).Msgf("Failed To Retry Ping %s", o.Driver)
		return nil, err
	}

	return db, nil
}
