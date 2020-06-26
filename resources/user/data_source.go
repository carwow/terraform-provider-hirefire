package user

import (
	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/config"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: read,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"notifications": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func setAttributes(d *schema.ResourceData, user *client.User) {
	d.SetId(user.Id)
	d.Set("email", user.Email)
	d.Set("notifications", user.Notifications)
}

func read(d *schema.ResourceData, m interface{}) error {
	user, err := config.Client(m).User.Get(d.Get("id").(string))
	if err != nil {
		return err
	}

	setAttributes(d, user)
	return nil
}
