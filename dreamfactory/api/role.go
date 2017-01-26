package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// TODO: ?related=role_service_access_by_role_id,role_lookup_by_role_id

func (c *Client) RoleCreate(r types.RolesRequest) (rr types.RolesRequest, err error) {
	err = c.send("POST", "/api/v2/system/role", 201, r, &rr)
	return
}

func (c *Client) RoleRead(id string) (r types.Role, err error) {
	err = c.send("GET", "/api/v2/system/role/"+id, 200, nil, &r)
	return
}

func (c *Client) RoleUpdate(id string, r types.Role) error {
	return c.send("PATCH", "/api/v2/system/role/"+id, 200, r, nil)
}

func (c *Client) RoleDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/role/"+id, 200, nil, nil)
}

func (c *Client) RoleExists(id string) error {
	return c.send("GET", "/api/v2/system/role/"+id, 200, nil, nil)
}
