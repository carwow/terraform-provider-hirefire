package account_test

import (
	"fmt"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const resourceName = "hirefire_account.foobar"

func TestAccAccount(t *testing.T) {
	orgName := fmt.Sprintf("test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: config(orgName),
				Check:  checkExists,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func config(orgName string) string {
	return fmt.Sprintf(`
resource "hirefire_organization" "foobar" {
	name = "%s"
	time_zone = "UTC"
}

resource "hirefire_account" "foobar" {
	organization_id = hirefire_organization.foobar.id
}`, orgName)
}

func checkExists(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().Account.Get(id)
	if err != nil {
		return err
	}
	return nil
}

func checkDestroy(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().Account.Get(id)
	if err == nil {
		return fmt.Errorf("Not destroyed: %s", resourceName)
	}
	return nil
}
