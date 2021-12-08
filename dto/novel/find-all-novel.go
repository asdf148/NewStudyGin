package dto

import "time"

type FindAllNovel interface {
}

type findAllNovel struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title" biding:"required"`
	Content   string    `json:"content" binding:"required"`
	UserID    uint      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewFindAllNovel(id uint, title string, content string, userId uint, createdAt time.Time) FindAllNovel {
	return &findAllNovel{
		ID:        id,
		Title:     title,
		Content:   content,
		UserID:    userId,
		CreatedAt: createdAt,
	}
}
