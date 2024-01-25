// package/database/mysql.go

package database

import (
	"github.com/elhardian/go-clean-architecture/libs/config"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	*Options
}

func NewMySQL(cfg *config.Config) Database {
	return &MySQL{Options: NewDatabase(cfg).(*Options)}
}
