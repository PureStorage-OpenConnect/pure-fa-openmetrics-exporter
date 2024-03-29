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

func TestDirectoriesSpaceCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/directories.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var directories client.DirectoriesList
	json.Unmarshal(res, &directories)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/directories$`)
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
	for _, d := range directories.Items {
		if d.Space.DataReduction != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", d.Name, *d.Space.DataReduction)] = true
		}
		if d.Space.Shared != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"shared\"} gauge:{value:%g}", d.Name, float64(*d.Space.Shared))] = true
		}
		if d.Space.Snapshots != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"snapshots\"} gauge:{value:%g}", d.Name, float64(*d.Space.Snapshots))] = true
		}
		if d.Space.System != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"system\"} gauge:{value:%g}", d.Name, float64(*d.Space.System))] = true
		}
		if d.Space.ThinProvisioning != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"thin_provisioning\"} gauge:{value:%g}", d.Name, *d.Space.ThinProvisioning)] = true
		}
		if d.Space.TotalPhysical != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_physical\"} gauge:{value:%g}", d.Name, float64(*d.Space.TotalPhysical))] = true
		}
		if d.Space.TotalProvisioned != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_provisioned\"} gauge:{value:%g}", d.Name, float64(*d.Space.TotalProvisioned))] = true
		}
		if d.Space.TotalReduction != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_reduction\"} gauge:{value:%g}", d.Name, *d.Space.TotalReduction)] = true
		}
		if d.Space.Unique != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"unique\"} gauge:{value:%g}", d.Name, float64(*d.Space.Unique))] = true
		}
		if d.Space.Virtual != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"virtual\"} gauge:{value:%g}", d.Name, float64(*d.Space.Virtual))] = true
		}
		if d.Space.Replication != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"replication\"} gauge:{value:%g}", d.Name, float64(*d.Space.Replication))] = true
		}
		if d.Space.SharedEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"shared_effective\"} gauge:{value:%g}", d.Name, float64(*d.Space.SharedEffective))] = true
		}
		if d.Space.UniqueEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"snapshots_effective\"} gauge:{value:%g}", d.Name, float64(*d.Space.SnapshotsEffective))] = true
		}
		if d.Space.UniqueEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"unique_effective\"} gauge:{value:%g}", d.Name, float64(*d.Space.UniqueEffective))] = true
		}
		if d.Space.TotalEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_effective\"} gauge:{value:%g}", d.Name, float64(*d.Space.TotalEffective))] = true
		}
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	hc := NewDirectoriesSpaceCollector(c)
	metricsCheck(t, hc, want)
}
