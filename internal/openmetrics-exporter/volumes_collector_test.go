package collectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
	"regexp"
	"strings"
	"testing"
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
		if v.QoS.BandwidthLimit != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.QoS.BandwidthLimit))] = true
		}
		if v.QoS.IopsLimit != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.QoS.IopsLimit))] = true
		}
		if v.Space.DataReduction != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, *v.Space.DataReduction)] = true
		}
		if v.Space.Shared != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"shared\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.Shared))] = true
		}
		if v.Space.Snapshots != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"snapshots\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.Snapshots))] = true
		}
		if v.Space.ThinProvisioning != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"system\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.ThinProvisioning))] = true
		}
		if v.Space.ThinProvisioning != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"thin_provisioning\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, *v.Space.ThinProvisioning)] = true
		}
		if v.Space.TotalPhysical != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_physical\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.TotalPhysical))] = true
		}
		if v.Space.TotalProvisioned != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_provisioned\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.TotalProvisioned))] = true
		}
		if v.Space.TotalReduction != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_reduction\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, *v.Space.TotalReduction)] = true
		}
		if v.Space.Unique != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"unique\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.Unique))] = true
		}
		if v.Space.Virtual != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"virtual\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.Virtual))] = true
		}
		if v.Space.Replication != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"replication\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.Replication))] = true
		}
		if v.Space.SharedEffective != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"shared_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.SharedEffective))] = true
		}
		if v.Space.SnapshotsEffective != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"snapshots_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.SnapshotsEffective))] = true
		}
		if v.Space.UniqueEffective != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"unique_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.UniqueEffective))] = true
		}
		if v.Space.TotalEffective != nil {
			want[fmt.Sprintf("label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"pod\" value:\"%s\"} label:{name:\"space\" value:\"total_effective\"} label:{name:\"volume_group\" value:\"%s\"} gauge:{value:%g}", purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, float64(*v.Space.TotalEffective))] = true
		}
	}
	defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)
	vl := c.GetVolumes()
	pc := NewVolumesCollector(vl)
	metricsCheck(t, pc, want)
}
