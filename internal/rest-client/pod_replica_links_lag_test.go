package client


import (
	"testing"
        "regexp"
        "strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"os"

	"github.com/google/go-cmp/cmp"
)

func TestPodReplicaLinksLag(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/pod_replica_links_lag.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var podrll PodReplicaLinksLagList
	json.Unmarshal(res, &podrll)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/pod-replica-links/lag$`)
                if r.URL.Path == "/api/api_version" {
                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
                } else if valid.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
		}
	   }))
        endp := strings.Split(server.URL, "/")
        e := endp[len(endp)-1]
        t.Run("pod_replica_links_lag_1", func(t *testing.T) {
            defer server.Close()
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    prll := c.GetPodReplicaLinksLag()
	    if diff := cmp.Diff(prll.Items, podrll.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
            }
        })
}
