package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/flavors"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud/testhelper/fixture"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixture.SetupHandler(t, "/flavors/fake_flavor", "GET", "", GetDSResp, 200)

	ds, err := flavors.Get(fake.ServiceClient(), "fake_flavor").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &ExampleDatastore, ds)
}
