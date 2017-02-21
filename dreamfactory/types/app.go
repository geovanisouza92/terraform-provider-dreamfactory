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
	var url string
	if a.Type == appValuesToInt["remote"] {
		url = a.LaunchURL
	} else {
		url = a.URL
	}

	return firstError([]func() error{
		setOrError(d, "name", a.Name),
		setOrError(d, "api_key", a.APIKey),
		setOrError(d, "description", a.Description),
		setOrError(d, "is_active", a.IsActive),
		setOrError(d, "type", appValuesToString[a.Type]),
		setOrError(d, "path", a.Path),
		setOrError(d, "url", url),
		setOrError(d, "storage_service_id", a.StorageServiceID),
		setOrError(d, "storage_container", a.StorageContainer),
		setOrError(d, "allow_fullscreen_toggle", a.AllowFullscreenToggle),
		setOrError(d, "role_id", a.RoleID),
		setOrError(d, "launch_url", a.LaunchURL),
	})
}
