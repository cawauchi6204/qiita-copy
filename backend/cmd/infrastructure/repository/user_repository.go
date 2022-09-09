package repository

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Description    string    `json:"description"`
	HpUrl          string    `json:"hp_url"`
	Location       string    `json:"location"`
	GithubAccount  string    `json:"github_account_id"`
	OrganizationId string    `json:"organization_id"`
	IsDeleted      int       `json:"is_deleted"`
	CreatedAt      time.Time `json:"createdAt"`
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

func FindUserByEmail(email string) (user User) {
	user = User{}
	if err := DB.First(&user, "email = ?", email).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateUser(user User) (result *gorm.DB) {
	result = DB.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
