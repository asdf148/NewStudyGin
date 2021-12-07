package dto

type ModifyNovel struct {
	Content string `json:"content" binding:"required"`
}
