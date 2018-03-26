package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/queues"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	listOpts := queues.ListOpts{}

	count := 0
	err := queues.List(fake.ServiceClient(), ClientID, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := queues.ExtractQueues(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedQueueSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	createOpts := queues.CreateOpts{
		MaxMessagesPostSize:       262144,
		DefaultMessageTTL:         3600,
		DefaultMessageDelay:       30,
		DeadLetterQueue:           "dead_letter",
		DeadLetterQueueMessageTTL: 3600,
		MaxClaimCount:             10,
		Description:               "Queue for unit testing.",
	}

	err := queues.Create(fake.ServiceClient(), QueueName, ClientID, createOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	updateOpts := queues.UpdateOpts{
		queues.UpdateQueueBody{
			Op:    "replace",
			Path:  "/metadata/_max_claim_count",
			Value: 10,
		},
	}
	updatedQueueResult := queues.QueueDetails{
		MaxClaimCount: 10,
	}

	actual, err := queues.Update(fake.ServiceClient(), QueueName, ClientID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, updatedQueueResult, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := queues.Delete(fake.ServiceClient(), QueueName, ClientID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := queues.Get(fake.ServiceClient(), QueueName, ClientID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, QueueDetails, actual)
}

func TestGetStat(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetStatsSuccessfully(t)

	actual, err := queues.GetStats(fake.ServiceClient(), QueueName, ClientID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedStats, actual)
}

func TestShare(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleShareSuccessfully(t)

	shareOpts := queues.ShareOpts{
		Paths:   []string{"messages", "claims", "subscriptions"},
		Methods: []string{"GET", "POST", "PUT", "PATCH"},
		Expires: "2016-09-01T00:00:00",
	}

	actual, err := queues.Share(fake.ServiceClient(), QueueName, ClientID, shareOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedShare, actual)
}

func TestPurge(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandlePurgeSuccessfully(t)

	purgeOpts := queues.PurgeOpts{
		ResourceTypes: []string{"messages", "subscriptions"},
	}

	err := queues.Purge(fake.ServiceClient(), QueueName, ClientID, purgeOpts).ExtractErr()
	th.AssertNoErr(t, err)
}
