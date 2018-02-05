package messages

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToMessageCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	TTL int `json:"ttl,omitempty"`

	Delay int `json:"delay,omitempty"`

	Body map[string]interface{} `json:"body" required:"true"`
}

// ToMessageCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToMessageCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "messages")
}

func Create(client *gophercloud.ServiceClient, queueName string, clientId string, opts CreateOpts) (r CreateResult) {
	b, err := opts.ToMessageCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, queueName, "messages"), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{201},
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToMessageListQuery() (string, error)
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
func ListMessages(client *gophercloud.ServiceClient, queueName string, clientId string, opts ListOptsBuilder) pagination.Pager {
	headers := map[string]string{"Client-ID": clientId}
	url := actionURL(client, queueName, "messages")
	if opts != nil {
		query, err := opts.ToMessageListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pager := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return MessagePage{pagination.LinkedPageBase{PageResult: r}}
	})
	pager.Headers = headers
	return pager
}

func DeleteMessages(client *gophercloud.ServiceClient, queueName string, clientId string) (r DeleteResult) {
	_, r.Err = client.Delete(actionURL(client, queueName, "messages"), &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}

func GetMessage(client *gophercloud.ServiceClient, queueName string, messageId string, clientId string) (r GetResult) {
	_, r.Err = client.Get(messageURL(client, queueName, messageId), &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}

func DeleteMessage(client *gophercloud.ServiceClient, queueName string, messageId string, clientId string) (r DeleteResult) {
	_, r.Err = client.Delete(messageURL(client, queueName, messageId), &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}
