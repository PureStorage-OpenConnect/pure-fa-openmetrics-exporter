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

func TestArraySpaceCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var arrs client.ArraysList
	json.Unmarshal(res, &arrs)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays$`)
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
	a := arrs.Items[0]
	sp := arrs.Items[0].Space
	if sp.DataReduction != nil {
		want[fmt.Sprintf("gauge:{value:%g}", *sp.DataReduction)] = true
	}
	want[fmt.Sprintf("label:{name:\"space\" value:\"capacity\"} gauge:{value:%g}", a.Capacity)] = true
	if sp.Shared != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"shared\"} gauge:{value:%g}", float64(*sp.Shared))] = true
	}
	if sp.Snapshots != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"snapshots\"} gauge:{value:%g}", float64(*sp.Snapshots))] = true
	}
	if sp.System != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"system\"} gauge:{value:%g}", float64(*sp.System))] = true
	}
	if sp.ThinProvisioning != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"thin_provisioning\"} gauge:{value:%g}", *sp.ThinProvisioning)] = true
	}
	if sp.TotalPhysical != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"total_physical\"} gauge:{value:%g}", float64(*sp.TotalPhysical))] = true
	}
	if sp.TotalProvisioned != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"total_provisioned\"} gauge:{value:%g}", float64(*sp.TotalProvisioned))] = true
	}
	if sp.TotalReduction != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"total_reduction\"} gauge:{value:%g}", *sp.TotalReduction)] = true
	}
	if sp.Unique != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"unique\"} gauge:{value:%g}", float64(*sp.Unique))] = true
	}
	if sp.Virtual != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"virtual\"} gauge:{value:%g}", float64(*sp.Virtual))] = true
	}
	if sp.Replication != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"replication\"} gauge:{value:%g}", float64(*sp.Replication))] = true
	}
	if sp.SharedEffective != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"shared_effective\"} gauge:{value:%g}", float64(*sp.SharedEffective))] = true
	}
	if sp.SnapshotsEffective != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"snapshots_effective\"} gauge:{value:%g}", float64(*sp.SnapshotsEffective))] = true
	}
	if sp.UniqueEffective != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"unique_effective\"} gauge:{value:%g}", float64(*sp.UniqueEffective))] = true
	}
	if sp.TotalEffective != nil {
		want[fmt.Sprintf("label:{name:\"space\" value:\"total_effective\"} gauge:{value:%g}", float64(*sp.TotalEffective))] = true
	}
	want[fmt.Sprintf("label:{name:\"space\" value:\"empty\"} gauge:{value:%g}", a.Capacity-(float64(*a.Space.System)+float64(*a.Space.Replication)+float64(*a.Space.Shared)+float64(*a.Space.Snapshots)+float64(*a.Space.Unique)))] = true
	want[fmt.Sprintf("gauge:{value:%g}", (float64(*a.Space.System)+float64(*a.Space.Replication)+float64(*a.Space.Shared)+float64(*a.Space.Snapshots)+float64(*a.Space.Unique))/a.Capacity*100)] = true
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)

	ac := NewArraySpaceCollector(c)
	metricsCheck(t, ac, want)
}
