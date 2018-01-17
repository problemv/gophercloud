package queues

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
type QueuePage struct {
	pagination.LinkedPageBase
}

type Queue struct {
	DeadLetterQueue           string   `json:"_dead_letter_queue"`
	DeadLetterQueueMessageTTL int      `json:"_dead_letter_queue_messages_ttl"`
	DefaultMessageDelay       string   `json:"_default_message_delay"`
	DefaultMessageTTL         int      `json:"_default_message_ttl"`
	Description               string   `json:"description"`
	Expires                   string   `json:"expires"`
	Flavor                    string   `json:"flavor"`
	Href                      string   `json:"href"`
	MaxClaimCount             int      `json:"_max_claim_count"`
	MaxMessagesPostSize       int      `json:"_max_messages_post_size"`
	Methods                   []string `json:"methods"`
	Name                      string   `json:"name"`
	Paths                     []string `json:"paths"`
	ResourceTypes             []string `json:"resource_types"`
}

func (r commonResult) ExtractQueue() (*Queue, error) {
	var s struct {
		Queue *Queue `json:"queue"`
	}
	err := r.ExtractInto(&s)
	return s.Queue, err
}

func (r commonResult) Extract() (*Queue, error) {
	return r.ExtractQueue()
}

func ExtractQueues(r pagination.Page) ([]Queue, error) {
	var s struct {
		Queues []Queue `json:"queues"`
	}
	err := (r.(QueuePage)).ExtractInto(&s)
	return s.Queues, err
}

// IsEmpty determines if a QueuesPage contains any results.
func (page QueuePage) IsEmpty() (bool, error) {
	clusters, err := ExtractQueues(page)
	return len(clusters) == 0, err
}

type QueueResult struct {
	Queue string `json:"queues"`
}

// Extract provides access to the individual Queue returned by the Get and
// Create functions.
func (r commonResult) ExtractQueues() (s *QueueResult, err error) {
	err = r.ExtractInto(&s)
	return s, err
}
