package request

type CreateTagRequest struct {
	ID     string `json:"id" bson:"id"`
	ImgUrl string `json:"imgUrl" bson:"imgUrl"`
	PostId string `json:"postId" bson:"postId"`
}
