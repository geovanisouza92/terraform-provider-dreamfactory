package api

import (
	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// EmailTemplateCreate creates a new email template
func (c *Client) EmailTemplateCreate(a types.EmailTemplatesRequest) (ar types.EmailTemplatesResponse, err error) {
	err = c.send("POST", "/api/v2/system/email_template", 201, a, &ar)
	return
}

// EmailTemplateRead gets email template's information
func (c *Client) EmailTemplateRead(id string) (a types.EmailTemplate, err error) {
	err = c.send("GET", "/api/v2/system/email_template/"+id, 200, nil, &a)
	return
}

// EmailTemplateUpdate changes email template's information
func (c *Client) EmailTemplateUpdate(id string, a types.EmailTemplate) error {
	return c.send("PATCH", "/api/v2/system/email_template/"+id, 200, a, nil)
}

// EmailTemplateDelete destroys an email template
func (c *Client) EmailTemplateDelete(id string) error {
	return c.send("DELETE", "/api/v2/system/email_template/"+id, 200, nil, nil)
}
