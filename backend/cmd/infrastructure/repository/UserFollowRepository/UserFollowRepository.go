package UserFollowRepository

import (
	"fmt"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

type UserFollow struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FollowUserId   int       `json:"follow_user_id"`
	FollowedUserId int       `json:"followed_user_id"`
	CreatedAt      time.Time `json:"createdAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

// TODO: 特定のidを入れてfollowしている人、followされている人を探す関数を作成する
func FindAll() {
	userFollows := []UserFollow{}
	if err := repository.DB.Find(&userFollows).Error; err != nil {
		fmt.Println(err)
	}
}
