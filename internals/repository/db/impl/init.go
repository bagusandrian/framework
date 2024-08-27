package impl

import (
	"github.com/bagusandrian/framework/internals/config"
	"github.com/bagusandrian/framework/internals/repository/db"

	"database/sql"
)

type repoDB struct {
	cfg *config.Config
	db  *sql.DB
}

func New(cfg *config.Config) db.DB {
	return &repoDB{
		cfg: cfg,
		db:  &sql.DB{},
	}
}
