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

func TestPodsSpaceCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/pods.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var pods client.PodsList
	json.Unmarshal(res, &pods)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/pods$`)
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
	for _, p := range pods.Items {
		var s float64
		if p.Space.DataReduction != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, *p.Space.DataReduction)] = true
		}
		if p.Space.Shared != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"shared\"} gauge:{value:%g}", p.Name, float64(*p.Space.Shared))] = true
		}
		if p.Space.Snapshots != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"snapshots\"} gauge:{value:%g}", p.Name, float64(*p.Space.Snapshots))] = true
		}
		if p.Space.System != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"system\"} gauge:{value:%g}", p.Name, float64(*p.Space.System))] = true
		}
		if p.Space.ThinProvisioning != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"thin_provisioning\"} gauge:{value:%g}", p.Name, *p.Space.ThinProvisioning)] = true
		}
		if p.Space.TotalPhysical != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_physical\"} gauge:{value:%g}", p.Name, float64(*p.Space.TotalPhysical))] = true
		}
		if p.Space.TotalProvisioned != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_provisioned\"} gauge:{value:%g}", p.Name, float64(*p.Space.TotalProvisioned))] = true
		}
		if p.Space.TotalReduction != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_reduction\"} gauge:{value:%g}", p.Name, *p.Space.TotalReduction)] = true
		}
		if p.Space.Unique != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"unique\"} gauge:{value:%g}", p.Name, float64(*p.Space.Unique))] = true
		}
		if p.Space.Virtual != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"virtual\"} gauge:{value:%g}", p.Name, float64(*p.Space.Virtual))] = true
		}
		if p.Space.Replication != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"replication\"} gauge:{value:%g}", p.Name, float64(*p.Space.Replication))] = true
		}
		if p.Space.SharedEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"shared_effective\"} gauge:{value:%g}", p.Name, float64(*p.Space.SharedEffective))] = true
		}
		if p.Space.SnapshotsEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"snapshots_effective\"} gauge:{value:%g}", p.Name, float64(*p.Space.SnapshotsEffective))] = true
		}
		if p.Space.UniqueEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"unique_effective\"} gauge:{value:%g}", p.Name, float64(*p.Space.UniqueEffective))] = true
		}
		if p.Space.TotalEffective != nil {
			want[fmt.Sprintf("label:{name:\"name\" value:\"%s\"} label:{name:\"space\" value:\"total_effective\"} gauge:{value:%g}", p.Name, float64(*p.Space.TotalEffective))] = true
		}
		for _, a := range p.Arrays {
			if a.MediatorStatus == "online" {
				s = 1
			} else {
				s = 0
			}
			want[fmt.Sprintf("label:{name:\"array\" value:\"%s\"} label:{name:\"mediator\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"status\" value:\"%s\"} gauge:{value:%g}", a.Name, p.Mediator, p.Name, a.MediatorStatus, s)] = true
		}
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)

	pc := NewPodsSpaceCollector(c)
	metricsCheck(t, pc, want)
}
