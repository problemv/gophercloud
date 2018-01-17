package claims

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"
var apiName = "queues"

func commonURL(client *gophercloud.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims")
}

func idURL(client *gophercloud.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims")
}

func actionURL(client *gophercloud.ServiceClient, queueName string, action string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, action)
}

func claimURL(client *gophercloud.ServiceClient, queueName string, claimId string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims", claimId)
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

func deleteURL(client *gophercloud.ServiceClient, queueName string) string {
	return idURL(client, queueName)
}

func getURL(client *gophercloud.ServiceClient, queueName string) string {
	return idURL(client, queueName)
}
