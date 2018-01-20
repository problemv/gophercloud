package subscriptions

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"
var apiName = "queues"

func commonURL(client *gophercloud.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "subscriptions")
}

func idURL(client *gophercloud.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "subscriptions")
}

func actionURL(client *gophercloud.ServiceClient, queueName string, action string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, action)
}

func subscriptionURL(client *gophercloud.ServiceClient, queueName string, subscriptionsId string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "subscriptions", subscriptionsId)
}

func createURL(client *gophercloud.ServiceClient, queueName string) string {
	return commonURL(client, queueName)
}

func listURL(client *gophercloud.ServiceClient, queueName string) string {
	return commonURL(client, queueName)
}

func updateURL(client *gophercloud.ServiceClient, queueName string) string {
	return idURL(client, queueName)
}

func deleteURL(client *gophercloud.ServiceClient, queueName string, subscriptionId string) string {
	return subscriptionURL(client, queueName, subscriptionId)
}

func getURL(client *gophercloud.ServiceClient, queueName string) string {
	return idURL(client, queueName)
}
