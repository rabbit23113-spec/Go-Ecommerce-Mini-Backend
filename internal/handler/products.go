package handler

import (
	"main/internal/service"
	"main/package/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	Service *service.Service
}

func (ph *ProductsHandler) Find(ctx *gin.Context) {
	resp, err := ph.Service.ProductService.Find()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductsHandler) Create(ctx *gin.Context) {
	var body dto.CreateProductDto
	ctx.ShouldBindJSON(&body)
	resp, err := ph.Service.ProductService.Create(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (ph *ProductsHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (ph *ProductsHandler) Delete(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if ok != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "you should define an id parameter"})
		return
	}
	if err := ph.Service.ProductService.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusNoContent, "")
}
