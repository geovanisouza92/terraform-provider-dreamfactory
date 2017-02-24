package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// EventScriptsRequest represents a bulk request for event scripts
type EventScriptsRequest struct {
	Resource []EventScript `json:"resource,omitempty"`
	IDs      []int         `json:"ids,omitempty"`
}

// EventScriptsResponse represents a bulk response for event scripts
type EventScriptsResponse struct {
	Resource []EventScript `json:"resource,omitempty"`
	Meta     Metadata      `json:"meta,omitempty"`
}

// EventScript represents an event script in DreamFactory
type EventScript struct {
	ID                     int    `json:"id,omitempty"`
	Name                   string `json:"name"`
	Type                   string `json:"type"`
	Content                string `json:"content,omitempty"`
	Config                 string `json:"config,omitempty"`
	IsActive               bool   `json:"is_active,omitempty"`
	AllowEventModification bool   `json:"allow_event_modification,omitempty"`
}

// EventScriptFromResourceData create an EventScript instance from ResourceData
func EventScriptFromResourceData(d *schema.ResourceData) EventScript {
	return EventScript{
		Name:                   d.Get("name").(string),
		Type:                   d.Get("type").(string),
		Content:                d.Get("content").(string),
		Config:                 d.Get("config").(string),
		IsActive:               d.Get("is_active").(bool),
		AllowEventModification: d.Get("allow_event_modification").(bool),
	}
}

// FillResourceData fills ResourceData with EventScript information
func (es *EventScript) FillResourceData(d *schema.ResourceData) error {
	return firstError([]func() error{
		setOrError(d, "name", es.Name),
		setOrError(d, "type", es.Type),
		setOrError(d, "content", es.Content),
		setOrError(d, "config", es.Config),
		setOrError(d, "is_active", es.IsActive),
		setOrError(d, "allow_event_modification", es.AllowEventModification),
	})
}
