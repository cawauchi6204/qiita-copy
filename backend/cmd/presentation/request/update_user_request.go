package request

type UpdateUserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Description   string `json:"description"`
	HpUrl         string `json:"hpUrl"`
	Location      string `json:"location"`
	GithubAccount string `json:"githubAccountId"`
}
