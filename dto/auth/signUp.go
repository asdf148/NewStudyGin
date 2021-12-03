package dto

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Age      int8   `json:"age"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,max=10"`
}
