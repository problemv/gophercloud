package subscriptions

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
type SubscriptionPage struct {
	pagination.LinkedPageBase
}

type Subscription struct {
	Subscriber 		string 					`json:"subscriber"`
	TTL 			int						`json:"ttl"`
	Options 		map[string]interface{} 	`json:"options"`
	SubscriptionID 	string					`json:"subscription_id"`
	Age				int						`json:"age"`
	ID				string					`json:"id"`
	Source			string					`json:"source"`
}

func (r commonResult) ExtractSubscription() (*Subscription, error) {
	var s struct {
		Subscription *Subscription `json:"subscription"`
	}
	err := r.ExtractInto(&s)
	return s.Subscription, err
}

func (r commonResult) Extract() (*Subscription, error) {
	return r.ExtractSubscription()
}

func ExtractSubscriptions(r pagination.Page) ([]Subscription, error) {
	var s struct {
		Subscriptions []Subscription `json:"subscriptions"`
	}
	err := (r.(SubscriptionPage)).ExtractInto(&s)
	return s.Subscriptions, err
}

func (page SubscriptionPage) IsEmpty() (bool, error) {
	clusters, err := ExtractSubscriptions(page)
	return len(clusters) == 0, err
}

type SubscriptionResult struct {
	Subscription string `json:"subscriptions"`
}

// Create functions.
func (r commonResult) ExtractSubscriptions() (s *SubscriptionResult, err error) {
	err = r.ExtractInto(&s)
	return s, err
}
