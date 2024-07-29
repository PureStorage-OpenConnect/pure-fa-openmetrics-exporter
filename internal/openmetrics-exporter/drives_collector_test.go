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

func TestDriveCollector(t *testing.T) {

	rdr, _ := os.ReadFile("../../test/data/drives.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var drl client.DriveList
	json.Unmarshal(rdr, &drl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/drive$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if url.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(rdr))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	want := make(map[string]bool)
	for _, d := range drl.Items {
		want[fmt.Sprintf("label:{name:\"component_name\" value:\"%s\"} label:{name:\"component_status\" value:\"%s\"} label:{name:\"component_type\" value:\"%s\"} label:{name:\"component_type\" value:\"%s\"} gauge:{value:\"%g\"}", d.Name, d.Type, d.Status, d.Protocol, d.Capacity)] = true
	}

	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)

	dc := NewDriveCollector(c)
	metricsCheck(t, dc, want)
	server.Close()
}
