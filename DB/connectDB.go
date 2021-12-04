package DB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConnectDB interface {
	Connect() *gorm.DB
}

type connectDB struct {
}

func New() ConnectDB {
	return &connectDB{}
}

func (c *connectDB) Connect() *gorm.DB {
	dsn := "Gin:Gin@tcp(127.0.0.1:3306)/GinTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	return db
}
