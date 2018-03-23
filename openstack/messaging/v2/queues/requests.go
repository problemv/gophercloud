package queues

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToQueueListQuery() (string, error)
}

// ListOpts params to be used with List
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
func List(client *gophercloud.ServiceClient, clientID string, opts ListOptsBuilder) pagination.Pager {
	headers := map[string]string{"Client-ID": clientID}
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToQueueListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pager := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return QueuePage{pagination.LinkedPageBase{PageResult: r}}

	})
	pager.Headers = headers
	return pager
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToQueueCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies the queue creation parameters.
type CreateOpts struct {
	//The target incoming messages will be moved to when a message can’t
	// processed successfully after meet the max claim count is met.
	DeadLetterQueue string `json:"_dead_letter_queue,omitempty"`

	// The new TTL setting for messages when moved to dead letter queue.
	DeadLetterQueueMessageTTL int `json:"_dead_letter_queue_messages_ttl,omitempty"`

	// The delay of messages defined for a queue. When the messages send to
	// the queue, it will be delayed for some times and means it can not be
	// claimed until the delay expired.
	DefaultMessageDelay int `json:"_default_message_delay,omitempty"`

	// The default TTL of messages defined for a queue, which will effect for
	// any messages posted to the queue.
	DefaultMessageTTL int `json:"_default_message_ttl" required:"true"`

	// The flavor name which can tell Zaqar which storage pool will be used
	// to create the queue.
	Flavor string `json:"_flavor,omitempty"`

	// The max number the message can be claimed.
	MaxClaimCount int `json:"_max_claim_count,omitempty"`

	// The max post size of messages defined for a queue, which will effect
	// for any messages posted to the queue.
	MaxMessagesPostSize int `json:"_max_messages_post_size,omitempty"`

	// Description of the queue.
	Description string `json:"description,omitempty"`
}

// ToQueueCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToQueueCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new queue.
func Create(client *gophercloud.ServiceClient, queueName string, clientID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToQueueCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	// Zaqar uses PUT instead of Create for creating queues.
	_, r.Err = client.Put(updateURL(client, queueName), b, r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{201, 204},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})

	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// update request.
type UpdateOptsBuilder interface {
	ToQueueUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is an array of UpdateQueueBody.
type UpdateOpts []UpdateQueueBody

// UpdateOpts implements UpdateOpts.
type UpdateQueueBody struct {
	Op    string      `json:"op" required:"true"`
	Path  string      `json:"path" required:"true"`
	Value interface{} `json:"value" required:"true"`
}

// ToQueueUpdateMap constructs a request body from UpdateOpts.
func (opts UpdateOpts) ToQueueUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Update Updates the specified queue.
func Update(client *gophercloud.ServiceClient, queueName string, clientID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_, r.Err = client.Patch(updateURL(client, queueName), opts, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 204},
		MoreHeaders: map[string]string{
			"Client-ID":    clientID,
			"Content-Type": "application/openstack-messaging-v2.0-json-patch"},
	})
	return
}

// Delete deletes the specified queue.
func Delete(client *gophercloud.ServiceClient, queueName string, clientID string) (r DeleteResult) {
	_, r.Err = client.Request("DELETE", deleteURL(client, queueName), &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// Get requests details on a single queue, by name.
func Get(client *gophercloud.ServiceClient, queueName string, clientID string) (r GetResult) {
	_, r.Err = client.Get(idURL(client, queueName), &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientID}})
	return
}

// GetStats returns statistics for the specified queue.
func GetStats(client *gophercloud.ServiceClient, queueName string, clientID string) (r StatResult) {
	_, r.Err = client.Get(actionURL(client, queueName, "stats"), &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientID}})
	return
}

// CreateOpts specifies share creation parameters.
type ShareOpts struct {
	Paths   []string `json:"paths,omitempty"`
	Methods []string `json:"methods,omitempty"`
	Expires string   `json:"expires,omitempty"`
}

// ShareOptsBuilder allows extensions to add additional attributes to the
// Share request.
type ShareOptsBuilder interface {
	ToShareQueueMap() (map[string]interface{}, error)
}

// ToShareQueueMap formats a ShareOpts structure into a request body.
func (opts ShareOpts) ToShareQueueMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Share creates a pre-signed URL for a given queue.
func Share(client *gophercloud.ServiceClient, queueName string, clientID string, opts ShareOptsBuilder) (r ShareResult) {
	b, err := opts.ToShareQueueMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Post(actionURL(client, queueName, "share"), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// PurgeOpts specifies the purge parameters.
type PurgeOpts struct {
	ResourceTypes []string `json:"resource_types" required:"true"`
}

// PurgeOptsBuilder allows extensions to add additional attributes to the
// Purge request.
type PurgeOptsBuilder interface {
	ToPurgeQueueMap() (map[string]interface{}, error)
}

// ToPurgeQueueMap formats a PurgeOpts structure into a request body
func (opts PurgeOpts) ToPurgeQueueMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Purge purges particular resource of the queue.
func Purge(client *gophercloud.ServiceClient, queueName string, clientID string, opts PurgeOptsBuilder) (r CreateResult) {
	b, err := opts.ToPurgeQueueMap()
	if err != nil {
		r.Err = err
		return r
	}

	_, r.Err = client.Post(actionURL(client, queueName, "purge"), b, nil, &gophercloud.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}
