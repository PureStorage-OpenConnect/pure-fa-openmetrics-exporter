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
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > gauge:<value:%g > ", h.Name, h.Space.DataReduction)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"shared\" > gauge:<value:%g > ", h.Name, h.Space.Shared)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"snapshots\" > gauge:<value:%g > ", h.Name, h.Space.Snapshots)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"system\" > gauge:<value:%g > ", h.Name, h.Space.System)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"thin_provisioning\" > gauge:<value:%g > ", h.Name, h.Space.ThinProvisioning)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"total_physical\" > gauge:<value:%g > ", h.Name, h.Space.TotalPhysical)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"total_provisioned\" > gauge:<value:%g > ", h.Name, h.Space.TotalProvisioned)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"total_reduction\" > gauge:<value:%g > ", h.Name, h.Space.TotalReduction)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"unique\" > gauge:<value:%g > ", h.Name, h.Space.Unique)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"virtual\" > gauge:<value:%g > ", h.Name, h.Space.Virtual)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"replication\" > gauge:<value:%g > ", h.Name, h.Space.Replication)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"shared_effective\" > gauge:<value:%g > ", h.Name, h.Space.SharedEffective)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"snapshots_effective\" > gauge:<value:%g > ", h.Name, h.Space.SnapshotsEffective)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"unique_effective\" > gauge:<value:%g > ", h.Name, h.Space.UniqueEffective)] = true
		want[fmt.Sprintf("label:<name:\"host\" value:\"%s\" > label:<name:\"space\" value:\"total_effective\" > gauge:<value:%g > ", h.Name, h.Space.TotalEffective)] = true
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", false)
	hc := NewHostsSpaceCollector(c)
	metricsCheck(t, hc, want)
}
