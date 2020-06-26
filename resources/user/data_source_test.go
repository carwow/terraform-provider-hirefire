package user_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const dataSourceName = "data.hirefire_user.foobar"

func TestAccUser(t *testing.T) {
	user := &client.User{
		Id:    os.Getenv("HIREFIRE_TEST_USER_ID"),
		Email: os.Getenv("HIREFIRE_TEST_USER_EMAIL"),
	}

	if user.Id == "" || user.Email == "" {
		t.Skip("User data source test skipped unless env " +
			"'HIREFIRE_TEST_USER_ID' and 'HIREFIRE_TEST_USER_EMAIL' are set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  helper.PreCheck(t),
		Providers: helper.Providers(),
		Steps: []resource.TestStep{
			{
				Config: configDataSource(user),
				Check:  checkDataSource(*user),
			},
		},
	})
}

func configDataSource(user *client.User) string {
	return fmt.Sprintf(`
		data "hirefire_user" "foobar" {
			id = "%s"
		}`, user.Id)
}

func checkDataSource(user client.User) resource.TestCheckFunc {
	return helper.CheckResourceAttributes(dataSourceName, map[string]string{
		"email": user.Email,
	})
}
