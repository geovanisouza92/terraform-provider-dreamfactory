package types

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// EmailTemplatesRequest represents a bulk request for apps
type EmailTemplatesRequest struct {
	Resource []EmailTemplate `json:"resource,omitempty"`
	IDs      []int           `json:"ids,omitempty"`
}

// EmailTemplatesResponse represents a bulk response for apps
type EmailTemplatesResponse struct {
	Resource []EmailTemplate `json:"resource,omitempty"`
	Meta     Metadata        `json:"meta,omitempty"`
}

// EmailTemplate represents an app in DreamFactory
type EmailTemplate struct {
	ID           int            `json:"id,omitempty"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	To           EmailAddresses `json:"to"`
	Cc           EmailAddresses `json:",omitempty,omitempty"`
	Bcc          EmailAddresses `json:"bcc,omitempty"`
	Subject      string         `json:"subject,omitempty"`
	BodyText     string         `json:"body_text,omitempty"`
	BodyHTML     string         `json:"body_html,omitempty"`
	From         EmailAddress   `json:"from"`
	FromName     string         `json:"from_name,omitempty"`
	FromEmail    string         `json:"from_email,omitempty"`
	ReplyTo      EmailAddress   `json:"reply_to,omitempty"`
	ReplyToName  string         `json:"reply_to_name,omitempty"`
	ReplyToEmail string         `json:"reply_to_email,omitempty"`
	Defaults     []string       `json:"defaults"`
}

// EmailAddress represents an email address
type EmailAddress struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// EmailAddresses is a collection of email addresses
type EmailAddresses []EmailAddress

// EmailTemplateFromResourceData create an EmailTemplate instance from ResourceData
func EmailTemplateFromResourceData(d *schema.ResourceData) EmailTemplate {
	et := EmailTemplate{
		ID:           d.Get("id").(int),
		Name:         d.Get("name").(string),
		Description:  d.Get("description").(string),
		Subject:      d.Get("subject").(string),
		BodyText:     d.Get("body_text").(string),
		BodyHTML:     d.Get("body_html").(string),
		FromName:     d.Get("from_name").(string),
		FromEmail:    d.Get("from_email").(string),
		ReplyToName:  d.Get("reply_to_name").(string),
		ReplyToEmail: d.Get("reply_to_email").(string),
	}

	et.To = getEmailAddresses(d, "to")
	et.Cc = getEmailAddresses(d, "cc")
	et.Bcc = getEmailAddresses(d, "bcc")
	if from := getEmailAddress(d, "from"); from != nil {
		et.From = *from
	}
	if replyTo := getEmailAddress(d, "reply_to"); replyTo != nil {
		et.ReplyTo = *replyTo
	}

	et.Defaults = make([]string, 0)
	for _, d := range d.Get("defaults").(*schema.Set).List() {
		et.Defaults = append(et.Defaults, d.(string))
	}

	return et
}

// FillResourceData fills ResourceData with EmailTemplate information
func (et *EmailTemplate) FillResourceData(d *schema.ResourceData) error {
	return firstError([]func() error{
		setOrError(d, "id", et.ID),
		setOrError(d, "name", et.Name),
		setOrError(d, "description", et.Description),
		setOrError(d, "to", et.To.toSliceMap()),
		setOrError(d, "cc", et.Cc.toSliceMap()),
		setOrError(d, "bcc", et.Bcc.toSliceMap()),
		setOrError(d, "subject", et.Subject),
		setOrError(d, "body_text", et.BodyText),
		setOrError(d, "body_html", et.BodyHTML),
		setOrError(d, "from", et.From.toMap()),
		setOrError(d, "from_name", et.FromName),
		setOrError(d, "from_email", et.FromEmail),
		setOrError(d, "reply_to", et.ReplyTo.toMap()),
		setOrError(d, "reply_to_name", et.ReplyToName),
		setOrError(d, "reply_to_email", et.ReplyToEmail),
		setOrError(d, "defaults", et.Defaults),
	})
}

func (ea EmailAddress) toMap() map[string]interface{} {
	return map[string]interface{}{
		"name":  ea.Name,
		"email": ea.Email,
	}
}

func (eas EmailAddresses) toSliceMap() (out []map[string]interface{}) {
	for _, ea := range eas {
		out = append(out, ea.toMap())
	}
	return
}

func getEmailAddresses(d *schema.ResourceData, prefix string) (addresses []EmailAddress) {
	for i := 0; i < d.Get(prefix+".#").(int); i++ {
		subprefix := fmt.Sprintf(prefix+".%d", i)
		if ea := getEmailAddress(d, subprefix); ea != nil {
			addresses = append(addresses, *ea)
		}
	}
	return
}

func getEmailAddress(d *schema.ResourceData, prefix string) *EmailAddress {
	var n, e string
	if v, ok := d.GetOk(prefix + ".name"); ok {
		n = v.(string)
	}
	if v, ok := d.GetOk(prefix + ".email"); ok {
		e = v.(string)
	}
	if n != "" && e != "" {
		return &EmailAddress{n, e}
	}
	return nil
}
