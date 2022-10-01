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

func TestArraysCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
        var arrs client.ArraysList
        json.Unmarshal(res, &arrs)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays$`)
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
        for _, a := range arrs.Items {
                want[fmt.Sprintf("label:<name:\"array_name\" value:\"%s\" > label:<name:\"os\" value:\"%s\" > label:<name:\"system_id\" value:\"%s\" > label:<name:\"version\" value:\"%s\" > gauge:<value:1 > ", a.Name, a.Os, a.Id, a.Version)] = true
        }
        defer server.Close()
	c := client.NewRestClient(e, "fake-api-token", "latest", false)
        ac := NewArraysCollector(c)
	metricsCheck(t, ac, want)
}
