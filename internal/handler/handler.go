package handler

import (
	"main/internal/config"
	"main/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Config  *config.Config
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/signup", nil)
		auth.POST("/signin", nil)
		auth.POST("/logout", nil)
		auth.POST("/revoke", nil)
	}
	users := r.Group("/users")
	{
		users.GET("/me", nil)
		users.GET("/find/all", nil)
		users.GET("/find/id/", nil)
		users.GET("/find/email/", nil)
		users.POST("/create", nil)
		users.PATCH("/update", nil)
		users.DELETE("/delete", nil)
	}
	products := r.Group("/products")
	{
		products.GET("/find", nil)
		products.POST("/create", nil)
		products.PATCH("/update", nil)
		products.DELETE("/delete", nil)
	}
	warehouses := r.Group("/warehouses")
	{
		warehouses.GET("/find", nil)
		warehouses.GET("/find/product", nil)
		warehouses.POST("/create", nil)
		warehouses.PATCH("/update", nil)
		warehouses.DELETE("/delete", nil)
	}
	orders := r.Group("/orders")
	{
		orders.GET("/find/all", nil)
		orders.GET("/find/user", nil)
		orders.GET("/find/product", nil)
		orders.POST("/create", nil)
		orders.PATCH("/update", nil)
		orders.DELETE("/delete", nil)
	}
	return r
}
