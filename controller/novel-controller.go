package controller

import (
	"fmt"
	"strings"

	dto "asdf148.com/GinProject/dto/novel"
	"asdf148.com/GinProject/service"
	"github.com/gin-gonic/gin"
)

type NovelController interface {
	FindAll() gin.H
	Save(ctx *gin.Context) gin.H
	Modify(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type novelController struct {
	service service.NovelService
}

func NewNovelController(service service.NovelService) NovelController {
	return &novelController{
		service: service,
	}
}

func (c *novelController) FindAll() gin.H {
	return gin.H{
		"users": c.service.FindAll(),
	}
}

func (c *novelController) Save(ctx *gin.Context) gin.H {
	bearerToken := ctx.Request.Header["Authorization"][0]
	fmt.Println(bearerToken)
	token := strings.Split(bearerToken, " ")[1]

	var createNovel dto.CreateNovel
	ctx.ShouldBindJSON(&createNovel)

	return gin.H{
		"message": c.service.Save(token, createNovel),
	}
}

func (c *novelController) Modify(ctx *gin.Context) {

}

func (c *novelController) Delete(ctx *gin.Context) {

}
