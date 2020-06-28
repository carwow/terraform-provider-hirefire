package membership_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const resourceName = "hirefire_membership.foobar"

func TestAccMembership(t *testing.T) {
	orgName := fmt.Sprintf("test-%s", helper.RandString(10))
	mem := &client.Membership{
		UserId: os.Getenv("HIREFIRE_TEST_MEMBERSHIP_USER_ID"),
	}

	if mem.UserId == "" {
		t.Skip("Membership resource test skipped unless env " +
			"'HIREFIRE_TEST_MEMBERSHIP_USER_ID' is set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: func(orgName string, mem *client.Membership) string {
					*mem = client.Membership{
						UserId: mem.UserId,
						Owner:  helper.RandBool(),
					}
					return config(orgName, mem)
				}(orgName, mem),
				Check: checks(*mem),
			},
			{
				Config: func(orgName string, mem *client.Membership) string {
					*mem = client.Membership{
						UserId: mem.UserId,
						Owner:  helper.RandBool(),
					}
					return config(orgName, mem)
				}(orgName, mem),
				Check: checks(*mem),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func config(orgName string, mem *client.Membership) string {
	return fmt.Sprintf(`
		resource "hirefire_organization" "foobar" {
			name = "%s"
			time_zone = "UTC"
		}

		resource "hirefire_membership" "foobar" {
			organization_id = hirefire_organization.foobar.id
			user_id         = "%s"
			owner           = %t
		}`, orgName, mem.UserId, mem.Owner)
}

func checks(mem client.Membership) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		checkAttributes(mem),
		checkExists(mem),
	)
}

func checkAttributes(mem client.Membership) resource.TestCheckFunc {
	return helper.CheckResourceAttributes(resourceName, map[string]string{
		"owner": strconv.FormatBool(mem.Owner),
	})
}

func checkExists(mem client.Membership) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := helper.GetResourceID(state, resourceName)
		if err != nil {
			return err
		}

		actualMembership, err := helper.Client().Membership.Get(id)
		if err != nil {
			return err
		}

		mem.Id = actualMembership.Id
		mem.OrganizationId = actualMembership.OrganizationId
		return helper.Equals(mem, *actualMembership)
	}
}

func checkDestroy(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().Membership.Get(id)
	if err == nil {
		return fmt.Errorf("Not destroyed: %s", resourceName)
	}
	return nil
}
