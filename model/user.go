package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Age      int8
	Email    string
	Password string
	Novels   []Novel
}
