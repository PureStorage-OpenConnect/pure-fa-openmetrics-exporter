package collectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
	"regexp"
	"strings"
	"testing"
)

func TestVolumesSpaceCollector(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/volume_groups.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var volumeGroups client.VolumeGroupsList
	json.Unmarshal(res, &volumeGroups)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/volume-groups$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if valid.MatchString(r.URL.Path) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	want := make(map[string]bool)
	for _, vg := range volumeGroups.Items {
		if vg.QoS.BandwidthLimit != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", vg.Name, float64(*v.QoS.BandwidthLimit))] = true
		}
		if vg.QoS.IopsLimit != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", vg.Name, float64(*v.QoS.IopsLimit))] = true
		}
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false)
	vgc := NewVolumeGroupsCollector(c)
	metricsCheck(t, vgc, want)
}
