package messages

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

// MessagePage contains a single page of all clusters from a ListDetails call.
type MessagePage struct {
	pagination.LinkedPageBase
}

type Message struct {
	TTL   int                    `json:"ttl"`
	Delay int                    `json:"delay"`
	Body  map[string]interface{} `json:"body"`
	Age   int                    `json:"age"`
	Href  string                 `json:"href"`
	ID    string                 `json:"id"`
}

func (r commonResult) ExtractMessage() (*Message, error) {
	var s struct {
		Message *Message `json:"message"`
	}
	err := r.ExtractInto(&s)
	return s.Message, err
}

func (r commonResult) Extract() (*Message, error) {
	return r.ExtractMessage()
}

func ExtractMessages(r pagination.Page) ([]Message, error) {
	var s struct {
		Messages []Message `json:"messages"`
	}
	err := (r.(MessagePage)).ExtractInto(&s)
	return s.Messages, err
}

// IsEmpty determines if a MessagePage contains any results.
func (page MessagePage) IsEmpty() (bool, error) {
	messages, err := ExtractMessages(page)
	return len(messages) == 0, err
}

type MessageResult struct {
	Message string `json:"messages"`
}

// Extract provides access to the individual Message returned by the Get and
// Create functions.
func (r commonResult) ExtractMessages() (s *MessageResult, err error) {
	err = r.ExtractInto(&s)
	return s, err
}
