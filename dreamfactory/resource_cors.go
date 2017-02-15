package dreamfactory

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceCors() *schema.Resource {
	return &schema.Resource{
		Create: resourceCorsCreate,
		Read:   resourceCorsRead,
		Update: resourceCorsUpdate,
		Delete: resourceCorsDelete,
		Exists: resourceCorsExists,
		Importer: &schema.ResourceImporter{
			State: resourceCorsImporter,
		},
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"origin": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: types.MethodHashFunc,
			},
			"max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceCorsCreate(d *schema.ResourceData, c interface{}) error {
	cr := types.CorsFromResourceData(d)
	co, err := c.(*api.Client).CorsCreate(types.CorsRequest{Resource: []types.Cors{cr}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(co.Resource[0].ID))
	return resourceCorsRead(d, c)
}

func resourceCorsRead(d *schema.ResourceData, c interface{}) error {
	co, err := c.(*api.Client).CorsRead(d.Id())
	if err != nil {
		return err
	}
	return co.FillResourceData(d)
}

func resourceCorsUpdate(d *schema.ResourceData, c interface{}) error {
	co := types.CorsFromResourceData(d)
	return c.(*api.Client).CorsUpdate(d.Id(), co)
}

func resourceCorsDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).CorsDelete(d.Id())
}

func resourceCorsExists(d *schema.ResourceData, c interface{}) (bool, error) {
	co, err := c.(*api.Client).CorsRead(d.Id())
	return co.ID > 0, err
}

func resourceCorsImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	co, err := c.(*api.Client).CorsRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = co.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
