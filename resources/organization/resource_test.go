package organization_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/provider"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccImportBasic(t *testing.T) {
	name := fmt.Sprintf("test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			if os.Getenv("HIREFIRE_API_KEY") == "" {
				t.Fatal("HIREFIRE_API_KEY must be set")
			}
		},
		Providers: map[string]terraform.ResourceProvider{
			"hirefire": provider.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: config(name, "UTC"),
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
