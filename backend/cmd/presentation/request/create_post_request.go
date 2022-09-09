package request

type CreatePostRequest struct {
	ID             string `json:"id" bson:"id"`
	Title          string `json:"title" bson:"title"`
	Body           string `json:"body" bson:"body"`
	PostedBy       string `json:"postedBy" bson:"postedBy"`
	IsDeleted      int    `json:"isDeleted" bson:"isDeleted"`
	IsDraft        int    `json:"isDraft" bson:"isDraft"`
	OrganizationId string `json:"organizationId" bson:"organizationId"`
}
