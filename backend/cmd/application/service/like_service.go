package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"github.com/gin-gonic/gin"
)

func GetLikesByPostId(postId string) (likes []repository.Like) {
	likes = repository.FindLikesByPostId(postId)
	return
}

func UpdateLike(c *gin.Context) {
	var request request.UpdateLikeRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		// いいねがなければ追加、あれば削除
		like := repository.FindLikeById(request.PostId, request.UserId)
		// Like{}の中身が空だったらとかエラーが起きていたらなどの条件分岐にさせたい
		if like.PostId == "" {
			fmt.Println("")
			repository.CreateLike(repository.Like{
				PostId:     request.PostId,
				LikeUserId: request.UserId,
				CreatedAt:  time.Now(),
			})
		} else {
			repository.DeleteLike(like.ID)
		}
	}
}
