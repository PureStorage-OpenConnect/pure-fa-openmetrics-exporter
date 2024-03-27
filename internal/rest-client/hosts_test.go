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

func TestHosts(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/hosts.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var hosts HostsList
	json.Unmarshal(res, &hosts)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/hosts$`)
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
	t.Run("hosts_1", func(t *testing.T) {
		defer server.Close()
		c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
		hl := c.GetHosts()
		if diff := cmp.Diff(hl.Items, hosts.Items); diff != "" {
			t.Errorf("Mismatch (-want +got):\n%s", diff)
		}
	})
}
