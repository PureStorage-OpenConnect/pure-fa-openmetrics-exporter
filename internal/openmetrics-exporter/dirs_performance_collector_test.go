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

func TestDirectoriesPerformanceCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/directories_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var dirs client.DirectoriesPerformanceList
	json.Unmarshal(res, &dirs)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/directories/performance$`)
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
	want := make(map[string]bool)
	for _, p := range dirs.Items {
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_other_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.UsecPerOtherOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_read_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.UsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.UsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"read_bytes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ReadBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"write_bytes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.WriteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"others_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.OthersPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"reads_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ReadsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"writes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.WritesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.BytesPerOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_read\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.BytesPerRead)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_write\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.BytesPerWrite)] = true
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	pc := NewDirectoriesPerformanceCollector(c)
	metricsCheck(t, pc, want)
}
