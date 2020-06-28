package membership

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
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func setAttributes(d *schema.ResourceData, mem *client.Membership) {
	d.Set("organization_id", mem.OrganizationId)
	d.Set("user_id", mem.UserId)
	d.Set("owner", mem.Owner)
}

func getAttributes(d *schema.ResourceData) client.Membership {
	return client.Membership{
		Id:             d.Id(),
		OrganizationId: d.Get("organization_id").(string),
		UserId:         d.Get("user_id").(string),
		Owner:          d.Get("owner").(bool),
	}
}

func create(d *schema.ResourceData, m interface{}) error {
	mem, err := config.Client(m).Membership.Create(getAttributes(d))
	if err != nil {
		return err
	}

	d.SetId(mem.Id)
	setAttributes(d, mem)
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	mem, err := config.Client(m).Membership.Get(d.Id())
	if err != nil {
		return err
	}

	setAttributes(d, mem)
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	_, err := config.Client(m).Membership.Update(getAttributes(d))
	if err != nil {
		return err
	}

	d.Partial(false)
	return nil
}

func delete(d *schema.ResourceData, m interface{}) error {
	err := config.Client(m).Membership.Delete(d.Id())
	return err
}

func importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	err := read(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
