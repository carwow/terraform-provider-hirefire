package provider

import (
	"github.com/carwow/terraform-provider-hirefire/config"
	"github.com/carwow/terraform-provider-hirefire/resources/account"
	"github.com/carwow/terraform-provider-hirefire/resources/application"
	"github.com/carwow/terraform-provider-hirefire/resources/organization"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HIREFIRE_API_KEY", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"hirefire_organization": organization.Resource(),
			"hirefire_account":      account.Resource(),
			"hirefire_application":  application.Resource(),
		},

		ConfigureFunc: config.Init,
	}
}
