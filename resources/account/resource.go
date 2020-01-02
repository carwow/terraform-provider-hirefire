package account

import (
	"github.com/carwow/terraform-provider-hirefire/client"
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

func setAttributes(d *schema.ResourceData, acc *client.Account) {
	d.Set("organization_id", acc.OrganizationId)
}

func getAttributes(d *schema.ResourceData) client.Account {
	return client.Account{
		OrganizationId: d.Get("organization_id").(string),
	}
}

func create(d *schema.ResourceData, m interface{}) error {
	acc, err := config.Client(m).Account.Create(getAttributes(d))
	if err != nil {
		return err
	}

	d.SetId(acc.Id)
	setAttributes(d, acc)
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	acc, err := config.Client(m).Account.Get(d.Id())
	if err != nil {
		return err
	}

	setAttributes(d, acc)
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	_, err := config.Client(m).Account.Update(getAttributes(d))
	if err != nil {
		return err
	}

	d.Partial(false)
	return nil
}

func delete(d *schema.ResourceData, m interface{}) error {
	err := config.Client(m).Account.Delete(d.Id())
	return err
}

func importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	err := read(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
