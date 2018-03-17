package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/claims"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud.bak/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := claims.Get(fake.ServiceClient(), QueueName, ClaimID, ClientID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstClaim, actual)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	createOpts := claims.CreateOpts{
		TTL:   3600,
		Grace: 3600,
	}

	CreateQueryOpts := claims.CreateQueryOpts{
		Limit: 10,
	}

	actual, err := claims.Create(fake.ServiceClient(), QueueName, ClientID, createOpts, CreateQueryOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedClaim, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	updateOpts := claims.UpdateClaimOpts{
		Grace: 1600,
		TTL:   1200,
	}

	err := claims.Update(fake.ServiceClient(), QueueName, ClaimID, ClientID, updateOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestDelete(t *testing.T){
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := claims.Delete(client.ServiceClient(), QueueName, ClaimID, ClientID).ExtractErr()
	th.AssertNoErr(t, err)
}
