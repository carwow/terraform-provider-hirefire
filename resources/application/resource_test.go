package application_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/ptr"
	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const resourceName = "hirefire_application.foobar"

func TestAccApplication(t *testing.T) {
	orgName := fmt.Sprintf("test-%s", helper.RandString(10))
	app := &client.Application{}

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: func(orgName string, app *client.Application) string {
					*app = client.Application{
						Name: fmt.Sprintf("test-%s", helper.RandString(10)),
					}
					return config(orgName, app)
				}(orgName, app),
				Check: checks(*app),
			},
			{
				Config: func(orgName string, app *client.Application) string {
					*app = client.Application{
						Name: fmt.Sprintf("test-%s", helper.RandString(10)),
					}
					return config(orgName, app)
				}(orgName, app),
				Check: checks(*app),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccApplicationEverything(t *testing.T) {
	orgName := fmt.Sprintf("test-%s", helper.RandString(10))
	app := &client.Application{}
	logplexDrainToken := fmt.Sprintf("d.%s-%s-%s-%s-%s",
		helper.RandHex(8), helper.RandHex(4), helper.RandHex(4), helper.RandHex(4), helper.RandHex(12))

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: func(orgName string, app *client.Application) string {
					*app = client.Application{
						Name:                       fmt.Sprintf("test-%s", helper.RandString(10)),
						CustomDomain:               ptr.String(fmt.Sprintf("test-%s", helper.RandString(10))),
						LogplexDrainToken:          ptr.String(logplexDrainToken),
						Ssl:                        helper.RandBool(),
						RestartCrashedDynos:        helper.RandBool(),
						NewIssueNotifications:      helper.RandBool(),
						ResolvedIssueNotifications: helper.RandBool(),
					}
					return configEverything(orgName, app)
				}(orgName, app),
				Check: checks(*app),
			},
			{
				Config: func(orgName string, app *client.Application) string {
					*app = client.Application{
						Name:                       fmt.Sprintf("test-%s", helper.RandString(10)),
						CustomDomain:               ptr.String(fmt.Sprintf("test-%s", helper.RandString(10))),
						LogplexDrainToken:          ptr.String(logplexDrainToken),
						Ssl:                        helper.RandBool(),
						RestartCrashedDynos:        helper.RandBool(),
						NewIssueNotifications:      helper.RandBool(),
						ResolvedIssueNotifications: helper.RandBool(),
					}
					return configEverything(orgName, app)
				}(orgName, app),
				Check: checks(*app),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func configBase(orgName, appAttributes string) string {
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
			%s
		}`, orgName, appAttributes)
}

func config(orgName string, app *client.Application) string {
	return configBase(orgName, fmt.Sprintf(`name = "%s"`, app.Name))
}

func configEverything(orgName string, app *client.Application) string {
	return configBase(orgName, fmt.Sprintf(`
		name = "%s"
		custom_domain = "%s"
		logplex_drain_token = "%s"
		ssl = %t
		restart_crashed_dynos = %t
		new_issue_notifications = %t
		resolved_issue_notifications = %t
		checkup_frequency = %d
		`,
		app.Name,
		*app.CustomDomain,
		*app.LogplexDrainToken,
		app.Ssl,
		app.RestartCrashedDynos,
		app.NewIssueNotifications,
		app.ResolvedIssueNotifications,
		app.CheckupFrequency,
	))
}

func checks(app client.Application) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		checkAttributes(app),
		checkExists(app),
	)
}

func checkAttributes(app client.Application) resource.TestCheckFunc {
	return helper.CheckResourceAttributes(resourceName, map[string]string{
		"name":                         app.Name,
		"custom_domain":                helper.StringOrEmpty(app.CustomDomain),
		"logplex_drain_token":          helper.StringOrEmpty(app.LogplexDrainToken),
		"ssl":                          strconv.FormatBool(app.Ssl),
		"restart_crashed_dynos":        strconv.FormatBool(app.RestartCrashedDynos),
		"new_issue_notifications":      strconv.FormatBool(app.NewIssueNotifications),
		"resolved_issue_notifications": strconv.FormatBool(app.ResolvedIssueNotifications),
		"checkup_frequency":            strconv.Itoa(app.CheckupFrequency),
	})
}

func checkExists(app client.Application) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := helper.GetResourceID(state, resourceName)
		if err != nil {
			return err
		}

		actualApp, err := helper.Client().Application.Get(id)
		if err != nil {
			return err
		}

		app.Id = actualApp.Id
		app.AccountId = actualApp.AccountId
		app.Token = actualApp.Token
		return helper.Equals(app, *actualApp)
	}
}

func checkDestroy(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().Application.Get(id)
	if err == nil {
		return fmt.Errorf("Not destroyed: %s", resourceName)
	}
	return nil
}
