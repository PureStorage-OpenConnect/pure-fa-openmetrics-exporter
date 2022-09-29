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

func TestPods(t *testing.T) {

	res, _ := ioutil.ReadFile("../../test/data/pods.json")
	vers, _ := ioutil.ReadFile("../../test/data/versions.json")
	var pods PodsList
	json.Unmarshal(res, &pods)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	        valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/pods$`)
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
        t.Run("pods_1", func(t *testing.T) {
            defer server.Close()
            c := NewRestClient(e, "fake-api-token", "latest")
	    pl := c.GetPods()
	    if diff := cmp.Diff(pl.Items, pods.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
            }
        })
}
