package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/health"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud/testhelper/fixture"
)

func TestGetPing(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixture.SetupHandler(t, "/ping", "GET", "", GetDSResp, 200)

	ds, err := health.GetPing(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &ExampleDatastore, ds)
}

func TestGetHealth(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixture.SetupHandler(t, "/health", "GET", "", GetDSResp, 200)

	ds, err := health.GetHealth(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &ExampleDatastore, ds)
}
