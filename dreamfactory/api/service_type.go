package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// ServiceTypeList gets supported script types
func (c *Client) ServiceTypeList() (st types.ServiceTypesResponse, err error) {
	err = c.send("GET", "/api/v2/system/service_type", 200, nil, &st)
	return
}
