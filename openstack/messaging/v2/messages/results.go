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

// GetMessagesResult is the response of a GetMessages operations.
type GetMessagesResult struct {
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

// Message represents a message on a queue.
type Message struct {
	Body     map[string]interface{} `json:"body"`
	Age      int                    `json:"age"`
	Href     string                 `json:"href"`
	ID       string                 `json:"id"`
	TTL      int                    `json:"ttl"`
	Delay    int                    `json:"delay"`
	Checksum string                 `json:"checksum"`
}

// ResourceList represents the result of creating a message.
type ResourceList struct {
	Resources []string `json:"resources"`
}

// ExtractMessage extracts message into a Message struct.
func (r commonResult) ExtractMessage() (*Message, error) {
	var s struct {
		Message *Message `json:"message"`
	}
	err := r.ExtractInto(&s)
	return s.Message, err
}

// Extract interprets any commonResult as a Message.
func (r commonResult) Extract() (*Message, error) {
	return r.ExtractMessage()
}

// Extract interprets any CreateResult as a ResourceList.
func (r CreateResult) Extract() (ResourceList, error) {
	var s ResourceList
	err := r.ExtractInto(&s)
	return s, err
}

// Extract interprets any GetMessagesResult as a list of Message.
func (r GetMessagesResult) Extract() ([]Message, error) {
	var s struct {
		Messages []Message `json:"messages"`
	}
	err := r.ExtractInto(&s)
	return s.Messages, err
}

// Extract interprets any GetResult as a Message.
func (r GetResult) Extract() (Message, error) {
	var s Message
	err := r.ExtractInto(&s)
	return s, err
}

// ExtractMessage extracts message into a  list of Message.
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
