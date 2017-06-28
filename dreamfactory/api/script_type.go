package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// ScriptTypeList gets supported script types
func (c *Client) ScriptTypeList() (st types.ScriptTypes, err error) {
	err = c.send("GET", "/api/v2/system/script_type", 200, nil, &st)
	return
}
