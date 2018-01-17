package flavors

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
type FlavorPage struct {
	pagination.LinkedPageBase
}

type Flavor struct {
	Href 			string 		`json:"href"`
	PoolGroup 		string 		`json:"pool_group"`
	Name 			string		`json:"name"`
	Pool			int			`json:"pool"`
	Capabilities	[]string	`json:"capabilities"`
}

func (r commonResult) ExtractFlavor() (*Flavor, error) {
	var s struct {
		Flavor *Flavor `json:"pool"`
	}
	err := r.ExtractInto(&s)
	return s.Flavor, err
}

func (r commonResult) Extract() (*Flavor, error) {
	return r.ExtractFlavor()
}

func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	var s struct {
		Flavors []Flavor `json:"flavors"`
	}
	err := (r.(FlavorPage)).ExtractInto(&s)
	return s.Flavors, err
}

// IsEmpty determines if a ClusterPage contains any results.
func (page FlavorPage) IsEmpty() (bool, error) {
	clusters, err := ExtractFlavors(page)
	return len(clusters) == 0, err
}

type FlavorResult struct {
	Flavor string `json:"flavors"`
}

// Extract provides access to the individual Pool returned by the Get and
// Create functions.
func (r commonResult) ExtractFlavors() (s *FlavorResult, err error) {
	err = r.ExtractInto(&s)
	return s, err
}

