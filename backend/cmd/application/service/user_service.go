package service

import (
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"gorm.io/gorm"
)

func GetAllUsers() (users []entity.User) {
	users = repository.FindUsersAll()
	return
}

func GetUserById(userId string) (user entity.User) {
	user = repository.FindUserById(userId)
	return
}

func GetUsersById(userIds []string) (users []entity.User) {
	users = repository.FindUsersById(userIds)
	return
}

func GetUserByEmail(email string) (user entity.User) {
	user = repository.FindUserByEmail(email)
	return
}

func CreateUser(id, email, password string) (result *gorm.DB) {
	// TODO: これはEntityに持たせたい
	u := entity.User{
		ID:        id,
		Email:     email,
		Password:  password,
		IsDeleted: 0,
		CreatedAt: time.Now(),
	}
	repository.CreateUser(u)
	return
}

func UpdateUser(u entity.User, request request.UpdateUserRequest) {
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
