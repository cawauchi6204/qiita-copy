package UserRepository

import (
	"fmt"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Nickname       string    `json:"nickname"`
	HpUrl          string    `json:"hp_url"`
	Location       string    `json:"location"`
	GithubAccount  string    `json:"github_account_id"`
	OrganizationId int       `json:"organization_id"`
	IsDeleted      int       `json:"is_deleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// TODO: 今は特化関数だがBuilderパターンにしてuseCase層でimportしてビジネスロジック関数を組み立てたい

func FindAll() (users []User) {
	users = []User{}
	if err := repository.DB.Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return
}

func FindById(userId int) (user User) {
	user = User{}
	if err := repository.DB.First(&user, userId).Error; err != nil {
		fmt.Println(err)
	}
	return
}
