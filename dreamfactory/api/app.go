package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func (c *Client) AppCreate(a types.AppsRequest) (ar types.AppsResponse, err error) {
	err = c.send("POST", "/api/v2/system/app", 201, a, &ar)
	return
}

func (c *Client) AppRead(id string) (a types.App, err error) {
	err = c.send("GET", "/api/v2/system/app/"+id, 200, nil, &a)
	return
}

func (c *Client) AppUpdate(id string, a types.App) error {
	return c.send("PATCH", "/api/v2/system/app/"+id, 200, a, nil)
}

func (c *Client) AppDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/app"+id, 200, nil, nil)
}

func (c *Client) AppExists(id string) error {
	return c.send("GET", "/api/v2/system/app"+id, 200, nil, nil)
}
