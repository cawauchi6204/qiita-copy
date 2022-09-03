package service

import (
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

func GetAllUsers() (users []repository.User) {
	users = repository.FindUsersAll()
	return
}

func GetUserById(userId int) (user repository.User) {
	user = repository.FindUserById(userId)
	return
}

func CreateUser(user repository.User) {
	// TODO: これはEntityに持たせたい
	u := repository.User{
		Name:           user.Name,
		Email:          user.Email,
		Password:       user.Password,
		Nickname:       "",
		Description:    "",
		HpUrl:          "",
		Location:       "",
		GithubAccount:  "",
		OrganizationId: 0,
		IsDeleted:      0,
		CreatedAt:      time.Now(),
	}
	repository.CreateUser(u)
}
