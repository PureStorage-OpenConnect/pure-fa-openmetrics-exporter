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

func TestNetworkInterfacesPerformanceCollector(t *testing.T) {
	res, _ := os.ReadFile("../../test/data/network_interfaces_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var  nics client.NetworkInterfacesPerformanceList
	json.Unmarshal(res, &nics)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                valid := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/network-interfaces/performance$`)
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
	for _, n := range nics.Items {
		if n.InterfaceType == "eth" {
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_bytes_per_sec", n.Name, n.InterfaceType, n.Eth.ReceivedBytesPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "transmitted_bytes_per_sec", n.Name, n.InterfaceType, n.Eth.TransmittedBytesPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_packets_per_sec", n.Name, n.InterfaceType, n.Eth.ReceivedPacketsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "transmitted_packets_per_sec", n.Name, n.InterfaceType, n.Eth.TransmittedPacketsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "other_errors_per_sec", n.Name, n.InterfaceType, n.Eth.OtherErrorsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_crc_errors_per_sec", n.Name, n.InterfaceType, n.Eth.ReceivedCrcErrorsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_frame_errors_per_sec", n.Name, n.InterfaceType, n.Eth.ReceivedFrameErrorsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "total_errors_per_sec", n.Name, n.InterfaceType, n.Eth.TotalErrorsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "transmitted_carrier_errors_per_sec", n.Name, n.InterfaceType, n.Eth.TransmittedCarrierErrorsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "transmitted_dropped_errors_per_sec", n.Name, n.InterfaceType, n.Eth.TransmittedDroppedErrorsPerSec)] = true
		}
		if n.InterfaceType == "fc" {
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_bytes_per_sec", n.Name, n.InterfaceType, n.Fc.ReceivedBytesPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "transmitted_bytes_per_sec", n.Name, n.InterfaceType, n.Fc.TransmittedBytesPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_frames_per_sec", n.Name, n.InterfaceType, n.Fc.ReceivedFramesPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "transmitted_frames_per_sec", n.Name, n.InterfaceType, n.Fc.TransmittedFramesPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_crc_errors_per_sec", n.Name, n.InterfaceType, n.Fc.ReceivedCrcErrorsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_link_failures_per_sec", n.Name, n.InterfaceType, n.Fc.ReceivedLinkFailuresPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_loss_of_signal_per_sec", n.Name, n.InterfaceType, n.Fc.ReceivedLossOfSignalPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "received_loss_of_sync_per_sec", n.Name, n.InterfaceType, n.Fc.ReceivedLossOfSyncPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "total_errors_per_sec", n.Name, n.InterfaceType, n.Fc.TotalErrorsPerSec)] = true
			want[fmt.Sprintf("label:{name:\"dimension\" value:\"%s\"} label:{name:\"name\" value:\"%s\"} label:{name:\"type\" value:\"%s\"} gauge:{value:%g}", "transmitted_invalid_words_per_sec", n.Name, n.InterfaceType, n.Fc.TransmittedInvalidWordsPerSec)] = true
		}
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", false)
	pc := NewNetworkInterfacesPerformanceCollector(c)
	metricsCheck(t, pc, want)
}
