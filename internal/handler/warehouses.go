package handler

import (
	"main/internal/service"
	"main/package/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WarehousesHandler struct {
	Service *service.Service
}

func (wh *WarehousesHandler) Find(ctx *gin.Context) {
	resp, err := wh.Service.WarehouseService.Find()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (wh *WarehousesHandler) FindByProductId(ctx *gin.Context) {
	productId, ok := ctx.Params.Get("id")
	if ok != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "You should provide a product id"})
	}
	resp, err := wh.Service.WarehouseService.FindByProductId(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (wh *WarehousesHandler) Create(ctx *gin.Context) {
	var body dto.CreateWarehouseDto
	ctx.ShouldBindJSON(&body)
	resp, err := wh.Service.WarehouseService.Create(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (wh *WarehousesHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (wh *WarehousesHandler) Delete(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if ok != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "You should provide a product id"})
	}
	err := wh.Service.WarehouseService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, "")
}
