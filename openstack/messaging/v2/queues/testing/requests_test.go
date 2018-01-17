package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/queues"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListQueues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
		{
   "queues": [
      {
         "href":"/v2/queues/beijing",
         "name":"beijing"
      }
   ]
		}`)
	})

	count := 0

	queues.List(fake.ServiceClient(), queues.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := queues.ExtractQueues(page)
		if err != nil {
			t.Errorf("Failed to extract nodes: %v", err)
			return false, err
		}

		expected := []queues.Queue{
			{
				Href:		"/v2/queues/beijing",
				Name:		"beijing",
			},
		}

		th.AssertDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}
