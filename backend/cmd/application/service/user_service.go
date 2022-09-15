package service

import (
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"gorm.io/gorm"
)

func GetAllUsers() (users []repository.User) {
	users = repository.FindUsersAll()
	return
}

func GetUserById(userId string) (user repository.User) {
	user = repository.FindUserById(userId)
	return
}

func GetUsersById(userIds []string) (users []repository.User) {
	users = repository.FindUsersById(userIds)
	return
}

func GetUserByEmail(email string) (user repository.User) {
	user = repository.FindUserByEmail(email)
	return
}

func CreateUser(id, email, password string) (result *gorm.DB) {
	// TODO: これはEntityに持たせたい
	u := repository.User{
		ID:        id,
		Email:     email,
		Password:  password,
		IsDeleted: 0,
		CreatedAt: time.Now(),
	}
	repository.CreateUser(u)
	return
}

func UpdateUser(u repository.User, request request.UpdateUserRequest) {
	// TODO: 冗長すぎてやばいのでgolangでの三項演算子的なものが知りたい
	if request.Name != "" {
		u.Name = request.Name
	}
	if request.Description != "" {
		u.Email = request.Email
	}
	if request.HpUrl != "" {
		u.Description = request.Description
	}
	if request.HpUrl != "" {
		u.HpUrl = request.HpUrl
	}
	if request.Location != "" {
		u.Location = request.Location
	}
	if request.ImgUrl != "" {
		u.ImgUrl = request.ImgUrl
	}
	repository.UpdateUser(u)
}
