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

func TestArrayControllersCollector(t *testing.T) {

	rdr, _ := os.ReadFile("../../test/data/controllers.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var contl client.ControllersList
	json.Unmarshal(rdr, &contl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/controllers$`)
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
	for _, ctl := range contl.Items {
		want[fmt.Sprintf("label:{name:\"mode\" value:\"%s\"} label:{name:\"model\"value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"status\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} label:{name:\"version\" value:\"%s\"} gauge:{value:\"%g\"}", ctl.Mode, ctl.Model, ctl.Name, ctl.Status, ctl.Type, ctl.Version, (float64(ctl.ModeSince)/1000))] = true
	}

	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	dc := NewControllersCollector(c)
	metricsCheck(t, dc, want)
	server.Close()
}
