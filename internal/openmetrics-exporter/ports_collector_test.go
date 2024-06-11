package collectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
)

func TestPortsCollector(t *testing.T) {

	rhw, _ := os.ReadFile("../../test/data/ports.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var hwl client.PortsList
	json.Unmarshal(rhw, &hwl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/ports$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if url.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(rhw))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	want := make(map[string]bool)
	for _, h := range hwl.Items {
		want[fmt.Sprintf("label:{name:\"iqn\"  value:\"%s\"}  label:{name:\"name\"  value:\"%s\"}  label:{name:\"nqn\"  value:\"%s\"}  label:{name:\"portal\"  value:\"%s\"}  label:{name:\"wwn\"  value:\"%s\"}  gauge:{value:1}", h.Iqn, h.Name, h.Nqn, h.Portal, h.Wwn)] = true
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false)
	pc := NewPortsCollector(c)
	metricsCheck(t, pc, want)
	server.Close()
}
