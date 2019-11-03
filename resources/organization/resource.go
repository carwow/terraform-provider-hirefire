package organization

import (
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

func create(d *schema.ResourceData, m interface{}) error {
	return read(d, m)
}

func read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	return read(d, m)
}

func delete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func importer(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	org, err := config.Client(meta).Organization.Get(d.Id())
	if err != nil {
		return nil, err
	}

	d.SetId(org.Id)
	d.Set("name", org.Name)
	d.Set("time_zone", org.TimeZone)

	return []*schema.ResourceData{d}, nil
}
