package messages

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToQueueCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	TTL int `json:"ttl,omitempty"`

	Delay int `json:"delay,omitempty"`

	Body string `json:"body" required:"true"`
}

// ToMessageCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToMessageCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "messages")
}

func Create(client *gophercloud.ServiceClient, queueName string, messageId string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToQueueCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(messageURL(client, queueName, messageId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToQueueListQuery() (string, error)
}

// ListOpts params
type ListOpts struct {
	// Limit instructs List to refrain from sending excessively large lists of queues
	Limit int `q:"limit,omitempty"`

	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker,omitempty"`

	// Indicate if the messages can be echoed back to the client that posted them.
	Echo bool `q:"echo,omitempty"`

	// Indicate if the messages list should include the claimed messages.
	IncludeClaimed bool `q:"include_claimed,omitempty"`

	//Indicate if the messages list should include the delayed messages.
	IncludeDelayed bool `q:"include_delayed,omitempty"`
}

func (opts ListOpts) ToMessageListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListMessages instructs OpenStack to provide a list of messages.
func ListMessages(client *gophercloud.ServiceClient, opts ListOptsBuilder, queueName string) pagination.Pager {
	url := actionURL(client, queueName, "messages")
	if opts != nil {
		query, err := opts.ToQueueListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return MessagePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func DeleteMessages(client *gophercloud.ServiceClient, queueName string) (r DeleteResult) {
	_, r.Err = client.Delete(actionURL(client, queueName, "messages"), nil)
	return
}

func GetMessage(client *gophercloud.ServiceClient, queueName string, messageId string) (r GetResult) {
	_, r.Err = client.Get(messageURL(client, queueName, messageId), &r.Body, nil)
	return
}

func DeleteMessage(client *gophercloud.ServiceClient, queueName string, messageId string) (r DeleteResult) {
	_, r.Err = client.Delete(messageURL(client, queueName, messageId), nil)
	return
}
