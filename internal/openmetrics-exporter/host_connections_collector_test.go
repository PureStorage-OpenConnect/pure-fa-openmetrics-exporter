package collectors


import (
	"fmt"
	"testing"
        "regexp"
        "strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"os"

	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

func TestHostConnectionsCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/connections.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var conn client.ConnectionsList
	json.Unmarshal(res, &conn)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/connections$`)
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
	want := make(map[string]bool)
	for _, hc := range conn.Items {
		want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"hostgroup\" value:\"%s\"} label:{name:\"volume\" value:\"%s\"} gauge:{value:1}", hc.Host.Name, hc.HostGroup.Name, hc.Volume.Name)] = true
	}
        c := client.NewRestClient(e, "fake-api-token", "latest", false)
	hc := NewHostConnectionsCollector(c)
        metricsCheck(t, hc, want)
        server.Close()
}
