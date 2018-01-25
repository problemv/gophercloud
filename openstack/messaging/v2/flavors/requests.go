package flavors

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToFlavorCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	PoolGroup string `json:"pool_group" required:"true"`
}

// ToClusterCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToFlavorCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create uses put instead of post
func Create(client *gophercloud.ServiceClient, clientId string, poolName string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToFlavorCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(flavorURL(client, poolName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{201},
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToFlavorListQuery() (string, error)
}

// ListOpts params
type ListOpts struct {
	Limit int `q:"limit,omitempty"`

	Detailed bool `q:"detailed,omitempty"`

	Marker string `q:"marker,omitempty"`
}

// ToClusterListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List instructs OpenStack to provide a list of pools.
func List(client *gophercloud.ServiceClient, clientId string, opts ListOptsBuilder) pagination.Pager {
	headers := map[string]string{"Client-ID": clientId}
	url := commonURL(client)
	if opts != nil {
		query, err := opts.ToFlavorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pager := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return FlavorPage{pagination.LinkedPageBase{PageResult: r}}
	})
	pager.Headers = headers
	return pager
}

type UpdateOpts struct {
	PoolGroup int `json:"pool_group" required:"true"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToFlavorUpdateMap() (map[string]interface{}, error)
}

// ToClusterUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToFlavorUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Update implements profile updated request.
func Update(client *gophercloud.ServiceClient, flavorName string, opts UpdateOptsBuilder, clientId string) (r UpdateResult) {
	b, err := opts.ToFlavorUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Patch(flavorURL(client, flavorName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}

func Get(client *gophercloud.ServiceClient, flavorName string, clientId string) (r GetResult) {
	_, r.Err = client.Get(flavorURL(client, flavorName), &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, flavorName string, clientId string) (r DeleteResult) {
	_, r.Err = client.Delete(flavorURL(client, flavorName), &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId},
	})
	return
}
