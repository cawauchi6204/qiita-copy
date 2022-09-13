package repository

import (
	"fmt"
	"log"
)

type PostTag struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	PostId string `json:"postId"`
	TagId  string `json:"tagId"`
}

func FindTagsByPostId(postId string) (postTags []PostTag) {
	postTags = []PostTag{}
	if err := DB.Find(&postTags, "post_id = ?", postId).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreatePostTag(postTag PostTag) PostTag {
	result := DB.Create(&postTag)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return PostTag{
		ID:     postTag.ID,
		PostId: postTag.PostId,
		TagId:  postTag.TagId,
	}
}
