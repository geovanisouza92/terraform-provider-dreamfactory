package dreamfactory

import (
	"errors"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminCreate,
		Read:   resourceAdminRead,
		Update: resourceAdminUpdate,
		Delete: resourceAdminDelete,
		Exists: resourceAdminExists,
		Importer: &schema.ResourceImporter{
			State: resourceAdminImporter,
		},
		// TODO: Add validations, defaults and computed props
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_question": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"confirm_code": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"default_app_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oauth_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"confirmed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"expired": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"lookup": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"private": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAdminCreate(d *schema.ResourceData, c interface{}) error {
	ar, err := types.AdminFromResourceData(d)
	if err != nil {
		return err
	}
	a, err := c.(*api.Client).AdminCreate(types.AdminsRequest{Resource: []types.Admin{*ar}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(a.Resource[0].ID))
	return resourceAdminRead(d, c)
}

func resourceAdminRead(d *schema.ResourceData, c interface{}) error {
	a, err := c.(*api.Client).AdminRead(d.Id())
	if err != nil {
		return err
	}
	return a.FillResourceData(d)
}

func resourceAdminUpdate(d *schema.ResourceData, c interface{}) error {
	api := c.(*api.Client)

	actual, err := api.AdminRead(d.Id())
	if err != nil {
		return errors.New("Could not read admin from remote: " + err.Error())
	}

	if err = actual.UpdateMissingResourceData(d); err != nil {
		return err
	}

	a, err := types.AdminFromResourceData(d)
	if err != nil {
		return err
	}
	if err := api.AdminUpdate(d.Id(), a); err != nil {
		return err
	}
	return a.FillResourceData(d)
}

func resourceAdminDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).AdminDelete(d.Id())
}

func resourceAdminExists(d *schema.ResourceData, c interface{}) (bool, error) {
	a, err := c.(*api.Client).AdminRead(d.Id())
	return a.ID > 0, err
}

func resourceAdminImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	a, err := c.(*api.Client).AdminRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = a.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
