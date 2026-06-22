package handler

import (
	"main/internal/service"
	"main/package/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	Service *service.Service
}

func (uh *UsersHandler) Me(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := uh.Service.UserService.Me(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) FindAll(ctx *gin.Context) {
	resp, err := uh.Service.UserService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) FindById(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := uh.Service.UserService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) FindByEmail(ctx *gin.Context) {
	var req dto.FindByEmailDto
	ctx.ShouldBindJSON(&req)
	resp, err := uh.Service.UserService.FindByEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) Create(ctx *gin.Context) {
	var req dto.CreateUserDto
	ctx.ShouldBindJSON(req)
	resp, err := uh.Service.UserService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (uh *UsersHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (uh *UsersHandler) Delete(ctx *gin.Context) {
	id, err := ctx.Params.Get("id")
	if err != true {
		ctx.JSON(http.StatusBadRequest, "You should provide an id parameter")
	}
	uh.Service.UserService.Delete(id)
}
