package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// AppsRequest represents a bulk request for apps
type AppsRequest struct {
	Resource []App `json:"resource,omitempty"`
	IDs      []int `json:"ids,omitempty"`
}

// AppsResponse represents a bulk response for apps
type AppsResponse struct {
	Resource []App    `json:"resource,omitempty"`
	Meta     Metadata `json:"meta,omitempty"`
}

// App represents an app in DreamFactory
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
	AllowFullscreenToggle bool   `json:"allow_fullscreen_toggle,omitempty"`
	RoleID                int    `json:"role_id,omitempty"`
	LaunchURL             string `json:"launch_url,omitempty"`
}

var (
	appValuesToInt = map[string]int{
		"no_storage":   0,
		"provisioned":  1,
		"remote":       2,
		"on_webserver": 3,
	}
	appValuesToString = map[int]string{
		0: "no_storage",
		1: "provisioned",
		2: "remote",
		3: "on_webserver",
	}
)

// AppFromResourceData create an App instance from ResourceData
func AppFromResourceData(d *schema.ResourceData) App {
	apiKey, _ := d.Get("api_key").(string)
	return App{
		Name:                  d.Get("name").(string),
		APIKey:                apiKey,
		Description:           d.Get("description").(string),
		IsActive:              d.Get("is_active").(bool),
		Type:                  appValuesToInt[d.Get("type").(string)],
		Path:                  d.Get("path").(string),
		URL:                   d.Get("url").(string),
		StorageServiceID:      d.Get("storage_service_id").(int),
		StorageContainer:      d.Get("storage_container").(string),
		AllowFullscreenToggle: d.Get("allow_fullscreen_toggle").(bool),
		RoleID:                d.Get("role_id").(int),
		LaunchURL:             d.Get("launch_url").(string),
	}
}

// FillResourceData fills ResourceData with App information
func (a *App) FillResourceData(d *schema.ResourceData) error {
	d.Set("name", a.Name)
	d.Set("api_key", a.APIKey)
	d.Set("description", a.Description)
	d.Set("is_active", a.IsActive)
	d.Set("type", appValuesToString[a.Type])
	d.Set("path", a.Path)
	if a.Type == appValuesToInt["remote"] {
		d.Set("url", a.LaunchURL)
	} else {
		d.Set("url", a.URL)
	}
	d.Set("storage_service_id", a.StorageServiceID)
	d.Set("storage_container", a.StorageContainer)
	d.Set("allow_fullscreen_toggle", a.AllowFullscreenToggle)
	d.Set("role_id", a.RoleID)
	d.Set("launch_url", a.LaunchURL)

	return nil
}
