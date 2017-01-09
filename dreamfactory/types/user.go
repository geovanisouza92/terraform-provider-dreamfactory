package types

type UsersRequest struct {
	Resource []User `json:"resource,omitempty"`
	Ids      []int  `json:"ids,omitempty"`
}

type UsersResponse struct {
	Resource []User `json:"resource,omitempty"`
}

type User struct {
	Id               int    `json:"id,omitempty"`
	Name             string `json:"name"`
	Username         string `json:"username,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email"`
	IsActive         bool   `json:"is_active,omitempty"`
	Phone            string `json:"phone,omitempty"`
	SecurityQuestion string `json:"security_question,omitempty"`
	SecurityAnswer   string `json:"security_answer,omitempty"`
	DefaultAppID     int    `json:"default_app_id,omitempty"`
	OauthProvider    string `json:"oauth_provider,omitempty"`
}
