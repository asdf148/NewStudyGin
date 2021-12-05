package myUtil

import (
	"errors"
	"time"

	dto "asdf148.com/GinProject/dto/jwt"
	"github.com/dgrijalva/jwt-go"
)

type JwtUtil interface {
	CreateAccessToken(userId uint) (string, error)
	ParseTokenWithSecretKey(string) (uint, error)
}

type jwtUtil struct {
}

func New() JwtUtil {
	return &jwtUtil{}
}

func (c *jwtUtil) CreateAccessToken(userId uint) (string, error) {

	claims := &dto.JwtClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
			Issuer:    "me",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("testClaims"))
	if err != nil {
		return "error", err
	}

	return signedToken, nil
}

func (c *jwtUtil) ParseTokenWithSecretKey(signedToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&dto.JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("testClaims"), nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*dto.JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return 0, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return 0, err
	}

	return claims.UserId, nil
}
