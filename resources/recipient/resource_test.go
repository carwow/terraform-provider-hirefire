package recipient_test

import (
	"fmt"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const resourceName = "hirefire_recipient.foobar"

func TestAccRecipient(t *testing.T) {
	orgName := fmt.Sprintf("test-%s", helper.RandString(10))
	recipient := &client.Recipient{}

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: func(orgName string, recipient *client.Recipient) string {
					*recipient = client.Recipient{
						Email: "test-create@example.com",
					}
					return config(orgName, recipient)
				}(orgName, recipient),
				Check: checks(*recipient),
			},
			{
				Config: func(orgName string, recipient *client.Recipient) string {
					*recipient = client.Recipient{
						Email: "test-update@example.com",
					}
					return config(orgName, recipient)
				}(orgName, recipient),
				Check: checks(*recipient),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func config(orgName string, recipient *client.Recipient) string {
	return fmt.Sprintf(`
		resource "hirefire_organization" "foobar" {
			name = "%s"
			time_zone = "UTC"
		}

		resource "hirefire_account" "foobar" {
			organization_id = hirefire_organization.foobar.id
		}

		resource "hirefire_application" "foobar" {
			account_id = hirefire_account.foobar.id
			name = "foobar"
		}

		resource "hirefire_recipient" "foobar" {
			application_id = hirefire_application.foobar.id
			email = "%s"
		}`, orgName, recipient.Email)
}

func checks(recipient client.Recipient) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		checkAttributes(recipient),
		checkExists(recipient),
	)
}

func checkAttributes(recipient client.Recipient) resource.TestCheckFunc {
	return helper.CheckResourceAttributes(resourceName, map[string]string{
		"email": recipient.Email,
	})
}

func checkExists(recipient client.Recipient) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := helper.GetResourceID(state, resourceName)
		if err != nil {
			return err
		}

		actualRecipient, err := helper.Client().Recipient.Get(id)
		if err != nil {
			return err
		}

		recipient.Id = actualRecipient.Id
		recipient.ApplicationId = actualRecipient.ApplicationId
		return helper.Equals(recipient, *actualRecipient)
	}
}

func checkDestroy(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().Recipient.Get(id)
	if err == nil {
		return fmt.Errorf("Not destroyed: %s", resourceName)
	}
	return nil
}
