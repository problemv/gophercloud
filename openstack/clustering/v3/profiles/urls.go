package profiles

import "github.com/gophercloud/gophercloud"

var apiVersion = "v1"
var apiName = "profiles"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func idURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}

func profileURL(client *gophercloud.ServiceClient, id string) string {
	return idURL(client, id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}

func listURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}

func updateURL(client *gophercloud.ServiceClient, id string) string {
	return idURL(client, id)
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return idURL(client, id)
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return idURL(client, id)
}
