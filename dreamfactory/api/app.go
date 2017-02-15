package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// AppCreate creates a new app
func (c *Client) AppCreate(a types.AppsRequest) (ar types.AppsResponse, err error) {
	err = c.send("POST", "/api/v2/system/app", 201, a, &ar)
	return
}

// AppRead gets app's information
func (c *Client) AppRead(id string) (a types.App, err error) {
	err = c.send("GET", "/api/v2/system/app/"+id, 200, nil, &a)
	return
}

// AppUpdate changes app's information
func (c *Client) AppUpdate(id string, a types.App) error {
	return c.send("PATCH", "/api/v2/system/app/"+id, 200, a, nil)
}

// AppDelete destroys an app
func (c *Client) AppDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/app/"+id, 200, nil, nil)
}
