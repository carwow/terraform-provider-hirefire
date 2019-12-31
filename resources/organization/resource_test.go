package organization_test

import (
	"fmt"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccImportBasic(t *testing.T) {
	name := fmt.Sprintf("test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: config(name, "UTC"),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				ResourceName:      "hirefire_organization.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func config(name, timeZone string) string {
	return fmt.Sprintf(`
resource "hirefire_organization" "foobar" {
    name = "%s"
    time_zone = "%s"
}`, name, timeZone)
}

func checkDestroy(s *terraform.State) error {
	client := helper.Client()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "hirefire_organization" {
			continue
		}

		_, err := client.Organization.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Still exists")
		}
	}

	return nil
}
