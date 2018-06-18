/*
Package claims provides information and interaction with the Zaqar API
claims resource for the OpenStack Messaging service.

Example to Create a Claim on a specified Zaqar queue

	createOpts := claims.CreateOpts{
		TTL:		60,
		Grace:		120,
		Limit: 		20,
	}

	queueName := "my_queue"

	messages, err := claims.Create(messagingClient, queueName, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to get a claim for a specified Zaqar queue

	queueName := "my_queue"

	claim, err := claims.Get(messagingClient, queueName, claimID).Extract()
	if err != nil {
		panic(err)
	}

*/
package claims
