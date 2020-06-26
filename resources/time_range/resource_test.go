package time_range_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/testing/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const resourceName = "hirefire_time_range.foobar"

func TestAccTimeRange(t *testing.T) {
	orgName := fmt.Sprintf("test-%s", helper.RandString(10))
	timeRange := &client.TimeRange{}

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: func(orgName string, timeRange *client.TimeRange) string {
					*timeRange = client.TimeRange{
						FromMinute:  helper.RandInt(15, 30),
						UntilMinute: helper.RandInt(90, 120),
						Minimum:     helper.RandInt(1, 3),
						Maximum:     helper.RandInt(4, 6),
						Position:    helper.RandInt(0, 2),
						Monday:      helper.RandBool(),
						Tuesday:     helper.RandBool(),
						Wednesday:   helper.RandBool(),
						Thursday:    helper.RandBool(),
						Friday:      helper.RandBool(),
						Saturday:    helper.RandBool(),
						Sunday:      helper.RandBool(),
					}
					return config(orgName, timeRange)
				}(orgName, timeRange),
				Check: checks(*timeRange),
			},
			{
				Config: func(orgName string, timeRange *client.TimeRange) string {
					*timeRange = client.TimeRange{
						FromMinute:  helper.RandInt(15, 30),
						UntilMinute: helper.RandInt(90, 120),
						Minimum:     helper.RandInt(1, 3),
						Maximum:     helper.RandInt(4, 6),
						Position:    helper.RandInt(0, 2),
						Monday:      helper.RandBool(),
						Tuesday:     helper.RandBool(),
						Wednesday:   helper.RandBool(),
						Thursday:    helper.RandBool(),
						Friday:      helper.RandBool(),
						Saturday:    helper.RandBool(),
						Sunday:      helper.RandBool(),
					}
					return config(orgName, timeRange)
				}(orgName, timeRange),
				Check: checks(*timeRange),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func config(orgName string, timeRange *client.TimeRange) string {
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
			name = "foobar-time-range"
		}

		resource "hirefire_manager" "foobar" {
			application_id = hirefire_application.foobar.id
			name = "foobar-time-range"
			type = "Manager::Web::Logplex::RPM"
			enabled = false
			minimum = 1
			maximum = 2

			ratio                 = 100
			upscale_sensitivity   = 1
			downscale_sensitivity = 2
			upscale_timeout       = 1
			downscale_timeout     = 2
		}

		resource "hirefire_time_range" "foobar" {
			manager_id = hirefire_manager.foobar.id

			from_minute      = %d
			until_minute     = %d
			minimum          = %d
			maximum          = %d
			position         = %d
			monday           = %t
			tuesday          = %t
			wednesday        = %t
			thursday         = %t
			friday           = %t
			saturday         = %t
			sunday           = %t
		}`,
		orgName,
		timeRange.FromMinute,
		timeRange.UntilMinute,
		timeRange.Minimum,
		timeRange.Maximum,
		timeRange.Position,
		timeRange.Monday,
		timeRange.Tuesday,
		timeRange.Wednesday,
		timeRange.Thursday,
		timeRange.Friday,
		timeRange.Saturday,
		timeRange.Sunday,
	)
}

func checks(timeRange client.TimeRange) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		checkAttributes(timeRange),
		checkExists(timeRange),
	)
}

func checkAttributes(timeRange client.TimeRange) resource.TestCheckFunc {
	return helper.CheckResourceAttributes(resourceName, map[string]string{
		"from_minute":  strconv.Itoa(timeRange.FromMinute),
		"until_minute": strconv.Itoa(timeRange.UntilMinute),
		"minimum":      strconv.Itoa(timeRange.Minimum),
		"maximum":      strconv.Itoa(timeRange.Maximum),
		"position":     strconv.Itoa(timeRange.Position),
		"monday":       strconv.FormatBool(timeRange.Monday),
		"tuesday":      strconv.FormatBool(timeRange.Tuesday),
		"wednesday":    strconv.FormatBool(timeRange.Wednesday),
		"thursday":     strconv.FormatBool(timeRange.Thursday),
		"friday":       strconv.FormatBool(timeRange.Friday),
		"saturday":     strconv.FormatBool(timeRange.Saturday),
		"sunday":       strconv.FormatBool(timeRange.Sunday),
	})
}

func checkExists(timeRange client.TimeRange) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := helper.GetResourceID(state, resourceName)
		if err != nil {
			return err
		}

		actualTimeRange, err := helper.Client().TimeRange.Get(id)
		if err != nil {
			return err
		}

		timeRange.Id = actualTimeRange.Id
		timeRange.ManagerId = actualTimeRange.ManagerId
		return helper.Equals(timeRange, *actualTimeRange)
	}
}

func checkDestroy(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().TimeRange.Get(id)
	if err == nil {
		return fmt.Errorf("Not destroyed: %s", resourceName)
	}
	return nil
}
