package service

import "github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository/UserRepository"

func GetAllUser() (users []UserRepository.User) {
	return UserRepository.FindAll()
}

func GetUserById(userId int) (user UserRepository.User) {
	return UserRepository.FindById(userId)
}
