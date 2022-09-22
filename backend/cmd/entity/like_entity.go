package entity

import "time"

type Like struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PostId     string    `json:"postId"`
	LikeUserId string    `json:"likeUserId"`
	CreatedAt  time.Time `json:"createdAt"`
}
