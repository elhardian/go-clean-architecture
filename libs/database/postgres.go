// package/database/postgres.go

package database

import (
	"github.com/elhardian/go-clean-architecture/libs/config"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	*Options
}

func NewPostgreSQL(cfg *config.Config) Database {
	return &PostgreSQL{Options: NewDatabase(cfg).(*Options)}
}
