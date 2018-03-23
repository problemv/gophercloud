package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/queues"
)

func CreateQueue(t *testing.T, client *gophercloud.ServiceClient, clientID string) (string, error) {
	queueName := tools.RandomString("ACPTTEST", 5)

	t.Logf("Attempting to create Queue: %s", queueName)

	createOpts := queues.CreateOpts{
		MaxMessagesPostSize:       262143,
		DefaultMessageTTL:         3700,
		DefaultMessageDelay:       25,
		DeadLetterQueueMessageTTL: 3500,
		MaxClaimCount:             10,
		Description:               "Test queue for Gophercloud acceptance tests.",
	}

	createErr := queues.Create(client, queueName, clientID, createOpts).ExtractErr()
	if createErr != nil {
		t.Fatalf("Unable to create queue: %v", createErr)
	}

	t.Logf("Attempting to get Queue: %s", queueName)
	_, err := queues.Get(client, queueName, clientID).Extract()
	if err != nil {
		return queueName, err
	}

	t.Logf("Created Queue: %s", queueName)
	return queueName, nil
}

func CreateShare(t *testing.T, client *gophercloud.ServiceClient, queueName string, clientID string) (queues.QueueShare, error) {
	t.Logf("Attempting to create share for queue: %s", queueName)

	shareOpts := queues.ShareOpts{
		Paths:   []string{"messages"},
		Methods: []string{"POST"},
	}

	share, err := queues.Share(client, queueName, clientID, shareOpts).Extract()

	return share, err
}

func DeleteQueue(t *testing.T, client *gophercloud.ServiceClient, queueName string, clientID string) {
	t.Logf("Attempting to delete Queue: %s", queueName)
	err := queues.Delete(client, queueName, clientID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to delete queue %s: %v", queueName, err)
	}

	t.Logf("Deleted queue: %s", queueName)
}
