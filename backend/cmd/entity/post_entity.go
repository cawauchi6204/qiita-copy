package entity

import "time"

type Post struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	PostedBy       string    `json:"postedBy"`
	OrganizationId string    `json:"organizationId"`
	IsDraft        int       `json:"isDraft"`
	IsDeleted      int       `json:"isDeleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
