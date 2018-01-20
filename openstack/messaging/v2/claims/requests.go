package claims

import (
	"github.com/gophercloud/gophercloud"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToClaimCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	Limit int `q:"limit,omitempty"`

	TTL int `json:"ttl,omitempty"`

	Grace int `json:"grace,omitempty"`
}

func (opts CreateOpts) ToClaimCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func Create(client *gophercloud.ServiceClient, queueName string, clientId string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToClaimCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, queueName, "claims"), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201, 204},
		MoreHeaders: map[string]string{"Client-ID": clientId,},
	})
	return
}

func Get(client *gophercloud.ServiceClient, queueName string, claimId string, clientId string) (r GetResult) {
	_, r.Err = client.Get(claimURL(client, queueName, claimId), &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId},})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToClaimUpdateMap() (map[string]interface{}, error)
}

// UpdateClaimOpts implements UpdateOpts
type UpdateClaimOpts struct {
	TTL int `json:"ttl,omitempty"`

	Grace int `json:"grace,omitempty"`
}

// ToQueueUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateClaimOpts) ToClaimUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, queueName string, claimId string, clientId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToClaimUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Patch(claimURL(client, queueName, claimId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{204},
		MoreHeaders: map[string]string{"Client-ID": clientId,},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, queueName string, clientId string, claimId string,) (r DeleteResult) {
	_, r.Err = client.Delete(claimURL(client, queueName, claimId), &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId,},
	})
	return
}
