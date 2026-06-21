package handler

import (
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	Service *service.Service
}

func (ph *ProductsHandler) Find(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find")
}

func (ph *ProductsHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "mock create")
}

func (ph *ProductsHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (ph *ProductsHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock delete")
}
