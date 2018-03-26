package v2

import (
	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/messages"
	"github.com/gophercloud/gophercloud/pagination"
	"testing"
)

func TestCreateMessages(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	clientID = "3381af92-2b9e-11e3-b191-71861300734d"
	CreateMessage(t, client, createdQueueName, clientID)
}

func TestListMessages(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-718613007343"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	for i := 0; i < 3; i++ {
		CreateMessage(t, client, createdQueueName, clientID)
	}

	// Use a different clientID in order to see messages on the Queue
	clientID = "3381af92-2b9e-11e3-b191-71861300734d"
	listOpts := messages.ListOpts{}

	pager := messages.List(client, createdQueueName, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allMessages, err := messages.ExtractMessages(page)
		if err != nil {
			t.Fatalf("Unable to extract messages: %v", err)
		}

		for _, message := range allMessages {
			tools.PrintResource(t, message)
		}

		return true, nil
	})
}

func TestGetMessages(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-718613007343"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	CreateMessage(t, client, createdQueueName, clientID)
	CreateMessage(t, client, createdQueueName, clientID)

	// Use a different clientID in order to see messages on the Queue
	clientID = "3381af92-2b9e-11e3-b191-71861300734d"
	listOpts := messages.ListOpts{}

	var messageIDs []string

	pager := messages.List(client, createdQueueName, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allMessages, err := messages.ExtractMessages(page)
		if err != nil {
			t.Fatalf("Unable to extract messages: %v", err)
		}

		for _, message := range allMessages {
			messageIDs = append(messageIDs, message.ID)
		}

		return true, nil
	})

	getMessageOpts := messages.GetMessagesOpts{
		Ids: messageIDs,
	}
	t.Logf("Attempting to get messages from queue %s with ids: %v", createdQueueName, messageIDs)
	messagesList, err := messages.GetMessages(client, createdQueueName, clientID, getMessageOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to get messages from queue: %s", createdQueueName)
	}

	tools.PrintResource(t, messagesList)
}

func TestDeleteMessagesIDs(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-718613007343"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	CreateMessage(t, client, createdQueueName, clientID)
	CreateMessage(t, client, createdQueueName, clientID)

	// Use a different clientID in order to see messages on the Queue
	clientID = "3381af92-2b9e-11e3-b191-71861300734d"
	listOpts := messages.ListOpts{}

	var messageIDs []string

	pager := messages.List(client, createdQueueName, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allMessages, err := messages.ExtractMessages(page)
		if err != nil {
			t.Fatalf("Unable to extract messages: %v", err)
		}

		for _, message := range allMessages {
			messageIDs = append(messageIDs, message.ID)
		}

		return true, nil
	})

	deleteOpts := messages.DeleteMessageOpts{
		Ids: messageIDs,
	}

	t.Logf("Attempting to delete messages: %v", messageIDs)
	deleteErr := messages.DeleteMessages(client, createdQueueName, clientID, deleteOpts).ExtractErr()
	if deleteErr != nil {
		t.Fatalf("Unable to delete messages: %v", deleteErr)
	}

	t.Logf("Attempting to list messages.")
	messageList, err := ListMessages(t, client, createdQueueName, clientID)

	if len(messageList) > 0 {
		t.Fatalf("Did not delete all specified messages in the queue.")
	}
}

func TestDeleteMessagesPop(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-718613007343"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	for i := 0; i < 5; i++ {
		CreateMessage(t, client, createdQueueName, clientID)
	}

	// Use a different clientID in order to see messages on the Queue
	clientID = "3381af92-2b9e-11e3-b191-71861300734d"

	messageList, err := ListMessages(t, client, createdQueueName, clientID)

	messagesNumber := len(messageList)
	popNumber := 3

	deleteOpts := messages.DeleteMessageOpts{
		Pop: popNumber,
	}

	deleteErr := messages.DeleteMessages(client, createdQueueName, clientID, deleteOpts).ExtractErr()
	if deleteErr != nil {
		t.Fatalf("Unable to delete messages: %v", deleteErr)
	}

	messageList, err = ListMessages(t, client, createdQueueName, clientID)
	if len(messageList) != messagesNumber-popNumber {
		t.Fatalf("Unable to delete specified number of messages.")
	}

}

func TestGetMessage(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-718613007343"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	CreateMessage(t, client, createdQueueName, clientID)

	// Use a different clientID in order to see messages on the Queue
	clientID = "3381af92-2b9e-11e3-b191-71861300734d"
	listOpts := messages.ListOpts{}

	var messageIDs []string

	pager := messages.List(client, createdQueueName, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allMessages, err := messages.ExtractMessages(page)
		if err != nil {
			t.Fatalf("Unable to extract messages: %v", err)
		}

		for _, message := range allMessages {
			messageIDs = append(messageIDs, message.ID)
		}

		return true, nil
	})

	for _, messageID := range messageIDs {
		t.Logf("Ateempting to get message from queue %s: %s", createdQueueName, messageID)
		message, getErr := messages.Get(client, createdQueueName, messageID, clientID).Extract()
		if getErr != nil {
			t.Fatalf("Unable to get message from queue %s: %s", createdQueueName, messageID)
		}
		tools.PrintResource(t, message)
	}
}

func TestDeleteMessage(t *testing.T) {
	client, err := clients.NewMessagingV2Client()
	if err != nil {
		t.Fatalf("Unable to create a messaging service client: %v", err)
	}

	clientID := "3381af92-2b9e-11e3-b191-718613007343"

	createdQueueName, err := CreateQueue(t, client, clientID)
	defer DeleteQueue(t, client, createdQueueName, clientID)

	CreateMessage(t, client, createdQueueName, clientID)

	// Use a different clientID in order to see messages on the Queue
	clientID = "3381af92-2b9e-11e3-b191-71861300734d"
	listOpts := messages.ListOpts{}

	var messageIDs []string

	pager := messages.List(client, createdQueueName, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allMessages, err := messages.ExtractMessages(page)
		if err != nil {
			t.Fatalf("Unable to extract messages: %v", err)
		}

		for _, message := range allMessages {
			messageIDs = append(messageIDs, message.ID)
		}

		return true, nil
	})

	for _, messageID := range messageIDs {
		t.Logf("Ateempting to delete message from queue %s: %s", createdQueueName, messageID)
		deleteOpts := messages.DeleteOpts{}
		deleteErr := messages.Delete(client, createdQueueName, messageID, clientID, deleteOpts).ExtractErr()
		if deleteErr != nil {
			t.Fatalf("Unable to delete message from queue %s: %s", createdQueueName, messageID)
		} else {
			t.Logf("Successfully deleted message: %s", messageID)
		}
	}
}
