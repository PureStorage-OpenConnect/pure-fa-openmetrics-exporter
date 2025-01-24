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

func TestVolumesPerformanceCollector(t *testing.T) {
	purenaa := "naa.624a9370"
	volsperf, _ := os.ReadFile("../../test/data/volumes_performance.json")
	vols, _ := os.ReadFile("../../test/data/volumes.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var volumesperf client.VolumesPerformanceList
	var volumes client.VolumesList
	json.Unmarshal(volsperf, &volumesperf)
	json.Unmarshal(vols, &volumes)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vperfu := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/volumes/performance$`)
		vu := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/volumes$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if vperfu.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(volsperf))
		} else if vu.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vols))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	defer server.Close()
	naaid := make(map[string]string)
	for _, v := range volumes.Items {
		naaid[v.Name] = purenaa + v.Serial
	}
	want := make(map[string]bool)
	for _, p := range volumesperf.Items {
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"qos_rate_limit_usec_per_mirrored_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, float64(*p.QosRateLimitUsecPerMirroredWriteOp))] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"qos_rate_limit_usec_per_read_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, float64(*p.QosRateLimitUsecPerReadOp))] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"qos_rate_limit_usec_per_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, float64(*p.QosRateLimitUsecPerWriteOp))] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"queue_usec_per_mirrored_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.QueueUsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"queue_usec_per_read_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.QueueUsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"queue_usec_per_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.QueueUsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"san_usec_per_mirrored_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.SanUsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"san_usec_per_read_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.SanUsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"san_usec_per_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.SanUsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_mirrored_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.ServiceUsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_read_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.ServiceUsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.ServiceUsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_mirrored_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.UsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_read_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.UsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_write_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.UsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_read_op_cache_reduction\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.ServiceUsecPerReadOpCacheReduction)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"mirrored_write_bytes_per_sec\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.MirroredWriteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"read_bytes_per_sec\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.ReadBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"write_bytes_per_sec\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.WriteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"mirrored_writes_per_sec\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.MirroredWritesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"reads_per_sec\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.ReadsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"writes_per_sec\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.WritesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_mirrored_write\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.BytesPerMirroredWrite)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_op\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.BytesPerOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_read\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.BytesPerRead)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_write\"} label:{name:\"naa_id\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", naaid[p.Name], p.Name, p.BytesPerWrite)] = true
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)

	vl := c.GetVolumes()
	pc := NewVolumesPerformanceCollector(c, vl)
	metricsCheck(t, pc, want)
}
