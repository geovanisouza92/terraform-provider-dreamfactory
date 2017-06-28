package api

import (
	"net/url"
	"strings"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// UserCreate creates a new user
func (c *Client) UserCreate(u types.UsersRequest) (ur types.UsersResponse, err error) {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id", "user_to_app_to_role_by_user_id"}, ","))
	err = c.send("POST", "/api/v2/system/user?"+v.Encode(), 201, u, &ur)
	return
}

// UserRead gets user's information
func (c *Client) UserRead(id string) (u types.User, err error) {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id", "user_to_app_to_role_by_user_id"}, ","))
	err = c.send("GET", "/api/v2/system/user/"+id+"?"+v.Encode(), 200, nil, &u)
	return
}

// UserUpdate changes user's information
func (c *Client) UserUpdate(id string, u *types.User) error {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id", "user_to_app_to_role_by_user_id"}, ","))
	return c.send("PUT", "/api/v2/system/user/"+id+"?"+v.Encode(), 200, *u, u)
}

// UserDelete destroys a user
func (c *Client) UserDelete(id string) error {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id", "user_to_app_to_role_by_user_id"}, ","))
	return c.send("DELETE", "/api/v2/system/user/"+id+"?"+v.Encode(), 200, nil, nil)
}
