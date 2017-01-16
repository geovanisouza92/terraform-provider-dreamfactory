package dreamfactory

import (
	"errors"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,
		Exists: resourceAppExists,
		Importer: &schema.ResourceImporter{
			State: resourceAppImporter,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					found := false
					for _, appType := range appTypes {
						if value == appType {
							found = true
							break
						}
					}
					if !found {
						errors = append(errors, errInvalidAppType)
					}
					return
				},
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storage_service_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"storage_container": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_fullscreen_toggle": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"launch_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

var (
	appTypes          = []string{"no_storage", "provisioned", "remote", "on_webserver"}
	errInvalidAppType = errors.New(`Invalid dreafactory_app type. Possible values are:

	"no_storage":   No Storage Required - remote device, client, or desktop.
	"provisioned":  On a provisioned file storage service.
	"remote":       On a remote URL.
	"on_webserver": On this web server.
`)
)

func resourceAppCreate(d *schema.ResourceData, c interface{}) error {
	ar := types.AppFromResourceData(d)
	a, err := c.(*api.Client).AppCreate(types.AppsRequest{Resource: []types.App{ar}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(a.Resource[0].ID))
	return nil
}

func resourceAppRead(d *schema.ResourceData, c interface{}) error {
	a, err := c.(*api.Client).AppRead(d.Id())
	if err != nil {
		return err
	}
	a.FillResourceData(d)
	return nil
}

func resourceAppUpdate(d *schema.ResourceData, c interface{}) error {
	a := types.AppFromResourceData(d)
	return c.(*api.Client).AppUpdate(d.Id(), a)
}

func resourceAppDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).AppDelete(d.Id())
}

func resourceAppExists(d *schema.ResourceData, c interface{}) (bool, error) {
	// FIXME: Possible bug
	err := c.(*api.Client).AppExists(d.Id())
	return err == nil, err
}

func resourceAppImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	a, err := c.(*api.Client).AppRead(d.Id())
	if err != nil {
		return items, err
	}

	a.FillResourceData(d)
	items = append(items, d)

	return items, nil
}
