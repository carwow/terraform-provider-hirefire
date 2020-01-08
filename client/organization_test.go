package client

import (
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrganization(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/organizations/ID-123")
		rw.Write([]byte(`{
			"organization": {
				"id":        "ID-123",
				"name":      "carwow",
				"time_zone": "Europe/London"
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	organization, err := client.Organization.Get("ID-123")
	assert.Ok(t, err)

	expected := &Organization{
		Id:       "ID-123",
		Name:     "carwow",
		TimeZone: "Europe/London",
	}
	assert.Equals(t, expected, organization)

}
