package controller

import (
	"fmt"
	"strconv"

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
	var createNovel dto.CreateNovel
	ctx.ShouldBindJSON(&createNovel)

	UserId := fmt.Sprint(ctx.Keys["userId"])
	userId, err := strconv.ParseUint(UserId, 10, 64)
	if err != nil {
		panic(err)
	}

	return gin.H{
		"message": c.service.Save(uint(userId), createNovel),
	}
}

func (c *novelController) Modify(ctx *gin.Context) gin.H {
	UserId := fmt.Sprint(ctx.Keys["userId"])
	userId, err := strconv.ParseUint(UserId, 10, 64)
	if err != nil {
		panic(err)
	}

	var modifyNovel dto.ModifyNovel
	ctx.ShouldBindJSON(&modifyNovel)

	return gin.H{
		"message": c.service.Modify(uint(userId), ctx.Query("novel"), modifyNovel),
	}
}

func (c *novelController) Delete(ctx *gin.Context) gin.H {
	UserId := fmt.Sprint(ctx.Keys["userId"])
	userId, err := strconv.ParseUint(UserId, 10, 64)
	if err != nil {
		panic(err)
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return gin.H{
			"message": "string to uint convert failed",
		}
	}

	return gin.H{
		"message": c.service.Delete(uint(userId), uint(id)),
	}
}
