package client


import (
	"testing"
        "regexp"
        "strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"os"

	"github.com/google/go-cmp/cmp"
)

func TestAlerts(t *testing.T) {

	ropen, _ := os.ReadFile("../../test/data/alerts_open.json")
	rall, _ := os.ReadFile("../../test/data/alerts.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var aopen AlertsList
	var aall AlertsList
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
        t.Run("alerts_open", func(t *testing.T) {
            c := NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", false)
	    al := c.GetAlerts("state='open'")
	    if diff := cmp.Diff(al.Items, aopen.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        t.Run("alerts_all", func(t *testing.T) {
            c := NewRestClient(e, "fake-api-token", "latest", false)
	    al := c.GetAlerts("")
	    if diff := cmp.Diff(al.Items, aall.Items); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        	server.Close()
            }
        })
        server.Close()
}
