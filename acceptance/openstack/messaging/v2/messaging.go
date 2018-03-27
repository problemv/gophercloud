package v2

import (
	"strings"
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/claims"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/messages"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/queues"
	"github.com/gophercloud/gophercloud/pagination"
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

func CreateMessage(t *testing.T, client *gophercloud.ServiceClient, queueName string, clientID string) (messages.ResourceList, error) {
	t.Logf("Attempting to add message to Queue: %s", queueName)
	createOpts := messages.CreateOpts{
		Messages: []messages.Messages{
			{
				TTL:  300,
				Body: map[string]interface{}{"Key": tools.RandomString("ACPTTEST", 8)},
			},
		},
	}

	resource, err := messages.Create(client, queueName, clientID, createOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to add message to queue %s: %v", queueName, err)
	} else {
		t.Logf("Successfully added message to queue: %s", queueName)
	}

	return resource, err
}

func ListMessages(t *testing.T, client *gophercloud.ServiceClient, queueName string, clientID string) ([]messages.Message, error) {
	listOpts := messages.ListOpts{}
	var allMessages []messages.Message
	var listErr error

	t.Logf("Attempting to list messages on queue: %s", queueName)
	pager := messages.List(client, queueName, clientID, listOpts)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		allMessages, listErr = messages.ExtractMessages(page)
		if listErr != nil {
			t.Fatalf("Unable to extract messages: %v", listErr)
		}

		for _, message := range allMessages {
			tools.PrintResource(t, message)
		}

		return true, nil
	})
	return allMessages, err
}

func CreateClaim(t *testing.T, client *gophercloud.ServiceClient, queueName string, clientID string) ([]claims.Messages, error) {
	createOpts := claims.CreateOpts{}
	createQueryOpts := claims.CreateQueryOpts{}

	t.Logf("Attempting to create claim on queue: %s", queueName)
	claimedMessages, err := claims.Create(client, queueName, clientID, createOpts, createQueryOpts).Extract()
	tools.PrintResource(t, claimedMessages)
	if err != nil {
		t.Fatalf("Unable to create claim: %v", err)
	}

	return claimedMessages, err
}

func GetClaim(t *testing.T, client *gophercloud.ServiceClient, queueName string, claimID string, clientID string) (*claims.Claim, error) {
	t.Logf("Attempting to get claim: %s", claimID)
	claim, err := claims.Get(client, queueName, claimID, clientID).Extract()
	if err != nil {
		t.Fatalf("Unable to get claim: %s", claimID)
	}

	return claim, err
}

func DeleteClaim(t *testing.T, client *gophercloud.ServiceClient, queueName string, claimID string, clientID string) error {
	t.Logf("Attempting to delete claim: %s", claimID)
	err := claims.Delete(client, queueName, claimID, clientID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to delete claim: %s", clientID)
	}
	t.Logf("Sucessfully deleted claim: %s", claimID)

	return err
}

func ExtractIDs(claim []claims.Messages) ([]string, []string) {
	var claimIDs []string
	var messageID []string

	for _, msg := range claim {
		parts := strings.Split(msg.Href, "?claim_id=")
		if len(parts) == 2 {
			pieces := strings.Split(parts[0], "/")
			if len(pieces) > 0 {
				messageID = append(messageID, pieces[len(pieces)-1])
			}
			claimIDs = append(claimIDs, parts[1])
		}
	}
	encountered := map[string]bool{}
	for v:= range claimIDs {
		encountered[claimIDs[v]] = true
	}

	var uniqueClaimIDs []string

	for key := range encountered {
		uniqueClaimIDs = append(uniqueClaimIDs, key)
	}
	return uniqueClaimIDs, messageID
}
