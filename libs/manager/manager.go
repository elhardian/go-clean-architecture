// package/manager/manager.go

package manager

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/elhardian/go-clean-architecture/libs/config"
	"github.com/elhardian/go-clean-architecture/libs/database"
	gormDatabase "github.com/elhardian/go-clean-architecture/libs/database/gorm"
	httpClient "github.com/elhardian/go-clean-architecture/libs/http"
	"github.com/elhardian/go-clean-architecture/libs/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Manager interface {
	GetConfig() *config.Config
	GetServer() *server.Server
	GetDatabase() *sql.DB
	GetGorm() *gorm.DB
	GetHttp() httpClient.Http
}

type manager struct {
	config     *config.Config
	server     *server.Server
	db         *sql.DB
	dbGorm     *gorm.DB
	httpClient httpClient.Http
}

func NewInit() (Manager, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error().Err(err).Msg("Failed to Initialize Configuration")
		return nil, err
	}

	srv := server.NewServer(cfg)

	// Create the appropriate database connection based on the provided type
	var db *sql.DB
	switch cfg.DatabaseDriver {
	case "mysql":
		db, err = database.NewMySQL(cfg).Connect()
		if err != nil {
			log.Error().Err(err).Msg("Failed to Initialize Database MySQL")
			return nil, err
		}
	case "postgres":
		db, err = database.NewPostgreSQL(cfg).Connect()
		if err != nil {
			log.Error().Err(err).Msg("Failed to Initialize Database PostgreSQL")
			return nil, err
		}
	default:
		log.Error().Msgf("Unsupported database type: %s", cfg.DatabaseDriver)
		return nil, fmt.Errorf("unsupported database type: %s", cfg.DatabaseDriver)
	}

	// Create Gorm database connection
	dbGorm, err := gormDatabase.NewGorm(cfg).Connect()
	if err != nil {
		log.Error().Err(err).Msg("Failed to Initialize Database Gorm")
		return nil, err
	}

	clHttp := httpClient.NewHttp(cfg)
	clHttp.Connect()

	// Set up logger based on configuration
	log.Logger = log.With().Caller().Logger()
	if cfg.AppIsDev {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}).With().Caller().Logger()
	}

	return &manager{
		config:     cfg,
		server:     srv,
		db:         db,
		dbGorm:     dbGorm,
		httpClient: clHttp,
	}, nil
}

// Implement the Manager interface methods
func (sm *manager) GetConfig() *config.Config {
	return sm.config
}

func (sm *manager) GetServer() *server.Server {
	return sm.server
}

func (sm *manager) GetDatabase() *sql.DB {
	return sm.db
}

func (sm *manager) GetGorm() *gorm.DB {
	return sm.dbGorm
}

func (sm *manager) GetHttp() httpClient.Http {
	return sm.httpClient
}
