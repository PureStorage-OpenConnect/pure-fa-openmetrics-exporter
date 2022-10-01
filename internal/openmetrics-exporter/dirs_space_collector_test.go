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
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > gauge:<value:%g > ", d.Name, d.Space.DataReduction)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"shared\" > gauge:<value:%g > ", d.Name, d.Space.Shared)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"snapshots\" > gauge:<value:%g > ", d.Name, d.Space.Snapshots)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"system\" > gauge:<value:%g > ", d.Name, d.Space.System)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"thin_provisioning\" > gauge:<value:%g > ", d.Name, d.Space.ThinProvisioning)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_physical\" > gauge:<value:%g > ", d.Name, d.Space.TotalPhysical)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_provisioned\" > gauge:<value:%g > ", d.Name, d.Space.TotalProvisioned)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_reduction\" > gauge:<value:%g > ", d.Name, d.Space.TotalReduction)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"unique\" > gauge:<value:%g > ", d.Name, d.Space.Unique)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"virtual\" > gauge:<value:%g > ", d.Name, d.Space.Virtual)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"replication\" > gauge:<value:%g > ", d.Name, d.Space.Replication)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"shared_effective\" > gauge:<value:%g > ", d.Name, d.Space.SharedEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"snapshots_effective\" > gauge:<value:%g > ", d.Name, d.Space.SnapshotsEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"unique_effective\" > gauge:<value:%g > ", d.Name, d.Space.UniqueEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_effective\" > gauge:<value:%g > ", d.Name, d.Space.TotalEffective)] = true
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", false)
	hc := NewDirectoriesSpaceCollector(c)
	metricsCheck(t, hc, want)
}
