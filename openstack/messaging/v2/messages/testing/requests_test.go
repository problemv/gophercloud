package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/messages"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	createOpts := messages.CreateOpts{
		Messages: []messages.Messages{
			{
				TTL:   300,
				Delay: 20,
				Body: map[string]interface{}{
					"event":     "BackupStarted",
					"bakcup_id": "c378813c-3f0b-11e2-ad92-7823d2b0f3ce",
				},
			},
			{
				Body: map[string]interface{}{
					"event":         "BackupProgress",
					"current_bytes": "0",
					"total_bytes":   "99614720",
				},
			},
		},
	}

	actual, err := messages.Create(fake.ServiceClient(), QueueName, ClientID, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedResources, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	listOpts := messages.ListOpts{}

	count := 0
	err := messages.List(fake.ServiceClient(), QueueName, ClientID, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := messages.ExtractMessages(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedMessagesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestGetMessages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetMessagesSuccessfully(t)

	getMessagesOpts := messages.GetMessagesOpts{
		Ids: []string{"9988776655"},
	}

	actual, err := messages.GetMessages(fake.ServiceClient(), QueueName, ClientID, getMessagesOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExpectedMessagesSet, actual)
}

func TestDeleteMessages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteMessagesSuccessfully(t)

	deleteMessagesOpts := messages.DeleteMessageOpts{
		Ids: []string{"9988776655"},
	}

	err := messages.DeleteMessages(fake.ServiceClient(), QueueName, ClientID, deleteMessagesOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestGetMessage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := messages.Get(fake.ServiceClient(), QueueName, MessageID, ClientID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, FirstMessage, actual)
}

func TestDeleteMessage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	deleteOpts := messages.DeleteOpts{
		Claim: "12345",
	}

	err := messages.Delete(fake.ServiceClient(), QueueName, MessageID, ClientID, deleteOpts).ExtractErr()
	th.AssertNoErr(t, err)
}
