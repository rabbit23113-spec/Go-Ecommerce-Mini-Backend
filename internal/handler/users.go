package handler

import (
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	Service *service.Service
}

func (uh *UsersHandler) Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock me")
}

func (uh *UsersHandler) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find all")
}

func (uh *UsersHandler) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find by id")
}

func (uh *UsersHandler) FindByEmail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock find by email")
}

func (uh *UsersHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "mock create")
}

func (uh *UsersHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (uh *UsersHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock delete")
}
