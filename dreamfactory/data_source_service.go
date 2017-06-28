package dreamfactory

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
)

func dataSourceService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServiceRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			/*
				"doc": &schema.Schema{
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"format": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
							"content": &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
			*/
		},
	}
}

func dataSourceServiceRead(d *schema.ResourceData, c interface{}) error {
	rawFilters, ok := d.GetOk("filter")
	if !ok {
		return errFilter
	}
	filters := buildDataSourceFilters(rawFilters.(*schema.Set))

	s, err := c.(*api.Client).ServiceLoad(filters)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(s.ID))

	return s.FillResourceData(d)
}
