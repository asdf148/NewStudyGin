package service

import (
	"fmt"

	"asdf148.com/GinProject/DB"
	dto "asdf148.com/GinProject/dto/auth"
	"asdf148.com/GinProject/model"
	"asdf148.com/GinProject/myUtil"
	"golang.org/x/crypto/bcrypt"
)

var (
	database   DB.ConnectDB   = DB.New()
	customUtil myUtil.JwtUtil = myUtil.New()
)

type AuthService interface {
	Join(dto.SignUp) string
	Login(dto.Login) string
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (service *authService) Join(signUp dto.SignUp) string {
	db := database.Connect()

	// 	비밀번호 암호화
	hash, err := bcrypt.GenerateFromPassword([]byte(signUp.Password), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}

	// 저장
	db.Create(&model.User{Name: signUp.Name, Age: signUp.Age, Email: signUp.Email, Password: string(hash)})
	fmt.Println("sign up:")
	fmt.Println(signUp)
	return "success"
}

func (service *authService) Login(login dto.Login) string {
	db := database.Connect()

	var user model.User
	db.First(&user, "email = ?", login.Email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "password is different"
	}

	token, err := customUtil.CreateAccessToken(user.ID)
	if err != nil {
		return "fail to make token"
	}

	return token
}
