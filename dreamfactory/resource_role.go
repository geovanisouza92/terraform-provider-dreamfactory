package dreamfactory

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,
		Exists: resourceRoleExists,
		Importer: &schema.ResourceImporter{
			State: resourceRoleImporter,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  true,
			},
			"access": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_id": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"component": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"verb_mask": &schema.Schema{
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							// TODO: Transformar string em número "máscara"
						},
						"requestor_mask": &schema.Schema{
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							// TODO: Transformar string em número "máscara"
						},
						"filters": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"operator": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Default:  "=",
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
								},
							},
						},
						"filter_op": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "AND",
						},
					},
				},
			},
			"lookup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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

func resourceRoleCreate(d *schema.ResourceData, c interface{}) error {
	rr := types.RoleFromResourceData(d)
	r, err := c.(*api.Client).RoleCreate(types.RolesRequest{Resource: []types.Role{rr}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(r.Resource[0].ID))
	return nil
}

func resourceRoleRead(d *schema.ResourceData, c interface{}) error {
	r, err := c.(*api.Client).RoleRead(d.Id())
	if err != nil {
		return err
	}
	return r.FillResourceData(d)
}

func resourceRoleUpdate(d *schema.ResourceData, c interface{}) error {
	r := types.RoleFromResourceData(d)
	return c.(*api.Client).RoleUpdate(d.Id(), r)
}

func resourceRoleDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).RoleDelete(d.Id())
}

func resourceRoleExists(d *schema.ResourceData, c interface{}) (bool, error) {
	// FIXME: Possible bug
	err := c.(*api.Client).RoleExists(d.Id())
	return err == nil, err
}

func resourceRoleImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	r, err := c.(*api.Client).RoleRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = r.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
