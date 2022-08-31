package LikeRepository

import (
	"fmt"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

type Like struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LikeUserId  int       `json:"like_user_id"`
	LikedUserId int       `json:like_user_id`
	CreatedAt   time.Time `json:"createdAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindAll() {
	likes := []Like{}
	if err := repository.DB.Find(&likes).Error; err != nil {
		fmt.Println(err)
	}
}
