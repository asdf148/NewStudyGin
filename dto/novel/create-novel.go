package dto

type CreateNovel struct {
	Title   string `json:"title" biding:"required"`
	Content string `json:"content" binding:"required"`
}
