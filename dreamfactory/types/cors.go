package types

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

// CorsRequest represents a bulk request for apps
type CorsRequest struct {
	Resource []Cors `json:"resource,omitempty"`
	IDs      []int  `json:"ids,omitempty"`
}

// CorsResponse represents a bulk response for apps
type CorsResponse struct {
	Resource []Cors   `json:"resource"`
	Meta     Metadata `json:"meta,omitempty"`
}

// Cors represents an cors configuration in DreamFactory
type Cors struct {
	ID      int      `json:"id,omitempty"`
	Path    string   `json:"path"`
	Origin  string   `json:"origin"`
	Header  string   `json:"header"`
	Method  []string `json:"method"`
	MaxAge  int      `json:"max_age"`
	Enabled bool     `json:"enabled,omitempty"`
}

var methodHash = map[string]int{
	"GET":    1,
	"POST":   2,
	"PUT":    3,
	"PATCH":  4,
	"DELETE": 5,
}

// CorsFromResourceData create an Cors instance from ResourceData
func CorsFromResourceData(d *schema.ResourceData) Cors {
	method := []string{}
	for _, m := range d.Get("method").(*schema.Set).List() {
		log.Printf("eita: %v\n", m)
		method = append(method, strings.ToUpper(m.(string)))
	}
	return Cors{
		Path:    d.Get("path").(string),
		Origin:  d.Get("origin").(string),
		Header:  d.Get("header").(string),
		Method:  method,
		MaxAge:  d.Get("max_age").(int),
		Enabled: d.Get("enabled").(bool),
	}
}

// FillResourceData fills ResourceData with Cors information
func (c *Cors) FillResourceData(d *schema.ResourceData) error {
	method := []interface{}{}
	for _, m := range c.Method {
		method = append(method, m)
	}
	return firstError([]func() error{
		setOrError(d, "path", c.Path),
		setOrError(d, "origin", c.Origin),
		setOrError(d, "header", c.Header),
		setOrError(d, "method", method),
		setOrError(d, "max_age", c.MaxAge),
		setOrError(d, "enabled", c.Enabled),
	})
}

// MethodHashFunc convert a value into a unique ID
func MethodHashFunc(v interface{}) int {
	s, ok := v.(string)
	if !ok {
		return 0
	}
	i, ok := methodHash[strings.ToUpper(s)]
	if !ok {
		return 0
	}
	return i
}
