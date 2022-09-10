package service

import (
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
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

func GetUserByEmail(email string) (user repository.User) {
	user = repository.FindUserByEmail(email)
	return
}

func CreateUser(id, email, password string) (result *gorm.DB) {
	// TODO: これはEntityに持たせたい
	u := repository.User{
		ID:             id,
		Name:           "",
		Email:          email,
		Password:       password,
		Description:    "",
		HpUrl:          "",
		Location:       "",
		GithubAccount:  "",
		OrganizationId: "",
		IsDeleted:      0,
		CreatedAt:      time.Now(),
	}
	repository.CreateUser(u)
	return
}
