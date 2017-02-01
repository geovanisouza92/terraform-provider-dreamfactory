package dreamfactory

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Exists: resourceUserExists,
		Importer: &schema.ResourceImporter{
			State: resourceUserImporter,
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
			"security_answer": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, c interface{}) error {
	ur, err := types.UserFromResourceData(d)
	if err != nil {
		return err
	}
	u, err := c.(*api.Client).UserCreate(types.UsersRequest{Resource: []types.User{*ur}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(u.Resource[0].ID))
	return resourceUserRead(d, c)
}

func resourceUserRead(d *schema.ResourceData, c interface{}) error {
	u, err := c.(*api.Client).UserRead(d.Id())
	if err != nil {
		return err
	}
	return u.FillResourceData(d)
}

func resourceUserUpdate(d *schema.ResourceData, c interface{}) error {
	u, err := types.UserFromResourceData(d)
	if err != nil {
		return err
	}
	if err := c.(*api.Client).UserUpdate(d.Id(), u); err != nil {
		return err
	}
	return u.FillResourceData(d)
}

func resourceUserDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).UserDelete(d.Id())
}

func resourceUserExists(d *schema.ResourceData, c interface{}) (bool, error) {
	// FIXME: Possible bug
	err := c.(*api.Client).UserExists(d.Id())
	return err == nil, err
}

func resourceUserImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	u, err := c.(*api.Client).UserRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = u.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
