package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNetworkInterfacesPerformance(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/network_interfaces_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var nw NetworkInterfacesPerformanceList
	json.Unmarshal(res, &nw)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/network-interfaces/performance$`)
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
	t.Run("network_interfaces_performance_1", func(t *testing.T) {
		defer server.Close()
		c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
		nl := c.GetNetworkInterfacesPerformance()
		if diff := cmp.Diff(nl.Items, nw.Items); diff != "" {
			t.Errorf("Mismatch (-want +got):\n%s", diff)
		}
	})
}
