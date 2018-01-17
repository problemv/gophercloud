package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/messages"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListMessages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
		{
    "messages": [
        {
            "body": {
                "current_bytes": "0",
                "event": "BackupProgress",
                "total_bytes": "99614720"
            },
            "age": 482,
            "href": "/v2/queues/beijing/messages/578edfe6508f153f256f717b",
            "id": "578edfe6508f153f256f717b",
            "ttl": 3600
        }
	]
		}`)
	})

	count := 0

	messages.ListMessages(fake.ServiceClient(), messages.ListOpts{}, "fake_queue").EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := messages.ExtractMessages(page)
		if err != nil {
			t.Errorf("Failed to extract nodes: %v", err)
			return false, err
		}

		expected := []messages.Message{
			{
				Body:	map[string]interface{}{},
				Age:	482,
				Href:	"/v2/queues/beijing/messages/578edfe6508f153f256f717b",
				ID:		"578edfe6508f153f256f717b",
				TTL:	3600,
			},
		}

		th.AssertDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}
