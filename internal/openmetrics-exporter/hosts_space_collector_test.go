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

func TestHostsSpaceCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/hosts.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var hosts client.HostsList
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
	want := make(map[string]bool)
	for _, h := range hosts.Items {
		if h.Space.DataReduction != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} gauge:{value:%g}", h.Name, *h.Space.DataReduction)] = true
		}
		if h.Space.Shared != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"shared\"} gauge:{value:%g}", h.Name, float64(*h.Space.Shared))] = true
		}
		if h.Space.Snapshots != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"snapshots\"} gauge:{value:%g}", h.Name, float64(*h.Space.Snapshots))] = true
		}
		if h.Space.System != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"system\"} gauge:{value:%g}", h.Name, float64(*h.Space.System))] = true
		}
		if h.Space.ThinProvisioning != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"thin_provisioning\"} gauge:{value:%g}", h.Name, *h.Space.ThinProvisioning)] = true
		}
		if h.Space.TotalPhysical != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"total_physical\"} gauge:{value:%g}", h.Name, float64(*h.Space.TotalPhysical))] = true
		}
		if h.Space.TotalProvisioned != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"total_provisioned\"} gauge:{value:%g}", h.Name, float64(*h.Space.TotalProvisioned))] = true
		}
		if h.Space.TotalReduction != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"total_reduction\"} gauge:{value:%g}", h.Name, *h.Space.TotalReduction)] = true
		}
		if h.Space.Unique != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"unique\"} gauge:{value:%g}", h.Name, float64(*h.Space.Unique))] = true
		}
		if h.Space.Virtual != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"virtual\"} gauge:{value:%g}", h.Name, float64(*h.Space.Virtual))] = true
		}
		if h.Space.Replication != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"replication\"} gauge:{value:%g}", h.Name, float64(*h.Space.Replication))] = true
		}
		if h.Space.SharedEffective != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"shared_effective\"} gauge:{value:%g}", h.Name, float64(*h.Space.SharedEffective))] = true
		}
		if h.Space.SnapshotsEffective != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"snapshots_effective\"} gauge:{value:%g}", h.Name, float64(*h.Space.SnapshotsEffective))] = true
		}
		if h.Space.UniqueEffective != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"unique_effective\"} gauge:{value:%g}", h.Name, float64(*h.Space.UniqueEffective))] = true
		}
		if h.Space.TotalEffective != nil {
			want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"space\" value:\"total_effective\"} gauge:{value:%g}", h.Name, float64(*h.Space.TotalEffective))] = true
		}
		want[fmt.Sprintf("label:{name:\"host\" value:\"%s\"} label:{name:\"details\" value:\"%s\"} label:{name:\"status\" value:\"%s\"} gauge:{value:1}", h.Name, h.PortConnectivity.Details, h.PortConnectivity.Status)] = true
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)

	hc := NewHostsSpaceCollector(c)
	metricsCheck(t, hc, want)
}
