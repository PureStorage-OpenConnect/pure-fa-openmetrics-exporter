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

func TestPodsPerformanceReplicationCollector(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/pods_performance_replication.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var pods client.PodsPerformanceReplicationList
	json.Unmarshal(res, &pods)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/pods/performance/replication$`)
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
	for _, p := range pods.Items {
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"continuous\"} label:{name:\"direction\" value:\"from_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.ContinuousBytesPerSec.FromRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"resync\"} label:{name:\"direction\" value:\"from_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.ResyncBytesPerSec.FromRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"sync\"} label:{name:\"direction\" value:\"from_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.SyncBytesPerSec.FromRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"periodic\"} label:{name:\"direction\" value:\"from_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.PeriodicBytesPerSec.FromRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"continuous\"} label:{name:\"direction\" value:\"to_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.ContinuousBytesPerSec.ToRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"resync\"} label:{name:\"direction\" value:\"to_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.ResyncBytesPerSec.ToRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"sync\"} label:{name:\"direction\" value:\"to_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.SyncBytesPerSec.ToRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"periodic\"} label:{name:\"direction\" value:\"to_remote\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.PeriodicBytesPerSec.ToRemoteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"continuous\"} label:{name:\"direction\" value:\"total\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.ContinuousBytesPerSec.TotalBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"resync\"} label:{name:\"direction\" value:\"total\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.ResyncBytesPerSec.TotalBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"sync\"} label:{name:\"direction\" value:\"total\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.SyncBytesPerSec.TotalBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"periodic\"} label:{name:\"direction\" value:\"total\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Pod.Name, p.PeriodicBytesPerSec.TotalBytesPerSec)] = true
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	pc := NewPodsPerformanceReplicationCollector(c)
	metricsCheck(t, pc, want)
}
