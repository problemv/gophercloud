package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/pools"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListPools(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
		{
  	"pools": [
		{
	      	"href": "/v2/pools/test_pool1",
	      	"group": "poolgroup",
	      	"name": "test_pool1",
	      	"weight": 60,
	      	"uri": "mongodb://192.168.1.10:27017"
	    }
	  ]
		}`)
	})

	count := 0

	pools.List(fake.ServiceClient(), pools.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := pools.ExtractPools(page)
		if err != nil {
			t.Errorf("Failed to extract nodes: %v", err)
			return false, err
		}

		expected := []pools.Pool{
			{
				Href:   "/v2/pools/test_pool1",
				Group:  "poolgroup",
				Name:   "test_pool1",
				Weight: 60,
				URI:    "mongodb://192.168.1.10:27017",
			},
		}

		th.AssertDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}
