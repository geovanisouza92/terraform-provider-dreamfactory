package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// ServicesRequest represents a bulk request for services
type ServicesRequest struct {
	Resource []Service `json:"resource,omitempty"`
	IDs      []int     `json:"ids,omitempty"`
}

// ServicesResponse represents a bulk response for services
type ServicesResponse struct {
	Resource []Service `json:"resource,omitempty"`
	Meta     Metadata  `json:"meta,omitempty"`
}

// Service represents an service in DreamFactory
type Service struct {
	ID          int                     `json:"id,omitempty"`
	Name        string                  `json:"name"`
	Label       string                  `json:"label,omitempty"`
	Description string                  `json:"description,omitempty"`
	IsActive    bool                    `json:"is_active"`
	Type        string                  `json:"type"`      // one of ServiceType
	Mutable     bool                    `json:"mutable"`   // Error when trying to change
	Deletable   bool                    `json:"deletable"` // Error when trying to delete
	Doc         ServiceDoc              `json:"doc"`
	Config      *map[string]interface{} `json:"config"`
}

// ServiceDoc holds the Swagger spec of the service, when needed
type ServiceDoc struct {
	Format  int    `json:"format"`  // 0 = JSON | 1 = YAML
	Content string `json:"content"` // Swagger doc
}

// ServiceFromResourceData create an Service instance from ResourceData
func ServiceFromResourceData(d *schema.ResourceData) (*Service, error) {
	c := d.Get("config").(map[string]interface{})
	s := Service{
		Name:        d.Get("name").(string),
		Label:       d.Get("label").(string),
		Description: d.Get("description").(string),
		IsActive:    d.Get("is_active").(bool),
		Type:        d.Get("type").(string),
		Mutable:     d.Get("mutable").(bool),
		Deletable:   d.Get("deletable").(bool),
		Doc: ServiceDoc{
			Format:  d.Get("doc.format").(int),
			Content: d.Get("doc.content").(string),
		},
		Config: &c,
	}
	// TODO: Tratar ServiceConfig
	return &s, nil
}

// FillResourceData fills ResourceData with Service information
func (l *Service) FillResourceData(d *schema.ResourceData) error {
	// TODO: Tratar ServiceConfig
	return firstError([]func() error{
		setOrError(d, "name", l.Name),
		setOrError(d, "label", l.Label),
		setOrError(d, "description", l.Description),
		setOrError(d, "is_active", l.IsActive),
		setOrError(d, "type", l.Type),
		setOrError(d, "mutable", l.Mutable),
		setOrError(d, "deletable", l.Deletable),
		setOrError(d, "doc.format", l.Doc.Format),
		setOrError(d, "doc.content", l.Doc.Content),
		setOrError(d, "config", l.Config),
	})
}
