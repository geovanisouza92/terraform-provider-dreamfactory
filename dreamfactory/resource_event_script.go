package dreamfactory

import (
	"errors"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceEventScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventScriptCreate,
		Read:   resourceEventScriptRead,
		Update: resourceEventScriptUpdate,
		Delete: resourceEventScriptDelete,
		Exists: resourceEventScriptExists,
		Importer: &schema.ResourceImporter{
			State: resourceEventScriptImporter,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					found := false
					for _, eventScriptType := range eventScriptTypes {
						if value == eventScriptType {
							found = true
							break
						}
					}
					if !found {
						errors = append(errors, errInvalidEventScriptType)
					}
					return
				},
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allow_event_modification": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

var (
	eventScriptTypes          = []string{"python", "nodejs", "php", "v8js"}
	errInvalidEventScriptType = errors.New(`Invalid dreamfactory_event_script type. Possible values are:

	- "python"
	- "nodejs"
	- "php"
	- "v8js"
`)
)

func resourceEventScriptCreate(d *schema.ResourceData, c interface{}) error {
	esr := types.EventScriptFromResourceData(d)
	es, err := c.(*api.Client).EventScriptCreate(types.EventScriptsRequest{Resource: []types.EventScript{esr}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(es.Resource[0].ID))
	return resourceEventScriptRead(d, c)
}

func resourceEventScriptRead(d *schema.ResourceData, c interface{}) error {
	es, err := c.(*api.Client).EventScriptRead(d.Id())
	if err != nil {
		return err
	}
	return es.FillResourceData(d)
}

func resourceEventScriptUpdate(d *schema.ResourceData, c interface{}) error {
	es := types.EventScriptFromResourceData(d)
	return c.(*api.Client).EventScriptUpdate(d.Id(), es)
}

func resourceEventScriptDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).EventScriptDelete(d.Id())
}

func resourceEventScriptExists(d *schema.ResourceData, c interface{}) (bool, error) {
	es, err := c.(*api.Client).EventScriptRead(d.Id())
	return es.ID > 0, err
}

func resourceEventScriptImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	es, err := c.(*api.Client).EventScriptRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = es.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
