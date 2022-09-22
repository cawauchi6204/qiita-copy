package repository

import (
	"fmt"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
)

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

// TODO: 特定のidを入れてfollowしている人、followされている人を探す関数を作成する
func FindUserFollowsAll() {
	userFollows := []entity.UserFollow{}
	if err := DB.Find(&userFollows).Error; err != nil {
		fmt.Println(err)
	}
}
