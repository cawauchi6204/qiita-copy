package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllPosts() (posts []entity.Post) {
	posts = repository.FindPostsAll()
	return
}

func GetAllPostsByUserId(userId string) (posts []entity.Post) {
	posts = repository.FindPostsAllByUserId(userId)
	return
}

func GetPostByUserId(userId, postId string) (post entity.Post) {
	post = repository.FindPostByUserId(userId, postId)
	return
}

func GetPostsByIds(ids []string) (posts []entity.Post) {
	posts = repository.FindPostsById(ids)
	return
}

func CreatePost(c *gin.Context) (post entity.Post) {
	var request request.CreatePostRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
		}
		uu := u.String()
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		p := entity.Post{
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
		post = repository.CreatePost(p)
	}
	return
}
