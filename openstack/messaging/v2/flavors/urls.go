package flavors

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"
var apiName = "flavors"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func flavorURL(client *gophercloud.ServiceClient, poolName string) string {
	return client.ServiceURL(apiVersion, apiName, poolName)
}
