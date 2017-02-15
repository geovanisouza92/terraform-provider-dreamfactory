package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// CorsCreate creates a new cors
func (c *Client) CorsCreate(crq types.CorsRequest) (crs types.CorsResponse, err error) {
	err = c.send("POST", "/api/v2/system/cors", 201, crq, &crs)
	return
}

// CorsRead gets cors's information
func (c *Client) CorsRead(id string) (co types.Cors, err error) {
	err = c.send("GET", "/api/v2/system/cors/"+id, 200, nil, &co)
	return
}

// CorsUpdate changes cors's information
func (c *Client) CorsUpdate(id string, co types.Cors) error {
	return c.send("PATCH", "/api/v2/system/cors/"+id, 200, co, nil)
}

// CorsDelete destroys an cors
func (c *Client) CorsDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/cors/"+id, 200, nil, nil)
}
