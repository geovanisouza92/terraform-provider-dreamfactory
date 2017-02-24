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
			"dreamfactory_admin": resourceAdmin(),
			// "dreamfactory_app_group"
			"dreamfactory_app":  resourceApp(),
			"dreamfactory_cors": resourceCors(),
			// "dreamfactory_custom" ???
			"dreamfactory_email_template": resourceEmailTemplate(),
			// "dreamfactory_event_script"
			"dreamfactory_lookup": resourceLookup(),
			"dreamfactory_role":   resourceRole(),
			// "dreamfactory_script"
			// "dreamfactory_service" (/system/service_type)
			"dreamfactory_user": resourceUser(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	return api.New(endpoint, email, password, &http.Client{})
}
