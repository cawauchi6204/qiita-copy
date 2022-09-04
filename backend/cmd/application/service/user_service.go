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

func GetUserById(userId int) (user repository.User) {
	user = repository.FindUserById(userId)
	return
}

func GetUserByEmail(email string) (user repository.User) {
	user = repository.FindUserByEmail(email)
	return
}

func CreateUser(name, email, password string) (result *gorm.DB) {
	// TODO: これはEntityに持たせたい
	u := repository.User{
		Name:           name,
		Email:          email,
		Password:       password,
		Nickname:       "",
		Description:    "",
		HpUrl:          "",
		Location:       "",
		GithubAccount:  "",
		OrganizationId: 0,
		IsDeleted:      0,
		CreatedAt:      time.Now(),
	}
	result = repository.CreateUser(u)
	return
}
