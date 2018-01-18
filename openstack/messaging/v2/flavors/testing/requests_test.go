package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/flavors"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"fmt"
	"net/http"
)


func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	var MockResp =`
	{
  		"href": "/v2/flavors/testflavor",
  		"capabilities": [
    		"FIFO",
    		"CLAIMS",
    		"DURABILITY",
    		"AOD",
    		"HIGH_THROUGHPUT"
  		],
  		"pool_group": "testgroup",
  		"name": "testflavor"
	}`

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, MockResp)
	})

	expected := &flavors.Flavor{
		Href:			"/v2/flavors/testflavor",
		Capabilities:	[]string {"FIFO", "CLAIMS", "DURABILITY", "AOD", "HIGH_THROUGHPUT"},
		PoolGroup:		"testgroup",
		Name:			"testflavor",
	}

	flavor, err := flavors.Get(fake.ServiceClient(), "fake_flavor").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expected, flavor)
}
