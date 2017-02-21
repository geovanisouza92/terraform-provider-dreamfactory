package dreamfactory

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceEmailTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailTemplateCreate,
		Read:   resourceEmailTemplateRead,
		Update: resourceEmailTemplateUpdate,
		Delete: resourceEmailTemplateDelete,
		Exists: resourceEmailTemplateExists,
		Importer: &schema.ResourceImporter{
			State: resourceEmailTemplateImporter,
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
			"to": &schema.Schema{
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				// Set:
			},
			"cc": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				// Set:
			},
			"bcc": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				// Set:
			},
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"body_text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"body_html": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from": &schema.Schema{
				Type:     schema.TypeSet,
				MinItems: 1,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				// Set:
			},
			"from_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reply_to": &schema.Schema{
				Type:     schema.TypeSet,
				MinItems: 1,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				// Set:
			},
			"reply_to_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reply_to_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"defaults": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				// Set:
			},
		},
	}
}

func resourceEmailTemplateCreate(d *schema.ResourceData, c interface{}) error {
	etr := types.EmailTemplateFromResourceData(d)
	et, err := c.(*api.Client).EmailTemplateCreate(types.EmailTemplatesRequest{Resource: []types.EmailTemplate{etr}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(et.Resource[0].ID))
	return resourceEmailTemplateRead(d, c)
}

func resourceEmailTemplateRead(d *schema.ResourceData, c interface{}) error {
	et, err := c.(*api.Client).EmailTemplateRead(d.Id())
	if err != nil {
		return err
	}
	return et.FillResourceData(d)
}

func resourceEmailTemplateUpdate(d *schema.ResourceData, c interface{}) error {
	et := types.EmailTemplateFromResourceData(d)
	return c.(*api.Client).EmailTemplateUpdate(d.Id(), et)
}

func resourceEmailTemplateDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).EmailTemplateDelete(d.Id())
}

func resourceEmailTemplateExists(d *schema.ResourceData, c interface{}) (bool, error) {
	et, err := c.(*api.Client).EmailTemplateRead(d.Id())
	return et.ID > 0, err
}

func resourceEmailTemplateImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	et, err := c.(*api.Client).EmailTemplateRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = et.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
