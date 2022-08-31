package OrganizationRepository

import (
	"fmt"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

type Organization struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Description string    `json:"description"`
	HpUrl       string    `json:"hp_url"`
	Location    string    `json:"location"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindAll() {
	organizations := []Organization{}
	if err := repository.DB.Find(&organizations).Error; err != nil {
		fmt.Println(err)
	}
}
