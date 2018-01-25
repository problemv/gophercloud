package pools

import "github.com/gophercloud/gophercloud"

var apiVersion = "v2"
var apiName = "pools"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func poolURL(client *gophercloud.ServiceClient, poolName string) string {
	return client.ServiceURL(apiVersion, apiName, poolName)
}
