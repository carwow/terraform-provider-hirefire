package organization_test

import (
	"fmt"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const resourceName = "hirefire_organization.foobar"

func TestAccOrganization(t *testing.T) {
	org := client.Organization{}

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: func(org *client.Organization) string {
					org.Name = fmt.Sprintf("test-%s", acctest.RandString(10))
					org.TimeZone = "London"
					return config(org)
				}(&org),
				Check: checks(org),
			},
			{
				Config: func(org *client.Organization) string {
					org.Name = fmt.Sprintf("test-%s", acctest.RandString(10))
					org.TimeZone = "Lisbon"
					return config(org)
				}(&org),
				Check: checks(org),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func config(org *client.Organization) string {
	return fmt.Sprintf(`
resource "hirefire_organization" "foobar" {
	name = "%s"
	time_zone = "%s"
}`, org.Name, org.TimeZone)
}

func checks(org client.Organization) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		checkAttributes(org),
		checkExists(org),
	)
}

func checkAttributes(org client.Organization) resource.TestCheckFunc {
	return helper.CheckResourceAttributes(resourceName, map[string]string{
		"name":      org.Name,
		"time_zone": org.TimeZone,
	})
}

func checkExists(org client.Organization) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := helper.GetResourceID(state, resourceName)
		if err != nil {
			return err
		}

		actualOrg, err := helper.Client().Organization.Get(id)
		if err != nil {
			return err
		}

		org.Id = actualOrg.Id
		return helper.Equals(org, *actualOrg)
	}
}

func checkDestroy(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().Organization.Get(id)
	if err == nil {
		return fmt.Errorf("Not destroyed: %s", resourceName)
	}
	return nil
}
