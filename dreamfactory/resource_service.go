package dreamfactory

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

func resourceService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,
		Exists: resourceServiceExists,
		Importer: &schema.ResourceImporter{
			State: resourceServiceImporter,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errs []error) {
					value := v.(string)
					if _, found := serviceTypes[value]; !found {
						errs = append(errs, errInvalidServiceType())
					}
					return
				},
			},
			"doc": &schema.Schema{
				Type:     schema.TypeList,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: func(v interface{}, k string) (ws []string, errs []error) {
								if value, ok := v.(string); !ok || (value != "json" && value != "yaml") {
									errs = append(errs, errors.New(`doc.format must be one of: "json" or "yaml"`))
								}
								return
							},
						},
						"content": &schema.Schema{
							Type: schema.TypeString,
							// TODO: ValidateFunc?
						},
					},
				},
			},
		},
	}
}

var (
	serviceTypes          = map[string]types.ServiceType{}
	errInvalidServiceType = func() error {
		types := []string{}
		for t := range serviceTypes {
			types = append(types, t)
		}
		options := strings.Join(types, "\n")
		return fmt.Errorf("Invalid dreamfactory_service type. Possible values are:\n\n%s", options)
	}
)

func resourceServiceInit(api *api.Client) error {
	sts, err := api.ServiceTypeList()
	if err != nil {
		return err
	}

	for _, t := range sts.Resource {
		serviceTypes[t.Name] = t
	}

	return nil
}

func resourceServiceCreate(d *schema.ResourceData, c interface{}) error {
	sr, err := types.ServiceFromResourceData(d)
	if err != nil {
		return err
	}

	if err = serviceTypes[sr.Type].Validate(*sr.Config); err != nil {
		return err
	}

	s, err := c.(*api.Client).ServiceCreate(types.ServicesRequest{Resource: []types.Service{*sr}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(s.Resource[0].ID))
	return resourceServiceRead(d, c)
}

func resourceServiceRead(d *schema.ResourceData, c interface{}) error {
	s, err := c.(*api.Client).ServiceRead(d.Id())
	if err != nil {
		return err
	}
	return s.FillResourceData(d)
}

func resourceServiceUpdate(d *schema.ResourceData, c interface{}) error {
	s, err := types.ServiceFromResourceData(d)
	if err != nil {
		return err
	}

	if err = serviceTypes[s.Type].Validate(*s.Config); err != nil {
		return err
	}

	return c.(*api.Client).ServiceUpdate(d.Id(), *s)
}

func resourceServiceDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).ServiceDelete(d.Id())
}

func resourceServiceExists(d *schema.ResourceData, c interface{}) (bool, error) {
	s, err := c.(*api.Client).ServiceRead(d.Id())
	return s.ID > 0, err
}

func resourceServiceImporter(d *schema.ResourceData, c interface{}) ([]*schema.ResourceData, error) {
	items := []*schema.ResourceData{}

	s, err := c.(*api.Client).ServiceRead(d.Id())
	if err != nil {
		return items, err
	}

	if err = s.FillResourceData(d); err != nil {
		return items, err
	}
	items = append(items, d)

	return items, nil
}
