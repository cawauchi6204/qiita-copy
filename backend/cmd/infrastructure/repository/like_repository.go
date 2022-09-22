package repository

import (
	"fmt"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
)

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい
func FindLikesByPostId(postId string) (likes []entity.Like) {
	if err := DB.Where("post_id = ?", postId).Find(&likes).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindLikeById(postId, userId string) (like entity.Like) {
	if err := DB.Where("post_id = ? AND like_user_id = ?", postId, userId).Take(&like).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func CreateLike(like entity.Like) {
	if err := DB.Create(&like).Error; err != nil {
		fmt.Println(err)
	}
}

func DeleteLike(id uint) {
	like := entity.Like{}
	if err := DB.Delete(&like, id).Error; err != nil {
		fmt.Println(err)
	}
}
