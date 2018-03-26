package messages

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToMessageCreateMap() (map[string]interface{}, error)
}

// Messages params to be used with CreateOpts.
type Messages struct {
	TTL int `json:"ttl,omitempty"`

	Delay int `json:"delay,omitempty"`

	Body map[string]interface{} `json:"body" required:"true"`
}

// CreateOpts params to be used with Create.
type CreateOpts struct {
	Messages []Messages `json:"messages" required:"true"`
}

// ToMessageCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToMessageCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create creates a message on a specific queue based of off queue name.
func Create(client *gophercloud.ServiceClient, queueName string, clientID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToMessageCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, queueName, "messages"), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{201},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMessageListQuery() (string, error)
}

// ListOpts params to be used with List.
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

// ListMessages lists messages on a specific queue based off queue name.
func List(client *gophercloud.ServiceClient, queueName string, clientID string, opts ListOptsBuilder) pagination.Pager {
	headers := map[string]string{"Client-ID": clientID}
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

// DeleteMessageOptsBuilder allows extensions to add additional parameters to the
// DeleteMessages request.
type DeleteMessageOptsBuilder interface {
	ToMessagesDeleteQuery() (string, error)
}

// DeleteMessageOpts params to be used with DeleteMessages.
type DeleteMessageOpts struct {
	Ids []string `q:"ids,omitempty"`
	Pop int      `q:"pop,omitempty"`
}

// ToMessagesDeleteQuery formats a DeleteMessageOpts structure into a query string.
func (opts DeleteMessageOpts) ToMessagesDeleteQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// DeleteMessages deletes multiple messages based off of ID or number of messages.
func DeleteMessages(client *gophercloud.ServiceClient, queueName string, clientID string, opts DeleteMessageOptsBuilder) (r DeleteResult) {
	url := actionURL(client, queueName, "messages")
	if opts != nil {
		query, err := opts.ToMessagesDeleteQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		OkCodes:     []int{200, 204},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// GetMessagesOptsBuilder allows extensions to add additional parameters to the
// GetMessages request.
type GetMessagesOptsBuilder interface {
	ToGetMessagesListQuery() (string, error)
}

// GetMessagesOpts params to be used with GetMessages.
type GetMessagesOpts struct {
	Ids []string `q:"ids,omitempty"`
}

// ToGetMessagesListQuery formats a GetMessagesOpts structure into a query string.
func (opts GetMessagesOpts) ToGetMessagesListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// GetMessages requests details on a multiple messages, by IDs.
func GetMessages(client *gophercloud.ServiceClient, queueName string, clientID string, opts GetMessagesOptsBuilder) (r GetMessagesResult) {
	url := getURL(client, queueName)
	if opts != nil {
		query, err := opts.ToGetMessagesListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// Get requests details on a single message, by ID.
func Get(client *gophercloud.ServiceClient, queueName string, messageID string, clientID string) (r GetResult) {
	_, r.Err = client.Get(messageURL(client, queueName, messageID), &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// DeleteOptsBuilder allows extensions to add additional parameters to the
// delete request.
type DeleteOptsBuilder interface {
	ToMessageDeleteQuery() (string, error)
}

// DeleteOpts params to be used with Delete.
type DeleteOpts struct {
	// Claim instructs Delete to delete a message that is associated with a claim ID
	Claim string `q:"claim_id,omitempty"`
}

// ToMessageDeleteQuery formats a DeleteOpts structure into a query string.
func (opts DeleteOpts) ToMessageDeleteQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Delete deletes a specific message from the queue.
func Delete(client *gophercloud.ServiceClient, queueName string, messageID string, clientID string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := messageURL(client, queueName, messageID)
	if opts != nil {
		query, err := opts.ToMessageDeleteQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}
