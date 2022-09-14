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

func FindPostsIdsByTagId(tagId string) (postIds []string) {
	postTag := []PostTag{}
	if err := DB.Find(&postTag, "tag_id = ?", tagId).Error; err != nil {
		fmt.Println(err)
	}

	postIds = []string{}
	for _, v := range postTag {
		postIds = append(postIds, v.PostId)
	}
	fmt.Println(postIds)
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
