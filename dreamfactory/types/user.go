package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type UsersRequest struct {
	Resource []User `json:"resource,omitempty"`
	Ids      []int  `json:"ids,omitempty"`
}

type UsersResponse struct {
	Resource []User   `json:"resource,omitempty"`
	Meta     Metadata `json:"meta,omitempty"`
}

type User struct {
	ID               int    `json:"id,omitempty"`
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

func UserFromResourceData(d *schema.ResourceData) User {
	return User{
		Name:             d.Get("name").(string),
		Username:         d.Get("username").(string),
		FirstName:        d.Get("first_name").(string),
		LastName:         d.Get("last_name").(string),
		Email:            d.Get("email").(string),
		IsActive:         d.Get("is_active").(bool),
		Phone:            d.Get("phone").(string),
		SecurityQuestion: d.Get("security_question").(string),
		SecurityAnswer:   d.Get("security_answer").(string),
		DefaultAppID:     d.Get("default_app_id").(int),
		OauthProvider:    d.Get("oauth_provider").(string),
	}
}

func (u *User) FillResourceData(d *schema.ResourceData) error {
	d.Set("name", u.Name)
	d.Set("username", u.Username)
	d.Set("first_name", u.FirstName)
	d.Set("last_name", u.LastName)
	d.Set("email", u.Email)
	d.Set("is_active", u.IsActive)
	d.Set("phone", u.Phone)
	d.Set("security_question", u.SecurityQuestion)
	d.Set("security_answer", u.SecurityAnswer)
	d.Set("default_app_id", u.DefaultAppID)
	d.Set("oauth_provider", u.OauthProvider)

	return nil
}
