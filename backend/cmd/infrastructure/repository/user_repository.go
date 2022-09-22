package repository

import (
	"fmt"
	"log"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
	"gorm.io/gorm"
)

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindUsersAll() (users []entity.User) {
	users = []entity.User{}
	if err := DB.Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindUserById(userId string) (user entity.User) {
	user = entity.User{}
	if err := DB.First(&user, "id = ?", userId).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindUsersById(userIds []string) (users []entity.User) {
	users = []entity.User{}
	if err := DB.Where("id IN ?", userIds).Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindUserByEmail(email string) (user entity.User) {
	user = entity.User{}
	if err := DB.First(&user, "email = ?", email).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateUser(user entity.User) (result *gorm.DB) {
	result = DB.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}

func UpdateUser(user entity.User) (result *gorm.DB) {
	result = DB.Save(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
