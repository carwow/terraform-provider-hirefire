package client

import (
	"github.com/carwow/terraform-provider-hirefire/ptr"
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApplication(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/applications/ID-123")
		rw.Write([]byte(`{
			"application": {
				"id":         "ID-123",
				"account_id": "ID-999",
				"name":       "app-name"
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	application, err := client.Application.Get("ID-123")
	assert.Ok(t, err)

	expected := &Application{
		Id:                         "ID-123",
		AccountId:                  "ID-999",
		Name:                       "app-name",
		CustomDomain:               nil,
		LogplexDrainToken:          nil,
		Ssl:                        false,
		RestartCrashedDynos:        false,
		NewIssueNotifications:      false,
		ResolvedIssueNotifications: false,
	}
	assert.Equals(t, expected, application)

}

func TestGetApplicationEverything(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/applications/ID-123")
		rw.Write([]byte(`{
			"application": {
				"id":                           "ID-123",
				"account_id":                   "ID-999",
				"name":                         "app-name",
				"custom_domain":                "custom-domain",
				"logplex_drain_token":          "drain-token",
				"ssl":                          true,
				"restart_crashed_dynos":        true,
				"new_issue_notifications":      true,
				"resolved_issue_notifications": true
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	application, err := client.Application.Get("ID-123")
	assert.Ok(t, err)

	expected := &Application{
		Id:                         "ID-123",
		AccountId:                  "ID-999",
		Name:                       "app-name",
		CustomDomain:               ptr.String("custom-domain"),
		LogplexDrainToken:          ptr.String("drain-token"),
		Ssl:                        true,
		RestartCrashedDynos:        true,
		NewIssueNotifications:      true,
		ResolvedIssueNotifications: true,
	}
	assert.Equals(t, expected, application)

}
