package handler

import (
	"main/internal/service"
	"main/package/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.Service
}

func (ah *AuthHandler) SignUp(ctx *gin.Context) {
	var req dto.CreateUserDto
	ctx.ShouldBindJSON(&req)
	ah.Service.AuthService.SignUp(req)
}

func (ah *AuthHandler) SignIn(ctx *gin.Context) {
	var req dto.SignInDto
	ctx.ShouldBindJSON(&req)
	ah.Service.AuthService.SignIn(req)
}

func (ah *AuthHandler) LogOut(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock log out")
}

func (ah *AuthHandler) Revoke(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock revoke")
}
