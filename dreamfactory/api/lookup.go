package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// LookupCreate creates a new lookup
func (c *Client) LookupCreate(l types.LookupsRequest) (lr types.LookupsResponse, err error) {
	err = c.send("POST", "/api/v2/system/lookup", 201, l, &lr)
	return
}

// LookupRead gets lookup's information
func (c *Client) LookupRead(id string) (l types.Lookup, err error) {
	err = c.send("GET", "/api/v2/system/lookup/"+id, 200, nil, &l)
	return
}

// LookupUpdate changes lookup's information
func (c *Client) LookupUpdate(id string, l types.Lookup) error {
	return c.send("PATCH", "/api/v2/system/lookup/"+id, 200, l, nil)
}

// LookupDelete destroys an lookup
func (c *Client) LookupDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/lookup/"+id, 200, nil, nil)
}
