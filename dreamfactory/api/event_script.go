package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// EventScriptCreate creates a new event script
func (c *Client) EventScriptCreate(es types.EventScriptsRequest) (esr types.EventScriptsResponse, err error) {
	err = c.send("POST", "/api/v2/system/event_script", 201, es, &esr)
	return
}

// EventScriptRead gets event script's information
func (c *Client) EventScriptRead(id string) (es types.EventScript, err error) {
	err = c.send("GET", "/api/v2/system/event_script/"+id, 200, nil, &es)
	return
}

// EventScriptUpdate changes event script's information
func (c *Client) EventScriptUpdate(id string, es types.EventScript) error {
	return c.send("PATCH", "/api/v2/system/event_script/"+id, 200, es, nil)
}

// EventScriptDelete destroys an event script
func (c *Client) EventScriptDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/event_script/"+id, 200, nil, nil)
}
