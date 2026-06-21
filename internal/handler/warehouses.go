package handler

import (
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WarehousesHandler struct {
	Service *service.Service
}

func (wh *WarehousesHandler) Find(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find")
}

func (wh *WarehousesHandler) FindByProductId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find by product id")
}

func (wh *WarehousesHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "mock create")
}

func (wh *WarehousesHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (wh *WarehousesHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock delete")
}
