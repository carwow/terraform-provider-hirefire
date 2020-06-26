package provider

import (
	"github.com/carwow/terraform-provider-hirefire/config"
	"github.com/carwow/terraform-provider-hirefire/resources/account"
	"github.com/carwow/terraform-provider-hirefire/resources/application"
	"github.com/carwow/terraform-provider-hirefire/resources/manager"
	"github.com/carwow/terraform-provider-hirefire/resources/organization"
	"github.com/carwow/terraform-provider-hirefire/resources/recipient"
	"github.com/carwow/terraform-provider-hirefire/resources/time_range"
	"github.com/carwow/terraform-provider-hirefire/resources/user"
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
			"hirefire_manager":      manager.Resource(),
			"hirefire_time_range":   time_range.Resource(),
			"hirefire_recipient":    recipient.Resource(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"hirefire_user": user.DataSource(),
		},

		ConfigureFunc: config.Init,
	}
}
