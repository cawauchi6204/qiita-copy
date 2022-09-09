package repository

import (
	"fmt"
	"time"
)

type Organization struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Description string    `json:"description"`
	HpUrl       string    `json:"hp_url"`
	Location    string    `json:"location"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindOrganizationsAll() {
	organizations := []Organization{}
	if err := DB.Find(&organizations).Error; err != nil {
		fmt.Println(err)
	}
}
