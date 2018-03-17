package claims

import (
	"github.com/gophercloud/gophercloud"
	"net/http"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToClaimCreateMap() (map[string]interface{}, error)
}

// CreateQueryOptsBuilder Builder.
type CreateQueryOptsBuilder interface {
	ToClaimCreateQuery() (string, error)
}

// CreateOpts params
type CreateOpts struct {
	// Sets the TTL for the claim. When the claim expires un-deleted messages will be able to be claimed again.
	TTL int `json:"ttl,omitempty"`

	// Sets the Grace period for the claimed messages. The server extends the lifetime of claimed messages
	// to be at least as long as the lifetime of the claim itself, plus the specified grace period.
	Grace int `json:"grace,omitempty"`
}

//CreateQueryOpts params
type CreateQueryOpts struct {
	//Set the limit of messages returned by create
	Limit int `q:"limit,omitempty"`
}

//ToClaimCreateQuery assembles url parmeters for a create request
func (opts CreateQueryOpts) ToClaimCreateQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ToClaimCreateMap assembles a body for a Create request based on
// the contents of a CreateOpts.
func (opts CreateOpts) ToClaimCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create creates a Claim that claims messages on a specified queue.
func Create(client *gophercloud.ServiceClient, queueName string, clientID string, opts CreateOptsBuilder, query CreateQueryOptsBuilder) (r CreateResult) {
	b, err := opts.ToClaimCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := actionURL(client, queueName, "claims")
	if query != nil {
		query, err := query.ToClaimCreateQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	var resp *http.Response
	resp, r.Err = client.Post(url, b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{201, 204},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	// If the Claim has no content return an empty CreateResult
	if resp.StatusCode == 204 {
		r.Body = CreateResult{}
	} else {
		r.Body = resp.Body
	}

	return
}

// Get queries the specified claim for the specified queue.
func Get(client *gophercloud.ServiceClient, queueName string, claimID string, clientID string) (r GetResult) {
	_, r.Err = client.Get(claimURL(client, queueName, claimID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToClaimUpdateMap() (map[string]interface{}, error)
}

// UpdateClaimOpts implements UpdateOpts.
type UpdateClaimOpts struct {
	// Update the TTL for the specified Claim.
	TTL int `json:"ttl,omitempty"`

	// Update the grace period for Messages in a specified Claim.
	Grace int `json:"grace,omitempty"`
}

// ToClaimUpdateMap assembles a request body based on the contents of
// UpdateClaimOpts.
func (opts UpdateClaimOpts) ToClaimUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Update will update the options for a specified claim.
func Update(client *gophercloud.ServiceClient, queueName string, claimID string, clientID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToClaimUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Patch(claimURL(client, queueName, claimID), &b, nil, &gophercloud.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}

// Delete will delete a Claim for a specified Queue.
func Delete(client *gophercloud.ServiceClient, queueName string, claimID string, clientID string) (r DeleteResult) {
	_, r.Err = client.Delete(claimURL(client, queueName, claimID), &gophercloud.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Client-ID": clientID},
	})
	return
}
