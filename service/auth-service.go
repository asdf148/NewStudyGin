package service

import dto "asdf148.com/GinProject/dto/auth"

type AuthService interface {
	Join(dto.SignUp) string
	Login(dto.Login) string
}

type authService struct {
}

func New() AuthService {
	return &authService{}
}

func (service *authService) Join(signUp dto.SignUp) string {
	return "success"
}

func (service *authService) Login(login dto.Login) string {
	return "success"
}
