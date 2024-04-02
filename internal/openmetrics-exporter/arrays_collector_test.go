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

func TestArraysCollector(t *testing.T) {

	res, _ := os.ReadFile("../../test/data/arrays.json")
	ressub, _ := os.ReadFile("../../test/data/subscriptions.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var arrs client.ArraysList
	var subs client.SubscriptionsList
	json.Unmarshal(res, &arrs)
	json.Unmarshal(ressub, &subs)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlarr := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/arrays$`)
		urlsubs := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/subscriptions$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if urlarr.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(res))
		} else if urlsubs.MatchString(r.URL.Path) {
			w.Header().Set("x-auth-token", "faketoken")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ressub))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	a := arrs.Items[0]
	s := subs.Items[0]
	want := make(map[string]bool)

	want[fmt.Sprintf("label:{name:\"array_name\" value:\"%s\"} label:{name:\"os\" value:\"%s\"} label:{name:\"subscription_type\" value:\"%s\"} label:{name:\"system_id\" value:\"%s\"} label:{name:\"version\" value:\"%s\"} gauge:{value:1}", a.Name, a.Os, s.Service, a.Id, a.Version)] = true

	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false)
	ac := NewArraysCollector(c)
	metricsCheck(t, ac, want)
	server.Close()
}
