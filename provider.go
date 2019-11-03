package main

import (
	"github.com/carwow/terraform-provider-hirefire/resources/organization"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"hirefire_organization": organization.Resource(),
		},
	}
}
