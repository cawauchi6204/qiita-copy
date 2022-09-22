package repository

import (
	"fmt"
	"log"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
	"gorm.io/gorm"
)

func FindTagById(id string) (tag entity.Tag) {
	tag = entity.Tag{}
	if err := DB.Find(&tag, "id = ?", id).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateTag(tag entity.Tag) (result *gorm.DB) {
	result = DB.Create(&tag)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
