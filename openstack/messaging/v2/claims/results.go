package claims

import (
	"github.com/gophercloud/gophercloud"
)

// commonResult is the response of a base result.
type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Claim, error) {
	var s *Claim
	err := r.ExtractInto(&s)
	return s, err
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
	gophercloud.ErrResult
}

type Messages struct {
	Age  float32 `json:"age"`
	Href string  `json:"href"`
	TTL  int     `json:"ttl"`
	Body string  `json:"body"`
}

type Claim struct {
	Age      float32   `json:"age"`
	Href     string    `json:"href"`
	Messages []Messages `json:"messages"`
	TTL      int       `json:"ttl"`
}
