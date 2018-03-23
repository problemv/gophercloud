/*
Package queues provides information and interaction with the queues through
the OpenStack Messaging(Zaqar) service.

Lists all queues and creates, shows information for updates, deletes, and actions on a queue.

Example to List Queues

	listOpts := queues.ListOpts{
		Limit: 10,
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	pager := queues.List(client, clientID, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		queues, err := queues.ExtractQueues(page)
		if err != nil {
			panic(err)
		}

		for _, queue := range queues {
			fmt.Printf("%+v\n", queue)
		}

		return true, nil
	})

Example to Create a Queue

	createOpts := queues.CreateOpts{
		MaxMessagesPostSize: 262143,
		DefaultMessageTTL: 3700,
		DefaultMessageDelay: 25,
		DeadLetterQueueMessageTTL: 3500,
		MaxClaimCount: 10,
		Description: "Test queue for Gophercloud acceptance tests.",
	}

	err := queues.Create(client, "My Queue", clientID, createOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Update a Queue

	updateOpts := queues.UpdateOpts{
		queues.UpdateQueueBody{
			Op:    "replace",
			Path:  "/metadata/_max_claim_count",
			Value: 15,
		},
		queues.UpdateQueueBody{
			Op: "replace",
			Path: "/metadata/description",
			Value: "Updated description for queues acceptance test.",
		},
	}

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	updateResult, err := queues.Update(client, queueName, clientID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Queue

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	err := queues.Delete(client, queueName, clientID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Get a Queue

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	queue, err := queues.Get(client, queueName, clientID).Extract()
	if err != nil {
		panic(err)
	}

Example to GetStat a Queue

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	queueStats, err := queues.GetStats(client, queueName, clientID).Extract()
	if err != nil {
		panic(err)
	}

Example to Share a queue

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	shareOpts := queues.ShareOpts{
		Paths: []string{"messages"},
		Methods: []string{"POST"},
	}

	queueShare, err := queues.Share(client, queueName, clientID, shareOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Purge a queue

	clientID := "3381af92-2b9e-11e3-b191-71861300734d"

	purgeOpts := queues.PurgeOpts{
		ResourceTypes: []string{
			"messages",
		},
	}

	err := queues.Purge(client, queueName, "3381af92-2b9e-11e3-b191-71861300734d", purgeOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

*/
package queues
