package request

type EmailSignUpRequest struct {
	Id       string `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
