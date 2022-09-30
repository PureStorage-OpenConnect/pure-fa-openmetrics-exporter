package collectors


import (
	"fmt"
	"testing"
	"regexp"
	"strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"io/ioutil"

	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

func TestArraySpaceCollector(t *testing.T) {

	res, _ := ioutil.ReadFile("../../test/data/arrays.json")
	vers, _ := ioutil.ReadFile("../../test/data/versions.json")
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
	sp := arrs.Items[0].Space
	want[fmt.Sprintf("gauge:<value:%g > ", sp.DataReduction)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"shared\" > gauge:<value:%g > ", sp.Shared)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"snapshots\" > gauge:<value:%g > ", sp.Snapshots)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"system\" > gauge:<value:%g > ", sp.System)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"thin_provisioning\" > gauge:<value:%g > ", sp.ThinProvisioning)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"total_physical\" > gauge:<value:%g > ", sp.TotalPhysical)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"total_provisioned\" > gauge:<value:%g > ", sp.TotalProvisioned)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"total_reduction\" > gauge:<value:%g > ", sp.TotalReduction)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"unique\" > gauge:<value:%g > ", sp.Unique)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"virtual\" > gauge:<value:%g > ", sp.Virtual)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"replication\" > gauge:<value:%g > ", sp.Replication)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"shared_effective\" > gauge:<value:%g > ", sp.SharedEffective)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"snapshots_effective\" > gauge:<value:%g > ", sp.SnapshotsEffective)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"unique_effective\" > gauge:<value:%g > ", sp.UniqueEffective)] = true
	want[fmt.Sprintf("label:<name:\"space\" value:\"total_effective\" > gauge:<value:%g > ", sp.TotalEffective)] = true
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest")
	ac := NewArraySpaceCollector(c)
	metricsCheck(t, ac, want)
}
