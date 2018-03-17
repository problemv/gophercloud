package claims

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"
var apiName = "queues"

func actionURL(client *gophercloud.ServiceClient, queueName string, action string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, action)
}

func claimURL(client *gophercloud.ServiceClient, queueName string, claimId string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims", claimId)
}
