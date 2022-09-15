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
	ImgUrl         string    `json:"imgUrl"`
	HpUrl          string    `json:"hpUrl"`
	Location       string    `json:"location"`
	GithubAccount  string    `json:"githubAccountId"`
	OrganizationId string    `json:"organizationId"`
	IsDeleted      int       `json:"isDeleted"`
	CreatedAt      time.Time `json:"createdAt" sql:"DEFAULT:current_timestamp"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindUsersAll() (users []User) {
	users = []User{}
	if err := DB.Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindUserById(userId string) (user User) {
	user = User{}
	if err := DB.First(&user, "id = ?", userId).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindUsersById(userIds []string) (users []User) {
	users = []User{}
	if err := DB.Where("id IN ?", userIds).Find(&users).Error; err != nil {
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

func UpdateUser(user User) (result *gorm.DB) {
	result = DB.Save(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
