package messages

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"
var apiName = "queues"

func actionURL(client *gophercloud.ServiceClient, queueName string, action string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, action)
}

func getURL(client *gophercloud.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages")
}

func messageURL(client *gophercloud.ServiceClient, queueName string, messageID string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages", messageID)
}
