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

func TestArrayPerformanceCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var arrs client.ArraysPerformanceList
	json.Unmarshal(res, &arrs)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays/performance$`)
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
	defer server.Close()
	p := arrs.Items[0]
	want := make(map[string]bool)
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"queue_usec_per_mirrored_write_op\" > gauge:<value:%g > ", p.QueueUsecPerMirroredWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"queue_usec_per_read_op\" > gauge:<value:%g > ", p.QueueUsecPerReadOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"queue_usec_per_write_op\" > gauge:<value:%g > ", p.QueueUsecPerWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"san_usec_per_mirrored_write_op\" > gauge:<value:%g > ", p.SanUsecPerMirroredWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"san_usec_per_read_op\" > gauge:<value:%g > ", p.SanUsecPerReadOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"san_usec_per_write_op\" > gauge:<value:%g > ", p.SanUsecPerWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"service_usec_per_mirrored_write_op\" > gauge:<value:%g > ", p.ServiceUsecPerMirroredWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"service_usec_per_read_op\" > gauge:<value:%g > ", p.ServiceUsecPerReadOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"service_usec_per_write_op\" > gauge:<value:%g > ", p.ServiceUsecPerWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"usec_per_mirrored_write_op\" > gauge:<value:%g > ", p.UsecPerMirroredWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"usec_per_read_op\" > gauge:<value:%g > ", p.UsecPerReadOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"usec_per_write_op\" > gauge:<value:%g > ", p.UsecPerWriteOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"service_usec_per_read_op_cache_reduction\" > gauge:<value:%g > ", p.ServiceUsecPerReadOpCacheReduction)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"local_queue_usec_per_op\" > gauge:<value:%g > ", p.LocalQueueUsecPerOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"usec_per_other_op\" > gauge:<value:%g > ", p.UsecPerOtherOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"mirrored_write_bytes_per_sec\" > gauge:<value:%g > ", p.MirroredWriteBytesPerSec)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"read_bytes_per_sec\" > gauge:<value:%g > ", p.ReadBytesPerSec)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"write_bytes_per_sec\" > gauge:<value:%g > ", p.WriteBytesPerSec)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"mirrored_writes_per_sec\" > gauge:<value:%g > ", p.MirroredWritesPerSec)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"reads_per_sec\" > gauge:<value:%g > ", p.ReadsPerSec)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"writes_per_sec\" > gauge:<value:%g > ", p.WritesPerSec)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"others_per_sec\" > gauge:<value:%g > ", p.OthersPerSec)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"bytes_per_mirrored_write\" > gauge:<value:%g > ", p.BytesPerMirroredWrite)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"bytes_per_op\" > gauge:<value:%g > ", p.BytesPerOp)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"bytes_per_read\" > gauge:<value:%g > ", p.BytesPerRead)] = true
	want[fmt.Sprintf("label:<name:\"dimension\" value:\"bytes_per_write\" > gauge:<value:%g > ", p.BytesPerWrite)] = true
	want[fmt.Sprintf("gauge:<value:%g > ", p.QueueDepth)] = true
	c := client.NewRestClient(e, "fake-api-token", "latest", false)
	pc := NewArraysPerformanceCollector(c)
	metricsCheck(t, pc, want)
}
