package client

import (
	"github.com/carwow/terraform-provider-hirefire/testing/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTimeRange(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equals(t, req.URL.String(), "/time_ranges/ID-123")
		rw.Write([]byte(`{
			"time_range": {
				"id":           "ID-123",
				"manager_id":   "ID-456",
				"from_minute":  15,
				"until_minute": 30,
				"minimum":      1,
				"maximum":      3,
				"position":     0,
				"monday":       true,
				"tuesday":      true,
				"wednesday":    true,
				"thursday":     true,
				"friday":       true,
				"saturday":     false,
				"sunday":       false
			}
		}`))
	}))
	defer server.Close()

	client := New("secret")
	client.URL = server.URL + "/"

	timeRange, err := client.TimeRange.Get("ID-123")
	assert.Ok(t, err)

	expected := &TimeRange{
		Id:          "ID-123",
		ManagerId:   "ID-456",
		FromMinute:  15,
		UntilMinute: 30,
		Minimum:     1,
		Maximum:     3,
		Position:    0,
		Monday:      true,
		Tuesday:     true,
		Wednesday:   true,
		Thursday:    true,
		Friday:      true,
		Saturday:    false,
		Sunday:      false,
	}
	assert.Equals(t, expected, timeRange)

}
