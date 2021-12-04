package myUtil

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtUtil interface {
	CreateAccessToken(userId uint) (string, error)
}

type jwtUtil struct {
}

func New() JwtUtil {
	return &jwtUtil{}
}

func (c *jwtUtil) CreateAccessToken(userId uint) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("testClaims"))
	if err != nil {
		panic(err)
	}
	return token, err
}
