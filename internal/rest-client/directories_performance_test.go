package client


import (
	"testing"
        "regexp"
        "strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"io/ioutil"

	"github.com/google/go-cmp/cmp"
)

func TestDirectoriesPerformance(t *testing.T) {

	res, _ := ioutil.ReadFile("../../test/data/directories_performance.json")
	vers, _ := ioutil.ReadFile("../../test/data/versions.json")
	var dirsp DirectoriesPerformanceList
	json.Unmarshal(res, &dirsp)
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
        t.Run("directories_performance_1", func(t *testing.T) {
            defer server.Close()
            c := NewRestClient(e, "fake-api-token", "latest")
	    dpl := c.GetDirectoriesPerformance()
	    if diff := cmp.Diff(dpl.Items, dirsp.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
            }
        })
}
