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

func TestAlertsCollector(t *testing.T) {

	ropen, _ := os.ReadFile("../../test/data/alerts_open.json")
	rall, _ := os.ReadFile("../../test/data/alerts.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var aopen client.AlertsList
	var aall client.AlertsList
	json.Unmarshal(ropen, &aopen)
	json.Unmarshal(rall, &aall)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        urlall := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/alerts$`)
	        urlopen := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/alerts\?filter=state%3D%27open%27$`)
                if r.URL.Path == "/api/api_version" {
                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
                } else if urlopen.MatchString(r.URL.Path + "?" + r.URL.RawQuery) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ropen))
                } else if urlall.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(rall))
		}
	   }))
        endp := strings.Split(server.URL, "/")
        e := endp[len(endp)-1]
	want := make(map[string]bool)
	for _, a := range aopen.Items {
		want[fmt.Sprintf("label:<name:\"component_name\" value:\"%s\" > label:<name:\"component_type\" value:\"%s\" > label:<name:\"severity\" value:\"%s\" > gauge:<value:1 > ", a.ComponentName, a.ComponentType, a.Severity)] = true
	}
        c := client.NewRestClient(e, "fake-api-token", "latest", false)
	ac := NewAlertsCollector(c)
        metricsCheck(t, ac, want)
        server.Close()
}
