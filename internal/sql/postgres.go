package sql

import (
	"main/internal/config"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	config *config.Config
}

func (p *Postgres) NewPostgres(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", p.config.DBConfig.Url)
	if err != nil {
		return &sqlx.DB{}, err
	}
	return db, nil
}
