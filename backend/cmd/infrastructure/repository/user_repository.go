package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Nickname       string    `json:"nickname"`
	HpUrl          string    `json:"hp_url"`
	Location       string    `json:"location"`
	GithubAccount  string    `json:"github_account_id"`
	OrganizationId int       `json:"organization_id"`
	IsDeleted      int       `json:"is_deleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindUsersAll() (users []User) {
	users = []User{}
	if err := DB.Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindUserById(userId int) (user User) {
	user = User{}
	if err := DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func SaveUser(user User) (result *gorm.DB) {
	if err := DB.Create(&user).Error; err != nil {
		fmt.Println(err)
	}
	return DB.Create(&user)
}
