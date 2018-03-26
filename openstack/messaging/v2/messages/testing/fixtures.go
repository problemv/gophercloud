package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/messages"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// QueueName is the name of the queue
var QueueName = "FakeTestQueue"

// ClientID is a required parameter used the the Header.
var ClientID = "1234567890"

// MessageID is the id of the message
var MessageID = "9988776655"

// CreateMessageResponse is a sample response to a Create message.
const CreateMessageResponse = `
{
  "resources": [
    "/v2/queues/demoqueue/messages/51db6f78c508f17ddc924357",
    "/v2/queues/demoqueue/messages/51db6f78c508f17ddc924358"
  ]
}`

// CreateMessageRequest is a sample request to create a message.
const CreateMessageRequest = `
{
  "messages": [
	{
	  "body": {
		"bakcup_id": "c378813c-3f0b-11e2-ad92-7823d2b0f3ce",
		"event": "BackupStarted"
	  },
	  "delay": 20,
	  "ttl": 300
	},
	{
	  "body": {
		"current_bytes": "0",
		"event": "BackupProgress",
		"total_bytes": "99614720"
	  }
	}
  ]
}`

// ListMessagesResponse is a sample response to list messages.
const ListMessagesResponse = `
{
    "messages": [
        {
            "body": {
                "current_bytes": "0",
                "event": "BackupProgress",
                "total_bytes": "99614720"
            },
            "age": 482,
            "href": "/v2/queues/beijing/messages/578edfe6508f153f256f717b",
            "id": "578edfe6508f153f256f717b",
            "ttl": 3600,
            "checksum": "MD5:abf7213555626e29c3cb3e5dc58b3515"
        },
        {
            "body": {
                "current_bytes": "0",
                "event": "BackupProgress",
                "total_bytes": "99614720"
            },
            "age": 456,
            "href": "/v2/queues/beijing/messages/578ee000508f153f256f717d",
            "id": "578ee000508f153f256f717d",
            "ttl": 3600,
            "checksum": "MD5:abf7213555626e29c3cb3e5dc58b3515"
        }
    ]
}`

// GetMessagesResponse is a sample response to GetMessages.
const GetMessagesResponse = `
{
    "messages": [
        {
            "body": {
                "current_bytes": "0",
                "event": "BackupProgress",
                "total_bytes": "99614720"
            },
            "age": 443,
            "href": "/v2/queues/beijing/messages/578f0055508f153f256f717f",
            "id": "578f0055508f153f256f717f",
            "ttl": 3600
        }
    ]
}`

// GetMessageResponse is a sample response to Get.
const GetMessageResponse = `
{
    "body": {
        "current_bytes": "0",
        "event": "BackupProgress",
        "total_bytes": "99614720"
    },
    "age": 482,
    "href": "/v2/queues/beijing/messages/578edfe6508f153f256f717b",
    "id": "578edfe6508f153f256f717b",
    "ttl": 3600,
    "checksum": "MD5:abf7213555626e29c3cb3e5dc58b3515"
}`

// ExpectedResources is the expected result in Create
var ExpectedResources = messages.ResourceList{
	Resources: []string{
		"/v2/queues/demoqueue/messages/51db6f78c508f17ddc924357",
		"/v2/queues/demoqueue/messages/51db6f78c508f17ddc924358",
	},
}

// FirstMessage is the first result in a List.
var FirstMessage = messages.Message{
	Body: map[string]interface{}{
		"current_bytes": "0",
		"event":         "BackupProgress",
		"total_bytes":   "99614720",
	},
	Age:      482,
	Href:     "/v2/queues/beijing/messages/578edfe6508f153f256f717b",
	ID:       "578edfe6508f153f256f717b",
	TTL:      3600,
	Checksum: "MD5:abf7213555626e29c3cb3e5dc58b3515",
}

// SecondMessage is the second result in a List.
var SecondMessage = messages.Message{
	Body: map[string]interface{}{
		"current_bytes": "0",
		"event":         "BackupProgress",
		"total_bytes":   "99614720",
	},
	Age:      456,
	Href:     "/v2/queues/beijing/messages/578ee000508f153f256f717d",
	ID:       "578ee000508f153f256f717d",
	TTL:      3600,
	Checksum: "MD5:abf7213555626e29c3cb3e5dc58b3515",
}

// ExpectedMessagesSet is the expected result in ListMessages
var ExpectedMessagesSet = []messages.Message{
	{
		Body: map[string]interface{}{
			"total_bytes":   "99614720",
			"current_bytes": "0",
			"event":         "BackupProgress",
		},
		Age:      443,
		Href:     "/v2/queues/beijing/messages/578f0055508f153f256f717f",
		ID:       "578f0055508f153f256f717f",
		TTL:      3600,
		Delay:    0,
		Checksum: "",
	},
}

// ExpectedMessagesSlice is the expected result in a List.
var ExpectedMessagesSlice = []messages.Message{FirstMessage, SecondMessage}

// HandleCreateSuccessfully configures the test server to respond to a Create request.
func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/messages", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, CreateMessageRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, CreateMessageResponse)
		})
}

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/messages", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, ListMessagesResponse)
		})
}

// HandleGetMessagesSuccessfully configures the test server to respond to a GetMessages request.
func HandleGetMessagesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/messages", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, GetMessagesResponse)
		})
}

// HandleCreateSuccessfully configures the test server to respond to a Create request.
func HandleDeleteMessagesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/messages", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}

// HandleGetSuccessfully configures the test server to respond to a Get request.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/messages/%s", QueueName, MessageID),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, GetMessageResponse)
		})
}

// HandleGetSuccessfully configures the test server to respond to a Get request.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/messages/%s", QueueName, MessageID),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}
