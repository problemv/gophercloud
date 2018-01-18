package queues

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
	DeadLetterQueue string `json:"_dead_letter_queue,omitempty"`

	DeadLetterQueueMessageTTL int `json:"_dead_letter_queue_messages_ttl,omitempty"`

	DefaultMessageDelay int `json:"_default_message_delay,omitempty"`

	DefaultMessageTTL int `json:"_default_message_ttl" required:"true"`

	Flavor string `json:"flavor,omitempty"`

	MaxClaimCount int `json:"_max_claim_count,omitempty"`

	MaxMessagesPostSize int `json:"_max_messages_post_size" required:"true"`

	Description string `json:"description,omitempty"`
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
	_, r.Err = client.Put(updateURL(client, queueName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
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

	// Specifies if showing the detailed information when querying queues
	Detailed bool `q:"detailed,omitempty"`
}

// ToQueueListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToQueueListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List instructs OpenStack to provide a list of queues.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToQueueListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return QueuePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// UpdateOpts implements UpdateOpts
type UpdateQueueOpts struct {
	Opts map[string]interface{} `json:"-"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToQueueUpdateMap() (map[string]interface{}, error)
}

// ToQueueUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateQueueOpts) ToQueueUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, queueName string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToQueueUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Patch(updateURL(client, queueName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201, 204},
	})
	return
}

func Get(client *gophercloud.ServiceClient, queueName string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, queueName), &r.Body, nil)
	return
}

// Delete deletes the specified queue.
func Delete(client *gophercloud.ServiceClient, queueName string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, queueName), nil)
	return
}

// GetStats returns statistics for the specified queue.
func GetStats(client *gophercloud.ServiceClient, queueName string) (r GetResult) {
	_, r.Err = client.Get(actionURL(client, queueName, "stats"), &r.Body, nil)
	return
}

type ShareOpts struct {
	Paths	[]string	`json:"paths,omitempty"`
	Methods []string	`json:"methods,omitempty"`
	Expires string		`json:"expires,omitempty"`
}

type ShareOptsBuilder interface {
	ToShareQueueMap() (map[string]interface{}, error)
}

func (opts ShareOpts) ToShareQueueMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Share(client *gophercloud.ServiceClient, queueName string, opts ShareOptsBuilder) (r PostResult) {
	b, err := opts.ToShareQueueMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Post(actionURL(client, queueName, "share"), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type PurgeOpts struct {
	ResourceTypes	[]string	`json:"resource_types,omitempty"`
}

type PurgeOptsBuilder interface {
	ToPurgeQueueMap() (map[string]interface{}, error)
}

func (opts ShareOpts) ToPurgeQueueMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Purge(client *gophercloud.ServiceClient, queueName string, opts ShareOptsBuilder) (r PostResult) {
	b, err := opts.ToShareQueueMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Post(actionURL(client, queueName, "purge"), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
