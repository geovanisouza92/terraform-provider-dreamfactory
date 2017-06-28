package api

import (
	"net/url"
	"strings"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// RoleCreate creates a new role
func (c *Client) RoleCreate(r types.RolesRequest) (rr types.RolesRequest, err error) {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"role_service_access_by_role_id", "role_lookup_by_role_id"}, ","))
	err = c.send("POST", "/api/v2/system/role?"+v.Encode(), 201, r, &rr)
	return
}

// RoleRead gets role's information
func (c *Client) RoleRead(id string) (r types.Role, err error) {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"role_service_access_by_role_id", "role_lookup_by_role_id"}, ","))
	err = c.send("GET", "/api/v2/system/role/"+id+"?"+v.Encode(), 200, nil, &r)
	return
}

// RoleUpdate changes role's information
func (c *Client) RoleUpdate(id string, r *types.Role) error {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"role_service_access_by_role_id", "role_lookup_by_role_id"}, ","))
	return c.send("PUT", "/api/v2/system/role/"+id+"?"+v.Encode(), 200, *r, r)
}

// RoleDelete destroys a role
func (c *Client) RoleDelete(id string) error {
	v := url.Values{}
	v.Set("fields", "*")
	v.Set("related", strings.Join([]string{"role_service_access_by_role_id", "role_lookup_by_role_id"}, ","))
	return c.send("DELETE", "/api/v2/system/role/"+id+"?"+v.Encode(), 200, nil, nil)
}
