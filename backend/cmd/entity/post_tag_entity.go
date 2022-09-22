package entity

type PostTag struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	PostId string `json:"postId"`
	TagId  string `json:"tagId"`
}
