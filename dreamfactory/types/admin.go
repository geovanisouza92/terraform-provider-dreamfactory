package types

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

// AdminsRequest is a request object for Admin CRUD
type AdminsRequest struct {
	Resource []Admin `json:"resource,omitempty"`
	IDs      []int   `json:"ids,omitempty"`
}

// AdminsResponse is a response object for Admin CRUD
type AdminsResponse struct {
	Resource []Admin  `json:"resource,omitempty"`
	Meta     Metadata `json:"meta,omitempty"`
}

// Admin represents a DreamFactory admin
type Admin struct {
	ID               int           `json:"id,omitempty"`
	Name             string        `json:"name"`
	Username         string        `json:"username,omitempty"`
	FirstName        string        `json:"first_name,omitempty"`
	LastName         string        `json:"last_name,omitempty"`
	Email            string        `json:"email"`
	IsActive         bool          `json:"is_active,omitempty"`
	Phone            string        `json:"phone,omitempty"`
	SecurityQuestion string        `json:"security_question,omitempty"`
	ConfirmCode      string        `json:"confirm_code,omitempty"`
	DefaultAppID     int           `json:"default_app_id,omitempty"`
	OauthProvider    string        `json:"oauth_provider,omitempty"`
	Confirmed        bool          `json:"confirmed,omitempty"`
	Expired          bool          `json:"expired,omitempty"`
	Lookups          []AdminLookup `json:"user_lookup_by_user_id"`
}

// AdminLookup represents a DreamFactory admin lookup
type AdminLookup struct {
	ID          int    `json:"id,omitempty"`
	AdminID     *int   `json:"user_id"`
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	Private     bool   `json:"private,omitempty"`
	Description string `json:"description,omitempty"`
}

// AdminFromResourceData creates a Admin object from Terraform ResourceData
func AdminFromResourceData(d *schema.ResourceData) (*Admin, error) {
	id, err := strconv.Atoi(d.Id())
	if d.Id() != "" && err != nil {
		return nil, err
	}

	a := Admin{
		ID:               id,
		Name:             d.Get("name").(string),
		Username:         d.Get("username").(string),
		FirstName:        d.Get("first_name").(string),
		LastName:         d.Get("last_name").(string),
		Email:            d.Get("email").(string),
		IsActive:         d.Get("is_active").(bool),
		Phone:            d.Get("phone").(string),
		SecurityQuestion: d.Get("security_question").(string),
		ConfirmCode:      d.Get("confirm_code").(string),
		DefaultAppID:     d.Get("default_app_id").(int),
		OauthProvider:    d.Get("oauth_provider").(string),
		Confirmed:        d.Get("confirmed").(bool),
		Expired:          d.Get("expired").(bool),
	}

	a.Lookups = make([]AdminLookup, 0)
	for i := 0; i < d.Get("lookup.#").(int); i++ {
		prefix := fmt.Sprintf("lookup.%d.", i)

		adminID := &a.ID
		lookupID := d.Get(prefix + "id").(int)
		if lookupID < 0 {
			lookupID = lookupID * -1
			adminID = nil
		}

		a.Lookups = append(a.Lookups, AdminLookup{
			ID:          lookupID,
			AdminID:     adminID,
			Name:        d.Get(prefix + "name").(string),
			Value:       d.Get(prefix + "value").(string),
			Private:     d.Get(prefix + "private").(bool),
			Description: d.Get(prefix + "description").(string),
		})
	}

	return &a, nil
}

// FillResourceData copy information from the Admin to Terraform ResourceData
func (a *Admin) FillResourceData(d *schema.ResourceData) error {
	lookup := []map[string]interface{}{}
	for _, l := range a.Lookups {
		lookup = append(lookup, map[string]interface{}{
			"id":          l.ID,
			"name":        l.Name,
			"value":       l.Value,
			"private":     l.Private,
			"description": l.Description,
		})
	}

	return firstError([]func() error{
		setOrError(d, "id", a.ID),
		setOrError(d, "name", a.Name),
		setOrError(d, "username", a.Username),
		setOrError(d, "first_name", a.FirstName),
		setOrError(d, "last_name", a.LastName),
		setOrError(d, "email", a.Email),
		setOrError(d, "is_active", a.IsActive),
		setOrError(d, "phone", a.Phone),
		setOrError(d, "security_question", a.SecurityQuestion),
		setOrError(d, "confirm_code", a.ConfirmCode),
		setOrError(d, "default_app_id", a.DefaultAppID),
		setOrError(d, "oauth_provider", a.OauthProvider),
		setOrError(d, "confirmed", a.Confirmed),
		setOrError(d, "expired", a.Expired),
		setOrError(d, "lookup", lookup),
	})
}

// UpdateMissingResourceData checks for remote lookup that doesn't exist locally anymore, marking the admin with missing ones, to allow remote removal
func (a *Admin) UpdateMissingResourceData(d *schema.ResourceData) error {
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
		prefix := fmt.Sprintf("lookup.%d.", i)
		lookup = append(lookup, map[string]interface{}{
			"id":          d.Get(prefix + "id"),
			"name":        d.Get(prefix + "name"),
			"value":       d.Get(prefix + "value"),
			"private":     d.Get(prefix + "private"),
			"description": d.Get(prefix + "description"),
		})
	}

	// Loop through remote lookups, marking items that was removed locally
	for _, l := range a.Lookups {
		if i := find(l.ID); i == -1 {
			lookup = append(lookup, map[string]interface{}{
				"id":          l.ID * -1,
				"name":        l.Name,
				"value":       l.Value,
				"private":     l.Private,
				"description": l.Description,
			})
		}
	}

	return firstError([]func() error{
		setOrError(d, "lookup", lookup),
	})
}
