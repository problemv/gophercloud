package pools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// commonResult is the response of a base result.
type commonResult struct {
	gophercloud.Result
}

// CreateResult is the response of a Create operations.
type CreateResult struct {
	commonResult
}

// GetResult is the response of a Get operations.
type GetResult struct {
	commonResult
}

// PostResult is the response of a Post operations.
type PostResult struct {
	commonResult
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult is the response of a Update operations.
type UpdateResult struct {
	commonResult
}

// ClusterPage contains a single page of all clusters from a ListDetails call.
type PoolPage struct {
	pagination.LinkedPageBase
}

type Pool struct {
	Href 	string 					`json:"href"`
	Group 	string 					`json:"group"`
	Name 	string					`json:"name"`
	Weight	int						`json:"weight"`
	URI		string					`json:"uri"`
}

func (r commonResult) ExtractPool() (*Pool, error) {
	var s struct {
		Pool *Pool `json:"pool"`
	}
	err := r.ExtractInto(&s)
	return s.Pool, err
}

func (r commonResult) Extract() (*Pool, error) {
	return r.ExtractPool()
}

func ExtractPools(r pagination.Page) ([]Pool, error) {
	var s struct {
		Pools []Pool `json:"pools"`
	}
	err := (r.(PoolPage)).ExtractInto(&s)
	return s.Pools, err
}

// IsEmpty determines if a ClusterPage contains any results.
func (page PoolPage) IsEmpty() (bool, error) {
	pools, err := ExtractPools(page)
	return len(pools) == 0, err
}

type PoolResult struct {
	Pool string `json:"pools"`
}

// Extract provides access to the individual Pool returned by the Get and
// Create functions.
func (r commonResult) ExtractPools() (s *PoolResult, err error) {
	err = r.ExtractInto(&s)
	return s, err
}
