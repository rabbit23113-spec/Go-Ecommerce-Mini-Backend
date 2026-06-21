package handler

import (
	"main/internal/config"
	"main/internal/service"
)

type Handler struct {
	Service *service.Service
	Config  *config.Config
}
