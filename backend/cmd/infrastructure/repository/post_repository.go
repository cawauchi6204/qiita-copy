package repository

import (
	"fmt"
	"time"
)

type Post struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	PostedBy       int       `json:"posted_by"`
	OrganizationId int       `json:"organization_id"`
	IsDeleted      int       `json:"is_deleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindPostsAll() {
	posts := []Post{}
	if err := DB.Find(&posts).Error; err != nil {
		fmt.Println(err)
	}
}
