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

func TestPodReplicaLinksLagCollector(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/pod_replica_links_lag.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var prlpl client.PodReplicaLinksLagList
	json.Unmarshal(res, &prlpl)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/pod-replica-links/lag$`)
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
                want[fmt.Sprintf("label:{name:\"direction\" value:\"%s\"} label:{name:\"local_pod\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} label:{name:\"remote_pod\" value:\"%s\"} label:{name:\"status\" value:\"%s\"} gauge:{value:%g}", p.Direction, p.LocalPod.Name, p.Remotes[0].Name, p.RemotePod.Name, p.Status, p.Lag.Avg)] = true
                want[fmt.Sprintf("label:{name:\"direction\" value:\"%s\"} label:{name:\"local_pod\" value:\"%s\"} label:{name:\"remote\" value:\"%s\"} label:{name:\"remote_pod\" value:\"%s\"} label:{name:\"status\" value:\"%s\"} gauge:{value:%g}", p.Direction, p.LocalPod.Name, p.Remotes[0].Name, p.RemotePod.Name, p.Status, p.Lag.Max)] = true

	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	pc := NewPodReplicaLinksLagCollector(c)
	metricsCheck(t, pc, want)
}
