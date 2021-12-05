package controller

import (
	dto "asdf148.com/GinProject/dto/auth"
	"asdf148.com/GinProject/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Join(ctx *gin.Context) gin.H
	Login(ctx *gin.Context) gin.H
}

type authController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &authController{
		service: service,
	}
}

func (c *authController) Join(ctx *gin.Context) gin.H {
	var user dto.SignUp
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	return gin.H{
		"message": c.service.Join(user),
	}
}

func (c *authController) Login(ctx *gin.Context) gin.H {
	var user dto.Login
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	return gin.H{
		"message": c.service.Login(user),
	}
}
