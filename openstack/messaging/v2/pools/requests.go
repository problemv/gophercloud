package pools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToPoolCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	Weight int `json:"weight" required:"true"`

	URI string `json:"uri" required:"true"`

	Group string `json:"group,omitempty"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// ToClusterCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToPoolCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create uses put instead of post
func Create(client *gophercloud.ServiceClient, poolName string, clientId string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPoolCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(poolURL(client, poolName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
		MoreHeaders: map[string]string{"Client-ID": clientId,},
	})
	return
}

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToPoolListQuery() (string, error)
}

// ListOpts params
type ListOpts struct {
	Detailed bool `q:"detailed,omitempty"`

	Marker string `q:"marker,omitempty"`
}

// ToClusterListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPoolListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List instructs OpenStack to provide a list of pools.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := commonURL(client)
	if opts != nil {
		query, err := opts.ToPoolListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PoolPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type UpdateOpts struct {
	Weight int `json:"weight" required:"true"`
	URI string `json:"uri" required:"true"`
	Group string `json:"group,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPoolUpdateMap() (map[string]interface{}, error)
}

// ToClusterUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToPoolUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Update implements profile updated request.
func Update(client *gophercloud.ServiceClient, poolName string, clientId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPoolUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Patch(poolURL(client, poolName), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{"Client-ID": clientId,},
	})
	return
}

func Get(client *gophercloud.ServiceClient, poolName string, clientId string,) (r GetResult) {
	_, r.Err = client.Get(poolURL(client, poolName), &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId,},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, poolName string, clientId string,) (r DeleteResult) {
	_, r.Err = client.Delete(poolURL(client, poolName), &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Client-ID": clientId,},
	})
	return
}
