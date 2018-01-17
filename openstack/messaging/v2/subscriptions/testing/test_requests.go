package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/subscriptions"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListSubscriptions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
		{
  	"subscriptions": [
    	{
      		"age": 13,
      		"id": "57692aa63990b48c644bb7e5",
      		"subscriber": "http://10.229.49.117:5678",
      		"source": "test",
      		"ttl": 360,
      		"options": {}
    	}
  	]
		}`)
	})

	count := 0

	subscriptions.List(fake.ServiceClient(), subscriptions.ListOpts{}, "fake_queue").EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := subscriptions.ExtractSubscriptions(page)
		if err != nil {
			t.Errorf("Failed to extract nodes: %v", err)
			return false, err
		}

		expected := []subscriptions.Subscription{
			{
				Age:		13,
				ID:			"57692aa63990b48c644bb7e5",
				Subscriber: "http://10.229.49.117:5678",
				Source:		"test",
				TTL: 		360,
				Options: 	map[string]interface{}{},
			},
		}

		th.AssertDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}
