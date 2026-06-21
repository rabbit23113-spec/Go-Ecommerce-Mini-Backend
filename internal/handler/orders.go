package handler

import (
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersHandler struct {
	Service *service.Service
}

func (oh *OrdersHandler) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find all")
}

func (oh *OrdersHandler) FindByUserId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find by user id")
}
func (oh *OrdersHandler) FindByProductId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find by product id")
}

func (oh *OrdersHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "mock create")
}

func (oh *OrdersHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (oh *OrdersHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock delete")
}
