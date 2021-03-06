package api

import (
	"net/url"
	"strings"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// AdminCreate creates a new admin
func (c *Client) AdminCreate(u types.AdminsRequest) (ur types.AdminsResponse, err error) {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id"}, ","))
	err = c.send("POST", "/api/v2/system/admin?"+v.Encode(), 201, u, &ur)
	return
}

// AdminRead gets admin's information
func (c *Client) AdminRead(id string) (u types.Admin, err error) {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id"}, ","))
	err = c.send("GET", "/api/v2/system/admin/"+id+"?"+v.Encode(), 200, nil, &u)
	return
}

// AdminUpdate changes admin's information
func (c *Client) AdminUpdate(id string, u *types.Admin) error {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id"}, ","))
	return c.send("PUT", "/api/v2/system/admin/"+id+"?"+v.Encode(), 200, *u, u)
}

// AdminDelete destroys a admin
func (c *Client) AdminDelete(id string) error {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"user_lookup_by_user_id"}, ","))
	return c.send("DELETE", "/api/v2/system/admin/"+id+"?"+v.Encode(), 200, nil, nil)
}
