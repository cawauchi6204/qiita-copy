package repository

import (
	"fmt"
	"log"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
)

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindPostsAll() (posts []entity.Post) {
	posts = []entity.Post{}
	if err := DB.Find(&posts, "is_deleted = ?", 0).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindPostsAllByUserId(userId string) (posts []entity.Post) {
	posts = []entity.Post{}
	if err := DB.Find(&posts, "posted_by = ?", userId).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindPostsById(ids []string) (posts []entity.Post) {
	posts = []entity.Post{}
	if err := DB.Find(&posts, "id IN (?)", ids).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println("postsは")
	fmt.Println(posts)
	return
}

func FindPostByUserId(userId, postId string) (post entity.Post) {
	post = entity.Post{}
	if err := DB.Where("posted_by = ? AND id = ?", userId, postId).First(&post).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreatePost(post entity.Post) entity.Post {
	result := DB.Create(&post)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return entity.Post{
		ID:             post.ID,
		Title:          post.Title,
		Body:           post.Body,
		PostedBy:       post.PostedBy,
		OrganizationId: post.OrganizationId,
		IsDraft:        post.IsDraft,
		IsDeleted:      post.IsDeleted,
	}
}
