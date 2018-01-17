package health

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"

func pingURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, "ping")
}

func healthURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, "health")
}
