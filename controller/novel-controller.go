package controller

import (
	"strconv"
	"strings"

	dto "asdf148.com/GinProject/dto/novel"
	"asdf148.com/GinProject/service"
	"github.com/gin-gonic/gin"
)

type NovelController interface {
	FindAll() gin.H
	Save(ctx *gin.Context) gin.H
	Modify(ctx *gin.Context) gin.H
	Delete(ctx *gin.Context) gin.H
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
		"novels": c.service.FindAll(),
	}
}

func (c *novelController) Save(ctx *gin.Context) gin.H {
	bearerToken := ctx.Request.Header["Authorization"][0]
	token := strings.Split(bearerToken, " ")[1]

	var createNovel dto.CreateNovel
	ctx.ShouldBindJSON(&createNovel)

	return gin.H{
		"message": c.service.Save(token, createNovel),
	}
}

func (c *novelController) Modify(ctx *gin.Context) gin.H {
	bearerToken := ctx.Request.Header["Authorization"][0]
	token := strings.Split(bearerToken, " ")[1]

	var modifyNovel dto.ModifyNovel
	ctx.ShouldBindJSON(&modifyNovel)

	return gin.H{
		"message": c.service.Modify(token, ctx.Query("novel"), modifyNovel),
	}
}

func (c *novelController) Delete(ctx *gin.Context) gin.H {
	bearerToken := ctx.Request.Header["Authorization"][0]
	token := strings.Split(bearerToken, " ")[1]

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return gin.H{
			"message": "string to uint convert failed",
		}
	}

	return gin.H{
		"message": c.service.Delete(token, uint(id)),
	}
}
