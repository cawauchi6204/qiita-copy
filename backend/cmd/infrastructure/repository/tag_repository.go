package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Tag struct {
	ID     string `json:"id"`
	ImgUrl string `json:"imgUrl"`
}

func FindTagById(id string) (tag Tag) {
	tag = Tag{}
	if err := DB.Find(&tag, "id = ?", id).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateTag(tag Tag) (result *gorm.DB) {
	result = DB.Create(&tag)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
