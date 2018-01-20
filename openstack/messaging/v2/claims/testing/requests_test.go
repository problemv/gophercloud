package testing

import (
	"testing"

	"fmt"
	"github.com/gophercloud/gophercloud/openstack/messaging/v2/claims"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	var MockResp = `
	{
  		"age": 57,
  		"href": "/v2/queues/demoqueue/claims/51db7067821e727dc24df754",
  		"messages": [
    			{
      				"body": {
        			"event": "BackupStarted"
      			},
      			"href": "/v2/queues/demoqueue/messages/51db6f78c508f17ddc924357?claim_id=51db7067821e727dc24df754"
    		}
  		],
  		"ttl": 300
	}`

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, MockResp)
	})

	expected := &claims.Claim{
		Age:  57,
		Href: "/v2/queues/demoqueue/claims/51db7067821e727dc24df754",
		Messages: []interface{}{
			map[string]interface{}{
				"href": "/v2/queues/demoqueue/messages/51db6f78c508f17ddc924357?claim_id=51db7067821e727dc24df754",
				"body": map[string]interface{}{
					"event": "BackupStarted"}}},
		TTL: 300,
	}

	claim, err := claims.Get(fake.ServiceClient(), "fake_queue", "1234").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expected, claim)
}
