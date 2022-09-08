package repository

import (
	"fmt"
	"time"
)

type Post struct {
	ID             string    `json:"id"`
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

func FindPostsAllByUserId(userId int) (posts []Post) {
	posts = []Post{
		{
			ID:             "ID1",
			Title:          "Title1",
			Body:           "Body1",
			PostedBy:       1,
			OrganizationId: 1,
			IsDeleted:      0,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             "ID2",
			Title:          "Title2",
			Body:           "Body2",
			PostedBy:       1,
			OrganizationId: 1,
			IsDeleted:      0,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             "ID3",
			Title:          "Title3",
			Body:           "Body3",
			PostedBy:       1,
			OrganizationId: 1,
			IsDeleted:      0,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             "ID4",
			Title:          "Title4",
			Body:           "Body4",
			PostedBy:       1,
			OrganizationId: 1,
			IsDeleted:      0,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             "ID5",
			Title:          "Title5",
			Body:           "Body5",
			PostedBy:       1,
			OrganizationId: 1,
			IsDeleted:      0,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             "ID6",
			Title:          "Title6",
			Body:           "Body6",
			PostedBy:       1,
			OrganizationId: 1,
			IsDeleted:      0,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}
	// if err := DB.Find(&posts, "posted_by = ?", userId).Error; err != nil {
	// 	fmt.Println(err)
	// }
	return
}
