package client

import (
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMembership(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/memberships/ID-123")
		rw.Write([]byte(`{
			"membership": {
				"id":              "ID-123",
				"organization_id": "ID-999",
				"user_id":         "ID-888",
				"owner":           true
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	membership, err := client.Membership.Get("ID-123")
	assert.Ok(t, err)

	expected := &Membership{
		Id:             "ID-123",
		OrganizationId: "ID-999",
		UserId:         "ID-888",
		Owner:          true,
	}
	assert.Equals(t, expected, membership)
}
