package model

import "gorm.io/gorm"

type Novel struct {
	gorm.Model
	ID      uint
	UserID  uint
	Title   string
	Context string
}
