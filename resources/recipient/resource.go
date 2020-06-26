package recipient

import (
	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/config"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: delete,

		Importer: &schema.ResourceImporter{
			State: importer,
		},

		Schema: map[string]*schema.Schema{
			"application_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func setAttributes(d *schema.ResourceData, recipient *client.Recipient) {
	d.Set("application_id", recipient.ApplicationId)
	d.Set("email", recipient.Email)
}

func getAttributes(d *schema.ResourceData) client.Recipient {
	return client.Recipient{
		Id:            d.Id(),
		ApplicationId: d.Get("application_id").(string),
		Email:         d.Get("email").(string),
	}
}

func create(d *schema.ResourceData, m interface{}) error {
	recipient, err := config.Client(m).Recipient.Create(getAttributes(d))
	if err != nil {
		return err
	}

	d.SetId(recipient.Id)
	setAttributes(d, recipient)
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	recipient, err := config.Client(m).Recipient.Get(d.Id())
	if err != nil {
		return err
	}

	setAttributes(d, recipient)
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	_, err := config.Client(m).Recipient.Update(getAttributes(d))
	if err != nil {
		return err
	}

	d.Partial(false)
	return nil
}

func delete(d *schema.ResourceData, m interface{}) error {
	err := config.Client(m).Recipient.Delete(d.Id())
	return err
}

func importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	err := read(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
