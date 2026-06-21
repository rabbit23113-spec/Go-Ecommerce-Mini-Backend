package handler

import (
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.Service
}

func (ah *AuthHandler) SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "mock sign up")
}

func (ah *AuthHandler) SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "mock sign in")
}

func (ah *AuthHandler) LogOut(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock log out")
}

func (ah *AuthHandler) Revoke(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, "mock revoke")
}
