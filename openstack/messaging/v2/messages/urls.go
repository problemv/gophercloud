package messages

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"
var apiName = "queues"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func idURL(client *gophercloud.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName)
}

func actionURL(client *gophercloud.ServiceClient, queueName string, action string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, action)
}

func messageURL(client *gophercloud.ServiceClient, queueName string, messageId string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages", messageId)
}

func createURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}

func listURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
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
