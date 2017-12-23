/*
Package webhooks enables triggers an action represented by a webhook from the OpenStack
Clustering Service.

Example to Trigger webhook action

	webhooks.TriggerOpts{
		V: "1.0",
	}

	trigger, err := webhooks.Trigger(webhookClient, TriggerOpts)
	if err != nil {
		panic(err)
	}

	action, err := webhooks.ExtractAction(trigger)
	if err != nil {
		panic(err)
	}

  fmt.Printf("%+v\n", action)
*/
package webhooks
