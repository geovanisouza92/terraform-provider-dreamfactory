package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// RoleCreate creates a new role
func (c *Client) RoleCreate(r types.RolesRequest) (rr types.RolesRequest, err error) {
	err = c.send("POST", "/api/v2/system/role?fields=*&related=role_service_access_by_role_id,role_lookup_by_role_id", 201, r, &rr)
	return
}

// RoleRead gets role's information
func (c *Client) RoleRead(id string) (r types.Role, err error) {
	err = c.send("GET", "/api/v2/system/role/"+id+"?fields=*&related=role_service_access_by_role_id,role_lookup_by_role_id", 200, nil, &r)
	return
}

// RoleUpdate changes role's information
func (c *Client) RoleUpdate(id string, r *types.Role) error {
	return c.send("PUT", "/api/v2/system/role/"+id+"?fields=*&related=role_service_access_by_role_id,role_lookup_by_role_id", 200, *r, r)
}

// RoleDelete destroys a role
func (c *Client) RoleDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/role/"+id+"?fields=*&related=role_service_access_by_role_id,role_lookup_by_role_id", 200, nil, nil)
}

// RoleExists checks if a role exists
func (c *Client) RoleExists(id string) error {
	return c.send("GET", "/api/v2/system/role/"+id+"?fields=*&related=role_service_access_by_role_id,role_lookup_by_role_id", 200, nil, nil)
}
