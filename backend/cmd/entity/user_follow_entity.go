package entity

import "time"

type UserFollow struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FollowUserId   string    `json:"followUserId"`
	FollowedUserId string    `json:"followedUserId"`
	CreatedAt      time.Time `json:"createdAt"`
}
