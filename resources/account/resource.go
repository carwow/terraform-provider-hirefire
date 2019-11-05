package account

import (
	"github.com/carwow/terraform-provider-hirefire/config"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Delete: delete,

		Importer: &schema.ResourceImporter{
			State: importer,
		},

		Schema: map[string]*schema.Schema{
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func delete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func importer(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	account, err := config.Client(meta).Account.Get(d.Id())
	if err != nil {
		return nil, err
	}

	d.SetId(account.Id)
	d.Set("organization_id", account.OrganizationId)

	return []*schema.ResourceData{d}, nil
}
