package main

import (
	"log"
	"main/internal/config"
	"main/internal/handler"
	"main/internal/service"
	db "main/internal/sql"
)

func main() {
	cfg := config.Config{}
	config, err := cfg.ReadConfig()
	if err != nil {
		log.Fatalf("Error while reading the config: %s", err)
	}
	db := db.Postgres{Config: config}
	conn, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s", err)
	}
	service := service.Service{Config: config, Postgres: conn}
	handler := handler.Handler{Config: config, Service: &service}
}
