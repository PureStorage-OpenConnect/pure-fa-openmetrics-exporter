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

func TestPodReplicaLinksPerformanceCollector(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/pod_replica_links_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var prlpl client.PodReplicaLinksPerformanceList
	json.Unmarshal(res, &prlpl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/pod-replica-links/performance/replication$`)
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
	for _, p := range prlpl.Items {
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_sec_to_remote\"} label:{name:\"direction\" value:\"%s\"} label:{name:\"local_pod\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} label:{name:\"remote_pod\" value:\"%s\"} gauge:{value:%g}", p.Direction, p.LocalPod.Name, p.Remotes[0].Name, p.RemotePod.Name, p.BytesPerSecToRemote)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_sec_from_remote\"} label:{name:\"direction\" value:\"%s\"} label:{name:\"local_pod\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} label:{name:\"remote_pod\" value:\"%s\"} gauge:{value:%g}", p.Direction, p.LocalPod.Name, p.Remotes[0].Name, p.RemotePod.Name, p.BytesPerSecFromRemote)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_sec_total\"} label:{name:\"direction\" value:\"%s\"} label:{name:\"local_pod\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} label:{name:\"remote_pod\" value:\"%s\"} gauge:{value:%g}", p.Direction, p.LocalPod.Name, p.Remotes[0].Name, p.RemotePod.Name, p.BytesPerSecTotal)] = true

	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)

	pc := NewPodReplicaLinksPerformanceCollector(c)
	metricsCheck(t, pc, want)
}
