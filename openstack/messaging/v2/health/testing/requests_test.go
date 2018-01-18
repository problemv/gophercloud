package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/health"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"fmt"
	"net/http"
)

func TestGetPing(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	var MockResp =`{"catalog_reachable": true}`

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, MockResp)
	})

	_, err := health.GetPing(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
}

func TestGetHealth(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	var MockResp =`{"catalog_reachable": true}`

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, MockResp)
	})

	expected := &health.Health{CatalogReachable:	true,}

	healthRet, err := health.GetHealth(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expected, healthRet)
}
