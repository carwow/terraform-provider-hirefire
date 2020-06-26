package client

import (
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/users/ID-123")
		rw.Write([]byte(`{
			"user": {
				"id":            "ID-123",
				"email":         "test@example.com",
				"notifications": true
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	user, err := client.User.Get("ID-123")
	assert.Ok(t, err)

	expected := &User{
		Id:            "ID-123",
		Email:         "test@example.com",
		Notifications: true,
	}
	assert.Equals(t, expected, user)

}
