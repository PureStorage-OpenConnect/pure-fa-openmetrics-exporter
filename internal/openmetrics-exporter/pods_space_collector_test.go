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

func TestPodsSpaceCollector(t *testing.T) {

	res, _ := ioutil.ReadFile("../../test/data/pods.json")
	vers, _ := ioutil.ReadFile("../../test/data/versions.json")
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
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > gauge:<value:%g > ", p.Name, p.Space.DataReduction)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"shared\" > gauge:<value:%g > ", p.Name, p.Space.Shared)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"snapshots\" > gauge:<value:%g > ", p.Name, p.Space.Snapshots)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"system\" > gauge:<value:%g > ", p.Name, p.Space.System)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"thin_provisioning\" > gauge:<value:%g > ", p.Name, p.Space.ThinProvisioning)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_physical\" > gauge:<value:%g > ", p.Name, p.Space.TotalPhysical)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_provisioned\" > gauge:<value:%g > ", p.Name, p.Space.TotalProvisioned)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_reduction\" > gauge:<value:%g > ", p.Name, p.Space.TotalReduction)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"unique\" > gauge:<value:%g > ", p.Name, p.Space.Unique)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"virtual\" > gauge:<value:%g > ", p.Name, p.Space.Virtual)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"replication\" > gauge:<value:%g > ", p.Name, p.Space.Replication)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"shared_effective\" > gauge:<value:%g > ", p.Name, p.Space.SharedEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"snapshots_effective\" > gauge:<value:%g > ", p.Name, p.Space.SnapshotsEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"unique_effective\" > gauge:<value:%g > ", p.Name, p.Space.UniqueEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_effective\" > gauge:<value:%g > ", p.Name, p.Space.TotalEffective)] = true
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest")
	pc := NewPodsSpaceCollector(c)
	metricsCheck(t, pc, want)
}
