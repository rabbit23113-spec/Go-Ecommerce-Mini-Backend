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
	productsHandler := ProductsHandler{Service: h.Service}
	warehousesHandler := WarehousesHandler{Service: h.Service}
	ordersHandler := OrdersHandler{Service: h.Service}

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
		products.GET("/find", productsHandler.Find)
		products.POST("/create", productsHandler.Create)
		products.PATCH("/update", productsHandler.Update)
		products.DELETE("/delete", productsHandler.Delete)
	}

	warehouses := r.Group("/warehouses")
	{
		warehouses.GET("/find", warehousesHandler.Find)
		warehouses.GET("/find/product", warehousesHandler.FindByProductId)
		warehouses.POST("/create", warehousesHandler.Create)
		warehouses.PATCH("/update", warehousesHandler.Update)
		warehouses.DELETE("/delete", warehousesHandler.Delete)
	}

	orders := r.Group("/orders")
	{
		orders.GET("/find/all", ordersHandler.FindAll)
		orders.GET("/find/user", ordersHandler.FindByUserId)
		orders.GET("/find/product", ordersHandler.FindByProductId)
		orders.POST("/create", ordersHandler.Create)
		orders.PATCH("/update", ordersHandler.Update)
		orders.DELETE("/delete", ordersHandler.Delete)
	}
	return r
}
