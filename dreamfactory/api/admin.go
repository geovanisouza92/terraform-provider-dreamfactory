package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// AdminCreate creates a new admin
func (c *Client) AdminCreate(u types.AdminsRequest) (ur types.AdminsResponse, err error) {
	err = c.send("POST", "/api/v2/system/admin?fields=*&related=user_lookup_by_user_id", 201, u, &ur)
	return
}

// AdminRead gets admin's information
func (c *Client) AdminRead(id string) (u types.Admin, err error) {
	err = c.send("GET", "/api/v2/system/admin/"+id+"?fields=*&related=user_lookup_by_user_id", 200, nil, &u)
	return
}

// AdminUpdate changes admin's information
func (c *Client) AdminUpdate(id string, u *types.Admin) error {
	return c.send("PUT", "/api/v2/system/admin/"+id+"?fields=*&related=user_lookup_by_user_id", 200, *u, u)
}

// AdminDelete destroys a admin
func (c *Client) AdminDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/admin/"+id+"?fields=*&related=user_lookup_by_user_id", 200, nil, nil)
}

// AdminExists checks if a admin exists
func (c *Client) AdminExists(id string) error {
	return c.send("GET", "/api/v2/system/admin/"+id+"?fields=*&related=user_lookup_by_user_id", 200, nil, nil)
}
