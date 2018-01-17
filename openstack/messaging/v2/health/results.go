package health

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

type Health struct {
	CatalogReachable 	bool 					`json:"catalog_reachable"`
	StorageReachable	bool 					`json:"storage_reachable"`
	OperationStatus 	map[string]interface{} 	`json:"operation_status"`
}

func (r commonResult) ExtractHealth() (*Health, error) {
	var s struct {
		Health *Health `json:"health"`
	}
	err := r.ExtractInto(&s)
	return s.Health, err
}

func (r commonResult) Extract() (*Health, error) {
	return r.ExtractHealth()
}
