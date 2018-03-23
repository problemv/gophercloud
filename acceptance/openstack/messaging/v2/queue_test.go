package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/queues"
	"github.com/gophercloud/gophercloud/pagination"
)

func TestCRUDQueues(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	createdQueue, err := queues.Get(client, createdQueueName, clientID).Extract()
	tools.PrintResource(t, createdQueue)

	updateOpts := queues.UpdateOpts{
		queues.UpdateQueueBody{
			Op:    "replace",
			Path:  "/metadata/_max_claim_count",
			Value: 15,
		},
		queues.UpdateQueueBody{
			Op:    "replace",
			Path:  "/metadata/description",
			Value: "Updated description for queues acceptance test.",
		},
	}

	updateResult, updateErr := queues.Update(client, createdQueueName, clientID, updateOpts).Extract()
	if updateErr != nil {
		t.Fatalf("Unable to update queue %s: %v", createdQueueName, updateErr)
	}

	tools.PrintResource(t, updateResult)
}

func TestListQueues(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	firstQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, firstQueueName, clientID)

	secondQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, secondQueueName, clientID)

	listOpts := queues.ListOpts{
		Limit:    10,
		Detailed: true,
	}

	pager := queues.List(client, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allQueues, err := queues.ExtractQueues(page)
		if err != nil {
			t.Fatalf("Unable to extract queues: %v", err)
		}

		for _, queue := range allQueues {
			tools.PrintResource(t, queue)
		}

		return true, nil
	})
}

func TestStatQueue(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	queueStats, err := queues.GetStats(client, createdQueueName, clientID).Extract()
	if err != nil {
		t.Fatalf("Unable to stat queue: %v", err)
	}

	tools.PrintResource(t, queueStats)
}

func TestShare(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}
	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	queueName, err := CreateQueue(t, client, clientID)
	if err != nil {
		t.Logf("Unable to create queue for share.")
	}
	defer DeleteQueue(t, client, queueName, clientID)

	t.Logf("Attempting to create share for queue: %s", queueName)
	share, shareErr := CreateShare(t, client, queueName, clientID)
	if shareErr != nil {
		t.Fatalf("Unable to create share: %v", shareErr)
	}

	tools.PrintResource(t, share)
}

func TestPurge(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}
	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	queueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, queueName, clientID)

	purgeOpts := queues.PurgeOpts{
		ResourceTypes: []string{
			"messages",
		},
	}

	t.Logf("Attempting to purge queue: %s", queueName)
	purgeErr := queues.Purge(client, queueName, "3381af92-2b9e-11e3-b191-71861300734d", purgeOpts).ExtractErr()
	if purgeErr != nil {
		t.Fatalf("Unable to purge queue %s: %v", queueName, purgeErr)
	}
}
