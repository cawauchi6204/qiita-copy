package request

type UpdateLikeRequest struct {
	PostId string `json:"postId"`
	UserId string `json:"userId"`
}
