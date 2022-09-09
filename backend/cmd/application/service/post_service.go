package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllPosts() (posts []repository.Post) {
	posts = repository.FindPostsAll()
	return
}

func GetAllPostsByUserId(userId string) (posts []repository.Post) {
	posts = repository.FindPostsAllByUserId(userId)
	return
}

func CreatePost(c *gin.Context) (result *gorm.DB) {
	var request request.CreatePostRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
			return
		}
		uu := u.String()
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		p := repository.Post{
			ID:             uu,
			Title:          request.Title,
			Body:           request.Body,
			PostedBy:       request.PostedBy,
			OrganizationId: request.OrganizationId,
			IsDraft:        request.IsDraft,
			IsDeleted:      request.IsDeleted,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		repository.CreatePost(p)
	}
	return
}
