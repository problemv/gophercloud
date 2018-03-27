package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/subscriptions"
	"github.com/gophercloud/gophercloud/pagination"
)

func TestListSubscriptions(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	createdQueueName, err := CreateQueue(t, client, clientID)

	listOpts := subscriptions.ListOpts{
		Limit: 1,
	}

	pager := subscriptions.List(client, createdQueueName, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allSubscriptions, err := subscriptions.ExtractSubscriptions(page)
		if err != nil {
			t.Fatalf("Unable to extract subscriptions: %v", err)
		}

		for _, subscription := range allSubscriptions {
			tools.PrintResource(t, subscription)
		}

		return true, nil
	})
}
