package entity

import "time"

type User struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Description    string    `json:"description"`
	ImgUrl         string    `json:"imgUrl"`
	HpUrl          string    `json:"hpUrl"`
	Location       string    `json:"location"`
	GithubAccount  string    `json:"githubAccountId"`
	OrganizationId string    `json:"organizationId"`
	IsDeleted      int       `json:"isDeleted"`
	CreatedAt      time.Time `json:"createdAt" sql:"DEFAULT:current_timestamp"`
}
