package api

import (
	"net/url"
	"strings"

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

// ServiceDelete destroys a service
func (c *Client) ServiceDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/service/"+id, 200, nil, nil)
}

// ServiceLoad load a service for data source
func (c *Client) ServiceLoad(filter []string) (s types.Service, err error) {
	sr := types.ServicesResponse{}

	v := url.Values{}
	v.Set("filter", strings.Join(filter, " "))
	err = c.send("GET", "/api/v2/system/service/?"+v.Encode(), 200, nil, &sr)
	if err != nil {
		return
	}

	if len(sr.Resource) > 0 {
		s = sr.Resource[0]
	}

	return
}
