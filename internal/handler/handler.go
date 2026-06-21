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
	authHandler := AuthHandler{Service: h.Service}
	usersHandler := UsersHandler{Service: h.Service}
	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/signin", authHandler.SignIn)
		auth.POST("/logout", authHandler.LogOut)
		auth.POST("/revoke", authHandler.Revoke)
	}
	users := r.Group("/users")
	{
		users.GET("/me", usersHandler.Me)
		users.GET("/find/all", usersHandler.FindAll)
		users.GET("/find/id/", usersHandler.FindById)
		users.GET("/find/email/", usersHandler.FindByEmail)
		users.POST("/create", usersHandler.Create)
		users.PATCH("/update", usersHandler.Update)
		users.DELETE("/delete", usersHandler.Delete)
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
