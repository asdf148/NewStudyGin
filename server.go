package main

import (
	"asdf148.com/GinProject/controller"
	"asdf148.com/GinProject/model"
	"asdf148.com/GinProject/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	authService    service.AuthService       = service.New()
	authController controller.AuthController = controller.New(authService)
)

func main() {
	server := gin.Default()

	dsn := "Gin:Gin@tcp(127.0.0.1:3306)/GinTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.Novel{})

	authRoutes := server.Group("auth")
	{
		// 회원가입
		authRoutes.POST("/join", func(ctx *gin.Context) {
			ctx.JSON(200, authController.Join(ctx))
		})

		// 로그인
		authRoutes.POST("/login", func(ctx *gin.Context) {

		})
	}
	novelRoutes := server.Group("novel")
	{
		// 다 가져오기
		novelRoutes.GET("/", func(ctx *gin.Context) {

		})

		// 쓰기
		novelRoutes.POST("/", func(c *gin.Context) {

		})

		// 수정
		novelRoutes.PUT("/", func(c *gin.Context) {

		})

		// 삭젠
		novelRoutes.DELETE("/", func(c *gin.Context) {

		})
	}

	server.Run()
}
