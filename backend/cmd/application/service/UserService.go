package service

import (
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository/UserRepository"
	"gorm.io/gorm"
)

func GetAllUser() (users []UserRepository.User) {
	return UserRepository.FindAll()
}

func GetUserById(userId int) (user UserRepository.User) {
	return UserRepository.FindById(userId)
}

func SaveUser(user UserRepository.User) (result *gorm.DB) {
	return UserRepository.Save(user)
}
