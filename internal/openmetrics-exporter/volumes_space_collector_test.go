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

func TestVolumesSpaceCollector(t *testing.T) {
	purenaa := "naa.624a9370"
	res, _ := os.ReadFile("../../test/data/volumes.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
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
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.DataReduction)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"shared\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.Shared)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"snapshots\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.Snapshots)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"system\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.System)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"thin_provisioning\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.ThinProvisioning)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_physical\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.TotalPhysical)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_provisioned\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.TotalProvisioned)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_reduction\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.TotalReduction)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"unique\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.Unique)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"virtual\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.Virtual)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"replication\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.Replication)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"shared_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.SharedEffective)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"snapshots_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.SnapshotsEffective)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"unique_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.UniqueEffective)] = true
		want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, v.Space.TotalEffective)] = true
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	vl := c.GetVolumes()
	pc := NewVolumesSpaceCollector(vl)
	metricsCheck(t, pc, want)
}
