package service

import (
	"main/internal/config"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	Config   *config.Config
	Postgres *sqlx.DB
}
