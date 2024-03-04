package collectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
)

func TestNetworkInterfacesCollectorTest(t *testing.T) {

	rhw, _ := os.ReadFile("../../test/data/network_interfaces.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var hwl client.NetworkInterfacesList
	json.Unmarshal(rhw, &hwl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/network-interfaces$`)
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
		want[fmt.Sprintf("label:{name:\"enabled\"  value:\"%s\"}  label:{name:\"ethsubtype\"  value:\"%s\"}  label:{name:\"name\"  value:\"%s\"}  label:{name:\"services\"  value:\"%s\"}  label:{name:\"type\"  value:\"%s\"}  gauge:{value:%g}", strconv.FormatBool(h.Enabled), h.Eth.Subtype, h.Name, strings.Join(h.Services, ", "), h.InterfaceType, float64(h.Speed))] = true
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", false)
	nc := NewNetworkInterfacesCollector(c)
	metricsCheck(t, nc, want)
	server.Close()
}
