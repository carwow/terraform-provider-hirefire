package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carwow/terraform-provider-hirefire/testing/assert"
)

func TestGetAccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/accounts/ID-123")
		rw.Write([]byte(`{
			"account": {
				"id":              "ID-123",
				"organization_id": "ID-999"
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	account, err := client.Account.Get("ID-123")
	assert.Ok(t, err)

	expected := &Account{
		Id:             "ID-123",
		OrganizationId: "ID-999",
	}
	assert.Equals(t, expected, account)
}
