package repository

import (
	"fmt"
	"log"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
)

func FindTagsByPostId(postId string) (postTags []entity.PostTag) {
	postTags = []entity.PostTag{}
	if err := DB.Find(&postTags, "post_id = ?", postId).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindPostsIdsByTagId(tagId string) (postIds []string) {
	postTag := []entity.PostTag{}
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

func CreatePostTag(postTag entity.PostTag) entity.PostTag {
	result := DB.Create(&postTag)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return entity.PostTag{
		ID:     postTag.ID,
		PostId: postTag.PostId,
		TagId:  postTag.TagId,
	}
}
