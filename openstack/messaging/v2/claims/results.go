package claims

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
type ClaimPage struct {
	pagination.LinkedPageBase
}

type Claim struct {
	Body 		map[string]interface{}	`json:"body"`
	Age			int						`json:"age"`
	Href		string					`json:"href"`
	TTL			int						`json:"ttl"`
	Messages 	map[string]interface{}	`json:"messages"`
	Grace		int						`json:"grace"`
}

func (r commonResult) ExtractClaim() (*Claim, error) {
	var s struct {
		Claim *Claim `json:"claim"`
	}
	err := r.ExtractInto(&s)
	return s.Claim, err
}

func (r commonResult) Extract() (*Claim, error) {
	return r.ExtractClaim()
}

// ExtractCluster provides access to the list of clusters in a page acquired from the ListDetail operation.
func ExtractClaims(r pagination.Page) ([]Cluster, error) {
	var s struct {
		Claims []Claim `json:"claims"`
	}
	err := (r.(ClaimPage)).ExtractInto(&s)
	return s.Claims, err
}

// IsEmpty determines if a ClusterPage contains any results.
func (page ClaimPage) IsEmpty() (bool, error) {
	clusters, err := ExtractClaims(page)
	return len(clusters) == 0, err
}

type ClusterResult struct {
	Claim string `json:"claims"`
}

// Extract provides access to the individual Cluster returned by the Get and
// Create functions.
func (r commonResult) ExtractClaims() (s *ClusterResult, err error) {
	err = r.ExtractInto(&s)
	return s, err
}
