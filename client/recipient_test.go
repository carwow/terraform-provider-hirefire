package client

import (
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecipient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/recipients/ID-123")
		rw.Write([]byte(`{
			"recipient": {
				"id":             "ID-123",
				"application_id": "ID-456",
				"email":          "test@example.com"
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	recipient, err := client.Recipient.Get("ID-123")
	assert.Ok(t, err)

	expected := &Recipient{
		Id:            "ID-123",
		ApplicationId: "ID-456",
		Email:         "test@example.com",
	}
	assert.Equals(t, expected, recipient)

}
