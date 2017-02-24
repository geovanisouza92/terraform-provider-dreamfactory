package dreamfactory

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceLookup() *schema.Resource {
	return &schema.Resource{
		Create: resourceLookupCreate,
		Read:   resourceLookupRead,
		Update: resourceLookupUpdate,
		Delete: resourceLookupDelete,
		Exists: resourceLookupExists,
		Importer: &schema.ResourceImporter{
			State: resourceLookupImporter,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceLookupCreate(d *schema.ResourceData, c interface{}) error {
	lr := types.LookupFromResourceData(d)
	l, err := c.(*api.Client).LookupCreate(types.LookupsRequest{Resource: []types.Lookup{lr}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(l.Resource[0].ID))
	return resourceLookupRead(d, c)
}

func resourceLookupRead(d *schema.ResourceData, c interface{}) error {
	l, err := c.(*api.Client).LookupRead(d.Id())
	if err != nil {
		return err
	}
	return l.FillResourceData(d)
}

func resourceLookupUpdate(d *schema.ResourceData, c interface{}) error {
	l := types.LookupFromResourceData(d)
	return c.(*api.Client).LookupUpdate(d.Id(), l)
}

func resourceLookupDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).LookupDelete(d.Id())
}

func resourceLookupExists(d *schema.ResourceData, c interface{}) (bool, error) {
	l, err := c.(*api.Client).LookupRead(d.Id())
	return l.ID > 0, err
}

func resourceLookupImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	l, err := c.(*api.Client).LookupRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = l.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
