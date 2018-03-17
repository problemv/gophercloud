package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/messaging/v2/claims"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

var QueueName = "FakeTestQueue"

var ClaimID = "51db7067821e727dc24df754"

var ClientID = "1234567890"

// GetClaimResponse is a sample response to a get claim
const GetClaimResponse = `
{
	"age": 50,
	"href": "/v2/queues/demoqueue/claims/51db7067821e727dc24df754",
	"messages": [
		{
			"body": "BackupStarted",
			"href": "/v2/queues/FakeTestQueue/messages/51db6f78c508f17ddc924357?claim_id=51db7067821e727dc24df754",
			"age": 57,
			"ttl": 300
		}
	],
	"ttl": 50
}`

// CreateClaimResponse is a sample response to a create claim
const CreateClaimResponse = `
{
	"messages": [
		{
			"body": "BackupStarted",
			"href": "/v2/queues/FakeTestQueue/messages/51db6f78c508f17ddc924357?claim_id=51db7067821e727dc24df754",
			"age": 57,
			"ttl": 300
		}
	]
}`

// CreateClaimRequest is a sample request to create a claim.
const CreateClaimRequest = `
{
	"ttl": 3600,
	"grace": 3600
}
`

// UpdateClaimRequest is a sample request to update a claim.
const UpdateClaimRequest = `
{
	"ttl": 1200,
	"grace": 1600
}`

var FirstClaim = claims.Claim{
	Age:  50,
	Href: "/v2/queues/demoqueue/claims/51db7067821e727dc24df754",
	Messages: []claims.Messages{
		{
			Age:  57,
			Href: fmt.Sprintf("/v2/queues/%s/messages/51db6f78c508f17ddc924357?claim_id=%s", QueueName, ClaimID),
			TTL:  300,
			Body: "BackupStarted",
		},
	},
	TTL: 50,
}

var CreatedClaim = claims.Claim{
	Messages: []claims.Messages{
		{
			Age:  57,
			Href: fmt.Sprintf("/v2/queues/%s/messages/51db6f78c508f17ddc924357?claim_id=%s", QueueName, ClaimID),
			TTL:  300,
			Body: "BackupStarted",
		},
	},
}

// HandleGetSuccessfully configures the test server to respond to a Get request.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/claims/%s", QueueName, ClaimID),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, GetClaimResponse)
		})
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/claims", QueueName),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, CreateClaimRequest)

			w.WriteHeader(http.StatusCreated)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, CreateClaimResponse)
		})
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/claims/%s", QueueName, ClaimID),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, UpdateClaimRequest)

			w.WriteHeader(http.StatusNoContent)
		})
}

// HandleDeleteSuccessfully configures the test server to respond to an Delete request.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/v2/queues/%s/claims/%s", QueueName, ClaimID),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.WriteHeader(http.StatusNoContent)
		})
}
