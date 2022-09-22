package repository

import (
	"fmt"

	"github.com/cawauchi6204/qiita-copy/cmd/entity"
)

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindOrganizationsAll() {
	organizations := []entity.Organization{}
	if err := DB.Find(&organizations).Error; err != nil {
		fmt.Println(err)
	}
}
