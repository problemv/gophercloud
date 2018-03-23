package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/queues"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// QueueName is the name of the queue
var QueueName = "FakeTestQueue"

// ClientID is a required parameter used the the Header.
var ClientID = "1234567890"

// ListQueuesResponse is a sample response to a List queues.
const ListQueuesResponse = `
{
   "queues":[
      {
         "href":"/v2/queues/beijing",
         "name":"beijing"
      },
      {
         "href":"/v2/queues/london",
         "name":"london"
      }
   ]
}`

// GetQueueResponse is a sample response to a get queue.
const GetQueueResponse = `
{
    "_max_messages_post_size": 262144,
    "_default_message_ttl": 3600,
    "description": "Queue used for unit testing."
}`

// GetStatsResponse is a sample response to a stats request.
const GetStatsResponse = `
{
    "messages":{
         "claimed": 10,
         "total": 20,
         "free": 10
    }
}`

// CreateQueueRequest is a sample request to create a queue.
const CreateQueueRequest = `
{
    "_max_messages_post_size": 262144,
    "_default_message_ttl": 3600,
    "_default_message_delay": 30,
    "_dead_letter_queue": "dead_letter",
    "_dead_letter_queue_messages_ttl": 3600,
    "_max_claim_count": 10,
    "description": "Queue for unit testing."
}`

// UpdateQueueRequest is a sample request to update a queue.
const UpdateQueueRequest = `
[
    {
        "op": "replace",
        "path": "/metadata/_max_claim_count",
        "value": 10
    }
]`

// UpdateQueueResponse is a sample response to a update queue.
const UpdateQueueResponse = `
{
    "_max_claim_count": 10
}`

// CreateShareRequest is a sample request to a share.
const CreateShareRequest = `
{
  "paths": ["messages", "claims", "subscriptions"],
  "methods": ["GET", "POST", "PUT", "PATCH"],
  "expires": "2016-09-01T00:00:00"
}`

// CreateShareResponse is a sample response to a share request.
const CreateShareResponse = `
{
    "project": "2887aabf368046a3bb0070f1c0413470",
    "paths": [
        "/v2/queues/test/messages",
        "/v2/queues/test/claims",
        "/v2/queues/test/subscriptions"
    ],
    "expires": "2016-09-01T00:00:00",
    "methods": [
        "GET",
        "PATCH",
        "POST",
        "PUT"
    ],
    "signature": "6a63d63242ebd18c3518871dda6fdcb6273db2672c599bf985469241e9a1c799"
}`

// CreatePurgeRequest is a sample request to a purge.
const CreatePurgeRequest = `
{
    "resource_types": ["messages", "subscriptions"]
}`

// FirstQueue is the first result in a List.
var FirstQueue = queues.Queue{
	Href: "/v2/queues/beijing",
	Name: "beijing",
}

// SecondQueue is the second result in a List.
var SecondQueue = queues.Queue{
	Href: "/v2/queues/london",
	Name: "london",
}

// QueueDetails is the expected result in a Get.
var QueueDetails = queues.QueueDetails{
	DefaultMessageTTL:   3600,
	MaxMessagesPostSize: 262144,
	Description:         "Queue used for unit testing.",
}

// ExpectedQueueSlice is the expected result in a List.
var ExpectedQueueSlice = []queues.Queue{FirstQueue, SecondQueue}

// ExpectedStats is the expected result in a GetStats.
var ExpectedStats = queues.Stats{
	Claimed: 10,
	Total:   20,
	Free:    10,
}

// ExpectedShare is the expected result in Share.
var ExpectedShare = queues.QueueShare{
	Project: "2887aabf368046a3bb0070f1c0413470",
	Paths: []string{
		"/v2/queues/test/messages",
		"/v2/queues/test/claims",
		"/v2/queues/test/subscriptions",
	},
	Expires: "2016-09-01T00:00:00",
	Methods: []string{
		"GET",
		"PATCH",
		"POST",
		"PUT",
	},
	Signature: "6a63d63242ebd18c3518871dda6fdcb6273db2672c599bf985469241e9a1c799",
}

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2/queues",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, ListQueuesResponse)
		})
}

// HandleGetSuccessfully configures the test server to respond to a Get request.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, GetQueueResponse)
		})
}

// HandleGetSuccessfully configures the test server to respond to a Get request.
func HandleGetStatsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/stats", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, GetStatsResponse)
		})
}

// HandleCreateSuccessfully configures the test server to respond to a Create request.
func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, CreateQueueRequest)

			w.WriteHeader(http.StatusNoContent)
		})
}

// HandleUpdateSuccessfully configures the test server to respond to an Update request.
func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, UpdateQueueRequest)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, UpdateQueueResponse)
		})
}

// HandleDeleteSuccessfully configures the test server to respond to a Delete request.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.WriteHeader(http.StatusNoContent)
		})
}

// HandleShareSuccessfully configures the test server to respond to a Share request.
func HandleShareSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/share", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, CreateShareRequest)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, CreateShareResponse)
		})
}

// HandlePurgeSuccessfully configures the test server to respond to a Purge request.
func HandlePurgeSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/purge", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, CreatePurgeRequest)

			w.WriteHeader(http.StatusNoContent)
		})
}
