package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type AppsRequest struct {
	Resource []App `json:"resource,omitempty"`
	IDs      []int `json:"ids,omitempty"`
}

type AppsResponse struct {
	Resource []App    `json:"resource,omitempty"`
	Meta     Metadata `json:"meta,omitempty"`
}

type App struct {
	ID                    int    `json:"id,omitempty"`
	Name                  string `json:"name"`
	APIKey                string `json:"api_key"`
	Description           string `json:"description,omitempty"`
	IsActive              bool   `json:"is_active,omitempty"`
	Type                  int    `json:"type,omitempty"`
	Path                  string `json:"path,omitempty"`
	URL                   string `json:"url,omitempty"`
	StorageServiceID      int    `json:"storage_service_id,omitempty"`
	StorageContainer      string `json:"storage_container,omitempty"`
	RequiresFullscreen    bool   `json:"requires_fullscreen,omitempty"`
	AllowFullscreenToggle bool   `json:"allow_fullscreen_toggle,omitempty"`
	ToggleLocation        string `json:"toggle_location,omitempty"`
	RoleID                int    `json:"role_id,omitempty"`
	LaunchURL             string `json:"launch_url,omitempty"`
}

func AppFromResourceData(d *schema.ResourceData) App {
	return App{
		Name:                  d.Get("name").(string),
		APIKey:                d.Get("api_key").(string),
		Description:           d.Get("description").(string),
		IsActive:              d.Get("is_active").(bool),
		Type:                  d.Get("type").(int),
		Path:                  d.Get("path").(string),
		URL:                   d.Get("url").(string),
		StorageServiceID:      d.Get("storage_service_id").(int),
		StorageContainer:      d.Get("storage_container").(string),
		RequiresFullscreen:    d.Get("requires_fullscreen").(bool),
		AllowFullscreenToggle: d.Get("allow_fullscreen_toggle").(bool),
		ToggleLocation:        d.Get("toggle_location").(string),
		RoleID:                d.Get("role_id").(int),
		LaunchURL:             d.Get("launch_url").(string),
	}
}

func (a *App) FillResourceData(d *schema.ResourceData) {
	d.Set("name", a.Name)
	d.Set("api_key", a.APIKey)
	d.Set("description", a.Description)
	d.Set("is_active", a.IsActive)
	d.Set("type", a.Type)
	d.Set("path", a.Path)
	d.Set("url", a.URL)
	d.Set("storage_service_id", a.StorageServiceID)
	d.Set("storage_container", a.StorageContainer)
	d.Set("requires_fullscreen", a.RequiresFullscreen)
	d.Set("allow_fullscreen_toggle", a.AllowFullscreenToggle)
	d.Set("toggle_location", a.ToggleLocation)
	d.Set("role_id", a.RoleID)
	d.Set("launch_url", a.LaunchURL)
}
