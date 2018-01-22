package claims

import (
	"github.com/gophercloud/gophercloud"
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

type Message struct {
	Age      int           `json:"age"`
	Href     string        `json:"href"`
	TTL      int           `json:"ttl"`
}

type Claim struct {
	Messages []Message `json:"messages"`
}

func (r commonResult) ExtractClaim() (*Claim, error) {
	claim := Claim{}
	err := r.ExtractInto(&claim)
	return &claim, err
}

func (r commonResult) Extract() (*Claim, error) {
	return r.ExtractClaim()
}
