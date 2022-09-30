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

func TestVolumesSpaceCollector(t *testing.T) {

	res, _ := ioutil.ReadFile("../../test/data/volumes.json")
	vers, _ := ioutil.ReadFile("../../test/data/versions.json")
	var volumes client.VolumesList
	json.Unmarshal(res, &volumes)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/volumes$`)
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
	for _, v := range volumes.Items {
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > gauge:<value:%g > ", v.Name, v.Space.DataReduction)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"shared\" > gauge:<value:%g > ", v.Name, v.Space.Shared)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"snapshots\" > gauge:<value:%g > ", v.Name, v.Space.Snapshots)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"system\" > gauge:<value:%g > ", v.Name, v.Space.System)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"thin_provisioning\" > gauge:<value:%g > ", v.Name, v.Space.ThinProvisioning)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_physical\" > gauge:<value:%g > ", v.Name, v.Space.TotalPhysical)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_provisioned\" > gauge:<value:%g > ", v.Name, v.Space.TotalProvisioned)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_reduction\" > gauge:<value:%g > ", v.Name, v.Space.TotalReduction)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"unique\" > gauge:<value:%g > ", v.Name, v.Space.Unique)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"virtual\" > gauge:<value:%g > ", v.Name, v.Space.Virtual)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"replication\" > gauge:<value:%g > ", v.Name, v.Space.Replication)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"shared_effective\" > gauge:<value:%g > ", v.Name, v.Space.SharedEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"snapshots_effective\" > gauge:<value:%g > ", v.Name, v.Space.SnapshotsEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"unique_effective\" > gauge:<value:%g > ", v.Name, v.Space.UniqueEffective)] = true
		want[fmt.Sprintf("label:<name:\"name\" value:\"%s\" > label:<name:\"space\" value:\"total_effective\" > gauge:<value:%g > ", v.Name, v.Space.TotalEffective)] = true
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest")
	pc := NewVolumesSpaceCollector(c)
	metricsCheck(t, pc, want)
}
