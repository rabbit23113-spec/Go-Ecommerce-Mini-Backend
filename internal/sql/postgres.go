package db

import (
	"main/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Config *config.Config
}

func (p *Postgres) NewPostgres() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", p.Config.DBConfig.Url)
	if err != nil {
		return &sqlx.DB{}, err
	}
	return db, nil
}
