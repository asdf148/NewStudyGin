package dto

import "github.com/dgrijalva/jwt-go"

type JwtClaim struct {
	UserId uint
	jwt.StandardClaims
}
