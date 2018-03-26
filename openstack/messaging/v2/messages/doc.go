/*
Package messages provides information and interaction with the messages through
the OpenStack Messaging(Zaqar) service.

Example to Create Messages
	createOpts := messages.CreateOpts{
		Messages: []messages.Messages{
			{
				TTL:   300,
				Delay: 20,
				Body: map[string]interface{}{
					"event": "BackupStarted",
					"bakcup_id": "c378813c-3f0b-11e2-ad92-7823d2b0f3ce",
				},
			},
			{
				Body: map[string]interface{}{
					"event": "BackupProgress",
					"current_bytes": "0",
					"total_bytes": "99614720",
				},
			},
		},
	}

	clientID := "fff121f5-c506-410a-a69e-2d73ef9cbdbd"
	queueName := "my_queue"

	resources, err := messages.Create(client, QueueName, clientID, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to List Messages

	listOpts := messages.ListOpts{}

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"
	queueName := "my_queue"

	pager := messages.List(client, queueName, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		allMessages, err := queues.ExtractQueues(page)
		if err != nil {
			panic(err)
		}

		for _, message := range allMessages {
			fmt.Printf("%+v\n", message)
		}

		return true, nil
	})


Example to Get a set of Messages

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"
	queueName := "my_queue"

	getMessageOpts := messages.GetMessagesOpts{
		Ids: "123456",
	}

	messagesList, err := messages.GetMessages(client, createdQueueName, clientID, getMessageOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a set of Messages

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"
	queueName := "my_queue"

	deleteMessagesOpts := messages.DeleteMessageOpts{
		Ids: []string{"9988776655"},
	}

	err := messages.DeleteMessages(client, queueName, clientID, deleteMessagesOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to get a singular Message

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"
	queueName := "my_queue"
	messageID := "123456"

	message, err := messages.Get(client, queueName, clientID, ClientID).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a singular Message

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"
	queueName := "my_queue"
	messageID := "123456"

	deleteOpts := messages.DeleteOpts{
		Claim: "12345",
	}

	err := messages.Delete(client), queueName, messageID, clientID, deleteOpts).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package messages
