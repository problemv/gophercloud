package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/claims"
)

func TestCRUDClaim(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	clientID = "3381af92-2b9e-11e3-b191-71861300734d"
	for i := 0; i < 3; i++ {
		CreateMessage(t, client, createdQueueName, clientID)
	}

	clientID = "3381af92-2b9e-11e3-b191-7186130073dd"
	claimedMessages, err := CreateClaim(t, client, createdQueueName, clientID)
	claimIDs, _ := ExtractIDs(claimedMessages)

	tools.PrintResource(t, claimedMessages)

	updateOpts := claims.UpdateClaimOpts{
		TTL: 600,
		Grace: 500,
	}

	for _, claimID := range claimIDs {
		t.Logf("Attempting to update claim: %s", claimID)
		updateErr := claims.Update(client, createdQueueName, claimID, clientID, updateOpts).ExtractErr()
		if updateErr != nil {
			t.Fatalf("Unable to update claim %s: %v", claimID,  err)
		} else {
			t.Logf("Successfully updated claim: %s", claimID)
		}

		updatedClaim, getErr := GetClaim(t, client, createdQueueName, claimID, clientID)
		if getErr != nil {
			t.Fatalf("Unable to retrieve claim %s: %v", claimID, getErr)
		}

		tools.PrintResource(t, updatedClaim)
	}

	for _, claimID := range claimIDs {
		DeleteClaim(t, client, createdQueueName, claimID, clientID)
	}
}
