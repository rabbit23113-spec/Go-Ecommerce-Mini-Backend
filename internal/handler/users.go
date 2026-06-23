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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) FindAll(ctx *gin.Context) {
	resp, err := uh.Service.UserService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) FindById(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if ok != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "You should provide an id parameter"})
		return
	}
	resp, err := uh.Service.UserService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) FindByEmail(ctx *gin.Context) {
	var req dto.FindByEmailDto
	ctx.ShouldBindJSON(&req)
	resp, err := uh.Service.UserService.FindByEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (uh *UsersHandler) Create(ctx *gin.Context) {
	var req dto.CreateUserDto
	ctx.ShouldBindJSON(req)
	resp, err := uh.Service.UserService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (uh *UsersHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock update")
}

func (uh *UsersHandler) Delete(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if ok != true {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "You should provide an id parameter"})
		return
	}
	if err := uh.Service.UserService.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
