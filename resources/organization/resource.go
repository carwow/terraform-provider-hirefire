package organization

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: delete,

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
