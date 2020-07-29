package manager_test

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

const resourceName = "hirefire_manager.foobar"

func TestAccManager(t *testing.T) {
	orgName := fmt.Sprintf("test-%s", helper.RandString(10))
	manager := &client.Manager{}

	resource.Test(t, resource.TestCase{
		PreCheck:     helper.PreCheck(t),
		Providers:    helper.Providers(),
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				Config: func(orgName string, manager *client.Manager) string {
					*manager = client.Manager{
						Name:    fmt.Sprintf("test-%s", helper.RandString(10)),
						Type:    "Manager::Web::Logplex::RPM",
						Enabled: false,
						Minimum: 1,
						Maximum: 5,

						Ratio:                ptr.Int(100),
						UpscaleSensitivity:   ptr.Int(1),
						DownscaleSensitivity: ptr.Int(2),
						UpscaleTimeout:       ptr.Int(2),
						DownscaleTimeout:     ptr.Int(5),
					}
					return config(orgName, manager)
				}(orgName, manager),
				Check: checks(*manager),
			},
			{
				Config: func(orgName string, manager *client.Manager) string {
					*manager = client.Manager{
						Name:    fmt.Sprintf("test-%s", helper.RandString(10)),
						Type:    "Manager::Web::Logplex::QueueTime",
						Enabled: true,
						Minimum: 2,
						Maximum: 10,

						Aggregation:          ptr.String("percentile"),
						Percentile:           ptr.Int(99),
						MinimumQueueTime:     ptr.Int(200),
						MaximumQueueTime:     ptr.Int(400),
						UpscaleQuantity:      ptr.Int(5),
						DownscaleQuantity:    ptr.Int(1),
						UpscaleSensitivity:   ptr.Int(2),
						DownscaleSensitivity: ptr.Int(1),
						UpscaleTimeout:       ptr.Int(1),
						DownscaleTimeout:     ptr.Int(2),
					}
					return config(orgName, manager)
				}(orgName, manager),
				Check: checks(*manager),
			},
			{
				Config: func(orgName string, manager *client.Manager) string {
					*manager = client.Manager{
						Name:    fmt.Sprintf("test-%s", helper.RandString(10)),
						Type:    "Manager::Web::Logplex::Load",
						Enabled: true,
						Minimum: 2,
						Maximum: 10,

						LastMinutes:          ptr.Int(5),
						MinimumLoad:          ptr.Int(40),
						MaximumLoad:          ptr.Int(85),
						UpscaleQuantity:      ptr.Int(5),
						DownscaleQuantity:    ptr.Int(1),
						UpscaleSensitivity:   ptr.Int(2),
						DownscaleSensitivity: ptr.Int(1),
						UpscaleTimeout:       ptr.Int(1),
						DownscaleTimeout:     ptr.Int(2),
					}
					return config(orgName, manager)
				}(orgName, manager),
				Check: checks(*manager),
			},
			{
				Config: func(orgName string, manager *client.Manager) string {
					*manager = client.Manager{
						Name:    fmt.Sprintf("test-%s", helper.RandString(10)),
						Type:    "Manager::Worker::HireFire::JobQueue",
						Enabled: true,
						Minimum: 2,
						Maximum: 10,

						Decrementable:        ptr.Bool(true),
						Ratio:                ptr.Int(10),
						UpscaleSensitivity:   ptr.Int(2),
						DownscaleSensitivity: ptr.Int(1),
						UpscaleTimeout:       ptr.Int(1),
						DownscaleTimeout:     ptr.Int(2),
						UpscaleLimit:         ptr.Int(3),
						DownscaleLimit:       ptr.Int(4),
					}
					return config(orgName, manager)
				}(orgName, manager),
				Check: checks(*manager),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func configBase(orgName string, manager *client.Manager, managerAttributes string) string {
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
			name = "foobar-manager"
		}

		resource "hirefire_manager" "foobar" {
			application_id = hirefire_application.foobar.id
			name = "%s"
			type = "%s"
			enabled = %t
			minimum = %d
			maximum = %d

			%s
		}`,
		orgName,
		manager.Name,
		manager.Type,
		manager.Enabled,
		manager.Minimum,
		manager.Maximum,
		managerAttributes,
	)
}

func config(orgName string, manager *client.Manager) string {
	switch manager.Type {
	case "Manager::Web::Logplex::RPM":
		return configBase(orgName, manager, fmt.Sprintf(`
			ratio                 = %d
			upscale_sensitivity   = %d
			downscale_sensitivity = %d
			upscale_timeout       = %d
			downscale_timeout     = %d
			`,
			*manager.Ratio,
			*manager.UpscaleSensitivity,
			*manager.DownscaleSensitivity,
			*manager.UpscaleTimeout,
			*manager.DownscaleTimeout,
		))
	case "Manager::Web::Logplex::QueueTime":
		return configBase(orgName, manager, fmt.Sprintf(`
			aggregation           = "%s"
			percentile            = %d
			minimum_queue_time    = %d
			maximum_queue_time    = %d
			upscale_quantity      = %d
			downscale_quantity    = %d
			upscale_sensitivity   = %d
			downscale_sensitivity = %d
			upscale_timeout       = %d
			downscale_timeout     = %d
			`,
			*manager.Aggregation,
			*manager.Percentile,
			*manager.MinimumQueueTime,
			*manager.MaximumQueueTime,
			*manager.UpscaleQuantity,
			*manager.DownscaleQuantity,
			*manager.UpscaleSensitivity,
			*manager.DownscaleSensitivity,
			*manager.UpscaleTimeout,
			*manager.DownscaleTimeout,
		))
	case "Manager::Web::Logplex::Load":
		return configBase(orgName, manager, fmt.Sprintf(`
			last_minutes          = %d
			minimum_load          = %d
			maximum_load          = %d
			upscale_quantity      = %d
			downscale_quantity    = %d
			upscale_sensitivity   = %d
			downscale_sensitivity = %d
			upscale_timeout       = %d
			downscale_timeout     = %d
			`,
			*manager.LastMinutes,
			*manager.MinimumLoad,
			*manager.MaximumLoad,
			*manager.UpscaleQuantity,
			*manager.DownscaleQuantity,
			*manager.UpscaleSensitivity,
			*manager.DownscaleSensitivity,
			*manager.UpscaleTimeout,
			*manager.DownscaleTimeout,
		))
	case "Manager::Worker::HireFire::JobQueue":
		return configBase(orgName, manager, fmt.Sprintf(`
			decrementable         = %t
			ratio                 = %d
			upscale_sensitivity   = %d
			downscale_sensitivity = %d
			upscale_timeout       = %d
			downscale_timeout     = %d
			upscale_limit         = %d
			downscale_limit       = %d
			`,
			*manager.Decrementable,
			*manager.Ratio,
			*manager.UpscaleSensitivity,
			*manager.DownscaleSensitivity,
			*manager.UpscaleTimeout,
			*manager.DownscaleTimeout,
			*manager.UpscaleLimit,
			*manager.DownscaleLimit,
		))
	default:
		return configBase(orgName, manager, "")
	}
}

func checks(manager client.Manager) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		checkAttributes(manager),
		checkExists(manager),
	)
}

func checkAttributes(manager client.Manager) resource.TestCheckFunc {
	return helper.CheckResourceAttributes(resourceName, map[string]string{
		"name":    manager.Name,
		"type":    manager.Type,
		"enabled": strconv.FormatBool(manager.Enabled),
		"minimum": strconv.Itoa(manager.Minimum),
		"maximum": strconv.Itoa(manager.Maximum),

		"aggregation":           helper.StringOrEmpty(manager.Aggregation),
		"percentile":            helper.ItoaOrZero(manager.Percentile),
		"minimum_response_time": helper.ItoaOrZero(manager.MinimumResponseTime),
		"maximum_response_time": helper.ItoaOrZero(manager.MaximumResponseTime),
		"minimum_load":          helper.ItoaOrZero(manager.MinimumLoad),
		"maximum_load":          helper.ItoaOrZero(manager.MaximumLoad),
		"minimum_apdex":         helper.ItoaOrZero(manager.MinimumApdex),
		"maximum_apdex":         helper.ItoaOrZero(manager.MaximumApdex),
		"last_minutes":          helper.ItoaOrZero(manager.LastMinutes),
		"ratio":                 helper.ItoaOrZero(manager.Ratio),
		"decrementable":         helper.BoolOrFalse(manager.Decrementable),
		"url":                   helper.StringOrEmpty(manager.Url),
		"upscale_quantity":      helper.ItoaOrZero(manager.UpscaleQuantity),
		"downscale_quantity":    helper.ItoaOrZero(manager.DownscaleQuantity),
		"upscale_sensitivity":   helper.ItoaOrZero(manager.UpscaleSensitivity),
		"downscale_sensitivity": helper.ItoaOrZero(manager.DownscaleSensitivity),
		"upscale_timeout":       helper.ItoaOrZero(manager.UpscaleTimeout),
		"downscale_timeout":     helper.ItoaOrZero(manager.DownscaleTimeout),
		"upscale_limit":         helper.ItoaOrZero(manager.UpscaleLimit),
		"downscale_limit":       helper.ItoaOrZero(manager.DownscaleLimit),
		"scale_up_on_503":       helper.BoolOrFalse(manager.ScaleUpOn503),
		"new_relic_api_key":     helper.StringOrEmpty(manager.NewRelicApiKey),
		"new_relic_account_id":  helper.StringOrEmpty(manager.NewRelicAccountId),
		"new_relic_app_id":      helper.StringOrEmpty(manager.NewRelicAppId),
		"notify":                helper.BoolOrFalse(manager.Notify),
		"notify_quantity":       helper.ItoaOrZero(manager.NotifyQuantity),
		"notify_after":          helper.ItoaOrZero(manager.NotifyAfter),
	})
}

func checkExists(manager client.Manager) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := helper.GetResourceID(state, resourceName)
		if err != nil {
			return err
		}

		actualManager, err := helper.Client().Manager.Get(id)
		if err != nil {
			return err
		}

		manager.Id = actualManager.Id
		manager.ApplicationId = actualManager.ApplicationId
		return helper.Equals(manager, *actualManager)
	}
}

func checkDestroy(state *terraform.State) error {
	id, err := helper.GetResourceID(state, resourceName)
	if err != nil {
		return err
	}

	_, err = helper.Client().Manager.Get(id)
	if err == nil {
		return fmt.Errorf("Not destroyed: %s", resourceName)
	}
	return nil
}
