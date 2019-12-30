package organization_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/provider"
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccImportBasic(t *testing.T) {
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
				Config:             config("carwow", "UTC"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:  "hirefire_organization.foobar",
				ImportState:   true,
				ImportStateId: "4d217e79-737c-4a37-96f8-61611e03a213",
				ImportStateCheck: func(instances []*terraform.InstanceState) error {
					attr := instances[0].Attributes
					assert.Equals(t, "carwow", attr["name"])
					assert.Equals(t, "UTC", attr["time_zone"])
					return nil
				},
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
