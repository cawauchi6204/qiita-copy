package repository

import (
	"fmt"
	"time"
)

type Like struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LikeUserId  int       `json:"like_user_id"`
	LikedUserId int       `json:like_user_id`
	CreatedAt   time.Time `json:"createdAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindLikesAll() {
	likes := []Like{}
	if err := DB.Find(&likes).Error; err != nil {
		fmt.Println(err)
	}
}
