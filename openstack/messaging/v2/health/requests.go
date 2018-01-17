package health

import "github.com/gophercloud/gophercloud"

func GetPing(client *gophercloud.ServiceClient) (r GetResult) {
	_, r.Err = client.Get(pingURL(client), &r.Body, nil)
	return
}

func GetHealth(client *gophercloud.ServiceClient) (r GetResult) {
	_, r.Err = client.Get(healthURL(client), &r.Body, nil)
	return
}
