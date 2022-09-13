package repository

import (
	"fmt"
	"time"
)

type Like struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PostId     string    `json:"postId"`
	LikeUserId string    `json:"likeUserId"`
	CreatedAt  time.Time `json:"createdAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindLikesAll() {
	likes := []Like{}
	if err := DB.Find(&likes).Error; err != nil {
		fmt.Println(err)
	}
}

func FindLikeById(postId, userId string) (like Like) {
	if err := DB.Where("post_id = ? AND like_user_id = ?", postId, userId).Take(&like).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateLike(like Like) {
	if err := DB.Create(&like).Error; err != nil {
		fmt.Println(err)
	}
}

func DeleteLike(id uint) {
	like := Like{}
	if err := DB.Delete(&like, id).Error; err != nil {
		fmt.Println(err)
	}
}
