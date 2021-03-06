package main

import (
	"net/http"

	"asdf148.com/GinProject/DB"
	"asdf148.com/GinProject/controller"
	"asdf148.com/GinProject/middleware"
	"asdf148.com/GinProject/model"
	"asdf148.com/GinProject/service"
	"github.com/gin-gonic/gin"
)

var (
	database DB.ConnectDB = DB.New()

	authService     service.AuthService        = service.NewAuthService()
	authController  controller.AuthController  = controller.NewAuthController(authService)
	novelService    service.NovelService       = service.NewNovelService()
	novelController controller.NovelController = controller.NewNovelController(novelService)

	tokenMiddleware middleware.TokenVerification = middleware.NewTokenVerification()
)

func main() {
	server := gin.Default()

	db := database.Connect()
	db.AutoMigrate(&model.User{}, &model.Novel{})

	authRoutes := server.Group("auth")
	{
		// 회원가입
		authRoutes.POST("/join", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, authController.Join(ctx))
		})

		// 로그인
		authRoutes.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, authController.Login(ctx))
		})
	}
	novelRoutes := server.Group("novel")
	{
		// 다 가져오기
		novelRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, novelController.FindAll())
		})
		novelRoutes.Use(tokenMiddleware.TokenVerify)
		// 쓰기
		novelRoutes.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, novelController.Save(ctx))
		})

		// 수정
		novelRoutes.PUT("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, novelController.Modify(ctx))
		})

		// 삭젠
		novelRoutes.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, novelController.Delete(ctx))
		})
	}

	server.Run(":3000")
}
