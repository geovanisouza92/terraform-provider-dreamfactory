package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// ServiceCreate creates a new service
func (c *Client) ServiceCreate(s types.ServicesRequest) (sr types.ServicesResponse, err error) {
	err = c.send("POST", "/api/v2/system/service", 201, s, &sr)
	return
}

// ServiceRead gets service's information
func (c *Client) ServiceRead(id string) (s types.Service, err error) {
	err = c.send("GET", "/api/v2/system/service/"+id, 200, nil, &s)
	return
}

// ServiceUpdate changes service's information
func (c *Client) ServiceUpdate(id string, s types.Service) error {
	return c.send("PATCH", "/api/v2/system/service/"+id, 200, s, nil)
}

// ServiceDelete destroys an service
func (c *Client) ServiceDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/service/"+id, 200, nil, nil)
}
