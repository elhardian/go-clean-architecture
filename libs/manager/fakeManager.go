package manager

import (
	"database/sql"

	"github.com/elhardian/go-clean-architecture/libs/config"
	httpClient "github.com/elhardian/go-clean-architecture/libs/http"
	"github.com/elhardian/go-clean-architecture/libs/server"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// FakeManager is a fake implementation of the Manager interface for testing purposes.
type FakeManager struct {
	mock.Mock
	config     *config.Config
	server     *server.Server
	db         *sql.DB
	dbGorm     *gorm.DB
	httpClient httpClient.Http
}

// NewFakeManager creates a new instance of FakeManager.
func NewFakeManager() *FakeManager {
	return &FakeManager{}
}

// SetConfig sets the configuration for the fake manager.
func (fm *FakeManager) SetConfig(cfg *config.Config) {
	fm.config = cfg
}

// SetServer sets the server for the fake manager.
func (fm *FakeManager) SetServer(srv *server.Server) {
	fm.server = srv
}

// SetDatabase sets the database for the fake manager.
func (fm *FakeManager) SetDatabase(db *sql.DB) {
	fm.db = db
}

// SetGorm sets the Gorm database for the fake manager.
func (fm *FakeManager) SetGorm(dbGorm *gorm.DB) {
	fm.dbGorm = dbGorm
}

// SetHttp sets the HTTP client for the fake manager.
func (fm *FakeManager) SetHttp(httpClient httpClient.Http) {
	fm.httpClient = httpClient
}

// Implement the Manager interface methods using mock methods.

func (fm *FakeManager) GetConfig() *config.Config {
	return fm.config
}

func (fm *FakeManager) GetServer() *server.Server {
	return fm.server
}

func (fm *FakeManager) GetDatabase() *sql.DB {
	return fm.db
}

func (fm *FakeManager) GetGorm() *gorm.DB {
	return fm.dbGorm
}

func (fm *FakeManager) GetHttp() httpClient.Http {
	return fm.httpClient
}
