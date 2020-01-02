package organization

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"time_zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func setAttributes(d *schema.ResourceData, org *client.Organization) {
	d.Set("name", org.Name)
	d.Set("time_zone", org.TimeZone)
}

func getAttributes(d *schema.ResourceData) client.Organization {
	return client.Organization{
		Id:       d.Id(),
		Name:     d.Get("name").(string),
		TimeZone: d.Get("time_zone").(string),
	}
}

func create(d *schema.ResourceData, m interface{}) error {
	org, err := config.Client(m).Organization.Create(getAttributes(d))
	if err != nil {
		return err
	}

	d.SetId(org.Id)
	setAttributes(d, org)
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	org, err := config.Client(m).Organization.Get(d.Id())
	if err != nil {
		return err
	}

	setAttributes(d, org)
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	_, err := config.Client(m).Organization.Update(getAttributes(d))
	if err != nil {
		return err
	}

	d.Partial(false)
	return nil
}

func delete(d *schema.ResourceData, m interface{}) error {
	err := config.Client(m).Organization.Delete(d.Id())
	return err
}

func importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	err := read(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
