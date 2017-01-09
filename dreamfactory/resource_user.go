package dreamfactory

import (
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"

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
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, c interface{}) error {
	ur := types.User{
		Name:             d.Get("name").(string),
		Username:         d.Get("username").(string),
		FirstName:        d.Get("first_name").(string),
		LastName:         d.Get("last_name").(string),
		Email:            d.Get("email").(string),
		IsActive:         d.Get("is_active").(bool),
		Phone:            d.Get("phone").(string),
		SecurityQuestion: d.Get("security_question").(string),
		SecurityAnswer:   d.Get("security_answer").(string),
		DefaultAppID:     d.Get("default_app_id").(int),
		OauthProvider:    d.Get("oauth_provider").(string),
	}
	u, err := c.(*api.Client).UserCreate(types.UsersRequest{Resource: []types.User{ur}})
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(u.Resource[0].Id))
	return nil
}

func resourceUserRead(d *schema.ResourceData, c interface{}) error {
	u, err := c.(*api.Client).UserRead(d.Id())
	d.Set("name", u.Name)
	d.Set("username", u.Username)
	d.Set("first_name", u.FirstName)
	d.Set("last_name", u.LastName)
	d.Set("email", u.Email)
	d.Set("is_active", u.IsActive)
	d.Set("phone", u.Phone)
	d.Set("security_question", u.SecurityQuestion)
	d.Set("security_answer", u.SecurityAnswer)
	d.Set("default_app_id", u.DefaultAppID)
	d.Set("oauth_provider", u.OauthProvider)
	return err
}

func resourceUserUpdate(d *schema.ResourceData, c interface{}) error {
	ur := types.User{
		Name:             d.Get("name").(string),
		Username:         d.Get("username").(string),
		FirstName:        d.Get("first_name").(string),
		LastName:         d.Get("last_name").(string),
		Email:            d.Get("email").(string),
		IsActive:         d.Get("is_active").(bool),
		Phone:            d.Get("phone").(string),
		SecurityQuestion: d.Get("security_question").(string),
		SecurityAnswer:   d.Get("security_answer").(string),
		DefaultAppID:     d.Get("default_app_id").(int),
		OauthProvider:    d.Get("oauth_provider").(string),
	}
	return c.(*api.Client).UserUpdate(d.Id(), ur)
}

func resourceUserDelete(d *schema.ResourceData, c interface{}) error {
	return c.(*api.Client).UserDelete(d.Id())
}

func resourceUserExists(d *schema.ResourceData, c interface{}) (bool, error) {
	err := c.(*api.Client).UserExists(d.Id())
	return err == nil, err
}
