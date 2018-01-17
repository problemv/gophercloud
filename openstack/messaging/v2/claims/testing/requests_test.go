package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/claims"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud/testhelper/fixture"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	fixture.SetupHandler(t, "/queues/fake_queue/claims/12345", "GET", "", GetDSResp, 200)

	ds, err := claims.Get(fake.ServiceClient(), "fake_queue", "12345").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &ExampleDatastore, ds)
}
