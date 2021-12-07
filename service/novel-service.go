package service

import (
	dto "asdf148.com/GinProject/dto/novel"
	"asdf148.com/GinProject/model"
)

type NovelService interface {
	FindAll() []model.Novel
	Save(string, dto.CreateNovel) string
	Modify(string, string, dto.ModifyNovel) string
	Delete(string, uint) string
}

type novelService struct {
}

func NewNovelService() NovelService {
	return &novelService{}
}

func (service *novelService) FindAll() []model.Novel {
	db := database.Connect()

	var novels []model.Novel
	db.Find(&novels)

	return novels
}

func (service *novelService) Save(token string, createNovel dto.CreateNovel) string {
	db := database.Connect()

	userId, err := customUtil.ParseTokenWithSecretKey(token)
	if err != nil {
		return "error" + err.Error()
	}

	db.Create(&model.Novel{UserID: userId, Title: createNovel.Title, Context: createNovel.Content})

	return "Success"
}

func (service *novelService) Modify(token string, novelId string, modifyNovel dto.ModifyNovel) string {
	db := database.Connect()

	_, err := customUtil.ParseTokenWithSecretKey(token)
	if err != nil {
		return "error" + err.Error()
	}

	var novel model.Novel
	db.First(&novel, novelId)

	novel.Context = modifyNovel.Content

	db.Save(&novel)

	return "Modified"
}

func (service *novelService) Delete(token string, novelId uint) string {
	db := database.Connect()

	_, err := customUtil.ParseTokenWithSecretKey(token)
	if err != nil {
		return "error" + err.Error()
	}

	db.Delete(&model.Novel{}, novelId)

	return "Deleted"
}
