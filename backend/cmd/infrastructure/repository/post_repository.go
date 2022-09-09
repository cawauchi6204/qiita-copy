package repository

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	PostedBy       string    `json:"posted_by"`
	OrganizationId string    `json:"organization_id"`
	IsDraft        int       `json:"is_draft"`
	IsDeleted      int       `json:"is_deleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindPostsAll() (posts []Post) {
	posts = []Post{}
	if err := DB.Find(&posts, "is_deleted = ?", 0).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindPostsAllByUserId(userId string) (posts []Post) {
	posts = []Post{}
	if err := DB.Find(&posts, "posted_by = ?", userId).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreatePost(post Post) (result *gorm.DB) {
	result = DB.Create(&post)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
