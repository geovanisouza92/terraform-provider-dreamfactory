package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// LookupsRequest represents a bulk request for lookups
type LookupsRequest struct {
	Resource []Lookup `json:"resource,omitempty"`
	IDs      []int    `json:"ids,omitempty"`
}

// LookupsResponse represents a bulk response for lookups
type LookupsResponse struct {
	Resource []Lookup `json:"resource,omitempty"`
	Meta     Metadata `json:"meta,omitempty"`
}

// Lookup represents an lookup in DreamFactory
type Lookup struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
}

// LookupFromResourceData create an Lookup instance from ResourceData
func LookupFromResourceData(d *schema.ResourceData) Lookup {
	return Lookup{
		Name:        d.Get("name").(string),
		Value:       d.Get("value").(string),
		Description: d.Get("description").(string),
		Private:     d.Get("private").(bool),
	}
}

// FillResourceData fills ResourceData with Lookup information
func (l *Lookup) FillResourceData(d *schema.ResourceData) error {
	return firstError([]func() error{
		setOrError(d, "name", l.Name),
		setOrError(d, "value", l.Value),
		setOrError(d, "description", l.Description),
		setOrError(d, "private", l.Private),
	})
}
