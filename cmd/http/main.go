package main

import (
	"log"
	"main/internal/config"
	"main/internal/handler"
	"main/internal/service"
	db "main/internal/sql"
	"net/http"
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
	service := service.Service{Postgres: conn}
	srvc := service.InitService()
	handler := handler.Handler{Config: config, Service: srvc}
	srv := new(Server)
	srv.Serve(cfg.Port, handler.InitRoutes())
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Serve(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 28,
		Handler:        handler,
	}
	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while listening to the server: %s", err)
	}
	return nil
}
