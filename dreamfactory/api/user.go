package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// UserCreate creates a new user
func (c *Client) UserCreate(u types.UsersRequest) (ur types.UsersResponse, err error) {
	err = c.send("POST", "/api/v2/system/user?fields=*&related=user_lookup_by_user_id,user_to_app_to_role_by_user_id", 201, u, &ur)
	return
}

// UserRead gets user's information
func (c *Client) UserRead(id string) (u types.User, err error) {
	err = c.send("GET", "/api/v2/system/user/"+id+"?fields=*&related=user_lookup_by_user_id,user_to_app_to_role_by_user_id", 200, nil, &u)
	return
}

// UserUpdate changes user's information
func (c *Client) UserUpdate(id string, u *types.User) error {
	return c.send("PUT", "/api/v2/system/user/"+id+"?fields=*&related=user_lookup_by_user_id,user_to_app_to_role_by_user_id", 200, *u, u)
}

// UserDelete destroys a user
func (c *Client) UserDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/user/"+id+"?fields=*&related=user_lookup_by_user_id,user_to_app_to_role_by_user_id", 200, nil, nil)
}

// UserExists checks if a user exists
func (c *Client) UserExists(id string) error {
	return c.send("GET", "/api/v2/system/user/"+id+"?fields=*&related=user_lookup_by_user_id,user_to_app_to_role_by_user_id", 200, nil, nil)
}
