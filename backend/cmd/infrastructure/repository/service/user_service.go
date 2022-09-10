package service

import "github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"

func GetAllUsers() {
	repository.FindUsersAll()
}

func GetUserById(userId string) {
	repository.FindUserById(userId)
}
