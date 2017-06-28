package dreamfactory

import (
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/api"
)

// Provider creates a Terraform Provider configuration
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dreamfactory_admin":          resourceAdmin(),
			"dreamfactory_app":            resourceApp(),
			"dreamfactory_cors":           resourceCors(),
			"dreamfactory_email_template": resourceEmailTemplate(),
			"dreamfactory_event_script":   resourceEventScript(),
			"dreamfactory_lookup":         resourceLookup(),
			"dreamfactory_role":           resourceRole(),
			"dreamfactory_service":        resourceService(),
			"dreamfactory_user":           resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"dreamfactory_service": dataSourceService(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	c, err := api.New(endpoint, email, password, &http.Client{})
	if err != nil {
		return nil, err
	}
	if err = resourceEventScriptInit(c); err != nil {
		return nil, err
	}
	if err = resourceServiceInit(c); err != nil {
		return nil, err
	}
	return c, nil
}
