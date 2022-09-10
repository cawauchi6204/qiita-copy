package repository

import (
	"fmt"
	"time"
)

type UserFollow struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FollowUserId   string    `json:"followUserId"`
	FollowedUserId string    `json:"followedUserId"`
	CreatedAt      time.Time `json:"createdAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

// TODO: 特定のidを入れてfollowしている人、followされている人を探す関数を作成する
func FindUserFollowsAll() {
	userFollows := []UserFollow{}
	if err := DB.Find(&userFollows).Error; err != nil {
		fmt.Println(err)
	}
}
