package service

import (
	"net/http"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"github.com/gin-gonic/gin"
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

func UpdateUser(c *gin.Context) {
	var request request.UpdateUserRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		u := GetUserByEmail(request.Email)
		u.Name = request.Name
		u.Email = request.Email
		u.Description = request.Description
		u.HpUrl = request.HpUrl
		u.Location = request.Location
		repository.UpdateUser(u)
	}
}
