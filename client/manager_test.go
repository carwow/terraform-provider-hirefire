package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/ptr"
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
)

func TestGetManager(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/managers/ID-123")
		rw.Write([]byte(`{
			"manager": {
				"id":             "ID-123",
				"application_id": "ID-999",
				"name":           "web",
				"type":           "Manager::Web::Logplex::RPM",
				"enabled":        true,
				"minimum":        2,
				"maximum":        5
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	manager, err := client.Manager.Get("ID-123")
	assert.Ok(t, err)

	expected := &Manager{
		Id:            "ID-123",
		ApplicationId: "ID-999",
		Name:          "web",
		Type:          "Manager::Web::Logplex::RPM",
		Enabled:       true,
		Minimum:       2,
		Maximum:       5,
	}
	assert.Equals(t, expected, manager)

}

func TestGetManagerEverything(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/managers/ID-123")
		rw.Write([]byte(`{
			"manager": {
				"id":             "ID-123",
				"application_id": "ID-999",
				"name":           "web",
				"type":           "Manager::Web::Logplex::RPM",
				"enabled":        true,
				"minimum":        2,
				"maximum":        5,

				"aggregation":            "percentile",
				"percentile":             99,
				"minimum_latency":        100,
				"maximum_latency":        150,
				"minimum_queue_time":     200,
				"maximum_queue_time":     400,
				"minimum_response_time":  500,
				"maximum_response_time":  1000,
				"minimum_connect_time":   300,
				"maximum_connect_time":   600,
				"minimum_load":           1,
				"maximum_load":           2,
				"minimum_apdex":          95,
				"maximum_apdex":          99,
				"ratio":                  10,
				"decrementable":          true,
				"url":                    "https://www.example.com",
				"upscale_quantity":       5,
				"downscale_quantity":     1,
				"upscale_sensitivity":    1,
				"downscale_sensitivity":  2,
				"upscale_timeout":        1,
				"downscale_timeout":      2,
				"upscale_limit":          0,
				"downscale_limit":        2,
				"scale_up_on_503":        true,
				"new_relic_api_key":      "newrelic-api-key",
				"new_relic_account_id":   "newrelic-account-id",
				"new_relic_app_id":       "newrelic-app-id",
				"notify":                 true,
				"notify_quantity":        5,
				"notify_after":           10,
				"upscale_on_initial_job": true
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	manager, err := client.Manager.Get("ID-123")
	assert.Ok(t, err)

	expected := &Manager{
		Id:            "ID-123",
		ApplicationId: "ID-999",
		Name:          "web",
		Type:          "Manager::Web::Logplex::RPM",
		Enabled:       true,
		Minimum:       2,
		Maximum:       5,

		Aggregation:          ptr.String("percentile"),
		Percentile:           ptr.Int(99),
		MinimumLatency:       ptr.Int(100),
		MaximumLatency:       ptr.Int(150),
		MinimumQueueTime:     ptr.Int(200),
		MaximumQueueTime:     ptr.Int(400),
		MinimumResponseTime:  ptr.Int(500),
		MaximumResponseTime:  ptr.Int(1000),
		MinimumConnectTime:   ptr.Int(300),
		MaximumConnectTime:   ptr.Int(600),
		MinimumLoad:          ptr.Int(1),
		MaximumLoad:          ptr.Int(2),
		MinimumApdex:         ptr.Int(95),
		MaximumApdex:         ptr.Int(99),
		Ratio:                ptr.Int(10),
		Decrementable:        ptr.Bool(true),
		Url:                  ptr.String("https://www.example.com"),
		UpscaleQuantity:      ptr.Int(5),
		DownscaleQuantity:    ptr.Int(1),
		UpscaleSensitivity:   ptr.Int(1),
		DownscaleSensitivity: ptr.Int(2),
		UpscaleTimeout:       1,
		DownscaleTimeout:     2,
		UpscaleLimit:         0,
		DownscaleLimit:       2,
		ScaleUpOn503:         ptr.Bool(true),
		NewRelicApiKey:       ptr.String("newrelic-api-key"),
		NewRelicAccountId:    ptr.String("newrelic-account-id"),
		NewRelicAppId:        ptr.String("newrelic-app-id"),
		Notify:               true,
		NotifyQuantity:       5,
		NotifyAfter:          10,
		UpscaleOnInitialJob:  ptr.Bool(true),
	}
	assert.Equals(t, expected, manager)
}
