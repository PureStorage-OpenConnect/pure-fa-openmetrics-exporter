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

func TestHardwareCollector(t *testing.T) {

	rhw, _ := os.ReadFile("../../test/data/hardware.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var hwl client.HardwareList
	json.Unmarshal(rhw, &hwl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        url := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/hardware$`)
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
	var s float64
	for _, h := range hwl.Items {
		if h.Status != "ok" {
                        s = 1
                } else {
                        s = 0
                }
		want[fmt.Sprintf("label:<name:\"component_name\" value:\"%s\" > label:<name:\"component_type\" value:\"%s\" > gauge:<value:%g > ", h.Name, h.Type, s)] = true
		if h.Temperature > 0 {
			want[fmt.Sprintf("label:<name:\"component_name\" value:\"%s\" > label:<name:\"component_type\" value:\"%s\" > gauge:<value:%g > ", h.Name, h.Type, float64(h.Temperature))] = true
		}
		if h.Voltage > 0 {
			want[fmt.Sprintf("label:<name:\"component_name\" value:\"%s\" > label:<name:\"component_type\" value:\"%s\" > gauge:<value:%g > ", h.Name, h.Type, float64(h.Voltage))] = true
		}
	}
        c := client.NewRestClient(e, "fake-api-token", "latest", false)
	hc := NewHardwareCollector(c)
        metricsCheck(t, hc, want)
        server.Close()
}
