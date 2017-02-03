package types

import (
	"fmt"
	"strconv"

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
	ID               int          `json:"id,omitempty"`
	Name             string       `json:"name"`
	Username         string       `json:"username,omitempty"`
	FirstName        string       `json:"first_name,omitempty"`
	LastName         string       `json:"last_name,omitempty"`
	Email            string       `json:"email"`
	IsActive         bool         `json:"is_active,omitempty"`
	Phone            string       `json:"phone,omitempty"`
	SecurityQuestion string       `json:"security_question,omitempty"`
	SecurityAnswer   string       `json:"security_answer,omitempty"`
	DefaultAppID     int          `json:"default_app_id,omitempty"`
	OauthProvider    string       `json:"oauth_provider,omitempty"`
	Lookups          []UserLookup `json:"user_lookup_by_user_id"`
}

type UserLookup struct {
	ID      int    `json:"id,omitempty"`
	UserID  *int   `json:"user_id"`
	Name    string `json:"name"`
	Value   string `json:"value,omitempty"`
	Private bool   `json:"private,omitempty"`
}

func UserFromResourceData(d *schema.ResourceData) (*User, error) {
	id, err := strconv.Atoi(d.Id())
	if d.Id() != "" && err != nil {
		return nil, err
	}

	u := User{
		ID:               id,
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

	u.Lookups = make([]UserLookup, 0)
	for i := 0; i < d.Get("lookup.#").(int); i++ {
		prefix := fmt.Sprintf("lookup.%d.", i)
		userID := &u.ID
		lookupID := d.Get(prefix + "id").(int)
		if lookupID < 0 {
			lookupID = lookupID * -1
			userID = nil
		}
		u.Lookups = append(u.Lookups, UserLookup{
			ID:      lookupID,
			UserID:  userID,
			Name:    d.Get(prefix + "name").(string),
			Value:   d.Get(prefix + "value").(string),
			Private: d.Get(prefix + "private").(bool),
		})
	}

	return &u, nil
}

func (u *User) FillResourceData(d *schema.ResourceData) error {
	lookup := []map[string]interface{}{}
	for _, l := range u.Lookups {
		lookup = append(lookup, map[string]interface{}{
			"id":      l.ID,
			"name":    l.Name,
			"value":   l.Value,
			"private": l.Private,
		})
	}

	return firstError([]func() error{
		setOrError(d, "name", u.Name),
		setOrError(d, "username", u.Username),
		setOrError(d, "first_name", u.FirstName),
		setOrError(d, "last_name", u.LastName),
		setOrError(d, "email", u.Email),
		setOrError(d, "is_active", u.IsActive),
		setOrError(d, "phone", u.Phone),
		setOrError(d, "security_question", u.SecurityQuestion),
		setOrError(d, "security_answer", u.SecurityAnswer),
		setOrError(d, "default_app_id", u.DefaultAppID),
		setOrError(d, "oauth_provider", u.OauthProvider),
		setOrError(d, "lookup", lookup),
	})
}

func (u *User) UpdateMissingResourceData(d *schema.ResourceData) error {
	count := d.Get("lookup.#").(int)

	// Find for specific lookup ID
	find := func(id int) /* index */ int {
		for i := 0; i < count; i++ {
			if _id, ok := d.Get(fmt.Sprintf("lookup.%d.id", i)).(int); ok && id == _id {
				return i
			}
		}
		return -1
	}

	// Load existing items from ResourceData
	lookup := []map[string]interface{}{}
	for i := 0; i < count; i++ {
		lookup = append(lookup, map[string]interface{}{
			"id":      d.Get(fmt.Sprintf("lookup.%d.id", i)),
			"name":    d.Get(fmt.Sprintf("lookup.%d.name", i)),
			"value":   d.Get(fmt.Sprintf("lookup.%d.value", i)),
			"private": d.Get(fmt.Sprintf("lookup.%d.private", i)),
		})
	}

	// Loop through remote lookups, marking items that was removed locally
	for _, l := range u.Lookups {
		if i := find(l.ID); i == -1 {
			lookup = append(lookup, map[string]interface{}{
				"id":      l.ID * -1,
				"name":    l.Name,
				"value":   l.Value,
				"private": l.Private,
			})
		}
	}

	return firstError([]func() error{
		setOrError(d, "lookup", lookup),
	})
}
