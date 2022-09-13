package repository

import (
	"log"

	"gorm.io/gorm"
)

type Tag struct {
	ID     string `json:"id"`
	ImgUrl string `json:"imgUrl"`
}

func CreateTags(tag Tag) (result *gorm.DB) {
	result = DB.Create(&tag)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
