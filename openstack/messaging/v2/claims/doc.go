/*
Package claims provides information and interaction with the Zaqar API
claims resource for the OpenStack Messaging service.

Example to Create a Claim on a specified Zaqar queue

	createOpts := claims.CreateOpts{
		TTL:	60
		Grace:	120
	}

	claimQueryOpts := claims.CreateQueryOpts{
		Limit: 20,
	}

	queueName := "my_queue"
	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	messages, err := claims.Create(messagingClient, queueName, clientID, createOpts, claimQueryOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to get a claim for a specified Zaqar queue

	queueName := "my_queue"
	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	claim, err := claims.Get(messagingClient, queueName, clientID, claimID).Extract()
	if err != nil {
		panic(err)
	}

Example to update a claim for a specified Zaqar queue

	updateOpts := claims.UpdateClaimOpts{
		TTL: 600
		Grace: 1200
	}

	queueName := "my_queue"
	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	err := claims.Update(messagingClient, queueName, claimID, clientID, updateOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to delete a claim for a specified Zaqar queue

	queueName := "my_queue"
	clientID := "3381af92-2b9e-11e3-b191-71861300734c"

	err := claims.Delete(messagingClient, queueName, clientID, claimID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package claims
