package subscriptions

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToSubscriptionListQuery() (string, error)
}

// ListOpts params
type ListOpts struct {
	// Limit instructs List to refrain from sending excessively large lists of queues
	Limit int `q:"limit,omitempty"`

	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker,omitempty"`
}

func (opts ListOpts) ToSubscriptionListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListMessages instructs OpenStack to provide a list of messages.
func List(client *gophercloud.ServiceClient, queueName string, opts ListOptsBuilder) pagination.Pager {
	url := actionURL(client, queueName, "subscriptions")
	if opts != nil {
		query, err := opts.ToSubscriptionListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SubscriptionPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToQueueCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	Subscriber string                 `json:"subscriber"`
	TTL        int                    `json:"ttl"`
	Options    map[string]interface{} `json:"options"`
}

// ToQueueCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToQueueCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new queue.
func Create(client *gophercloud.ServiceClient, queueName string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToQueueCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	// Zaqar uses PUT instead of Create for creating queues
	_, r.Err = client.Post(updateURL(client, queueName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSubscriptionUpdateMap() (map[string]interface{}, error)
}

// UpdateClaimOpts implements UpdateOpts
type UpdateSubscriptionOpts struct {
	Subscriber string                 `json:"subscriber" required:"true"`
	TTL        int                    `json:"TTL,omitempty"`
	Options    map[string]interface{} `json:"options,omitempty"`
}

// ToSubscriptionUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateSubscriptionOpts) ToSubscriptionUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, queueName string, claimId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSubscriptionUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Patch(subscriptionURL(client, queueName, claimId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func Get(client *gophercloud.ServiceClient, queueName string, subscriptionId string) (r GetResult) {
	_, r.Err = client.Get(subscriptionURL(client, queueName, subscriptionId), &r.Body, nil)
	return
}

func Delete(client *gophercloud.ServiceClient, queueName string, subscriptionId string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, queueName, subscriptionId), nil)
	return
}
