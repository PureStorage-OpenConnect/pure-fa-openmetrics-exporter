package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"purestorage/fa-openmetrics-exporter/internal/openmetrics-exporter"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

var version string = "0.2.0"
var debug bool = false

func main() {

	host := flag.String("host", "0.0.0.0", "Address of the exporter")
	port := flag.Int("port", 9490, "Port of the exporter")
	d := flag.Bool("debug", false, "Debug")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	debug = *d
	log.Printf("Start exporter on %s", addr)

	http.HandleFunc("/", index)
	http.HandleFunc("/metrics/volumes", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics/hosts", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics/pods", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics/directories", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	path := strings.Split(r.URL.Path, "/")
	metrics := ""
	if len(path) == 2 {
		metrics = "all"
	} else {
		metrics = path[2]
		switch metrics {
		case "volumes":
		case "hosts":
		case "pods":
		case "directories":
		default:
			metrics = "all"
		}
	}
	endpoint := params.Get("endpoint")
	if endpoint == "" {
		http.Error(w, "Endpoint parameter is missing", http.StatusBadRequest)
		return
	}
	apiver := params.Get("api-version")
	if apiver == "" {
		apiver = "latest"
	}
	authHeader := r.Header.Get("Authorization")
	authFields := strings.Fields(authHeader)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		http.Error(w, "Target authorization token is missing", http.StatusBadRequest)
		return
	}
	apitoken := authFields[1]

	registry := prometheus.NewRegistry()
	faclient := client.NewRestClient(endpoint, apitoken, apiver, debug)
	collectors.Collector(context.TODO(), metrics, registry, faclient)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
	faclient.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	msg := `<html>
<body>
<h1>Pure Storage Flashblade OpenMetrics Exporter</h1>
<table>
    <thead>
        <tr>
        <td>Type</td>
        <td>Endpoint</td>
        <td>GET parameters</td>
        <td>Description</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Full metrics</td>
            <td><a href="/metrics?endpoint=host">/metrics</a></td>
            <td>endpoint</td>
            <td>All array metrics. Expect slow response time.</td>
        </tr>
        <tr>
            <td>Volumes metrics</td>
            <td><a href="/metrics/volumes?endpoint=host">/metrics/volumes</a></td>
            <td>endpoint</td>
            <td>Provides only volumes related metrics.</td>
        </tr>
        <tr>
            <td>Hosts metrics</td>
            <td><a href="/metrics/hosts?endpoint=host">/metrics/hosts</a></td>
            <td>endpoint</td>
            <td>Provides only hosts related metrics.</td>
        </tr>
        <tr>
            <td>Pods metrics</td>
            <td><a href="/metrics/pods?endpoint=host">/metrics/pods</a></td>
            <td>endpoint</td>
            <td>Provides only pods related metrics.</td>
        </tr>
        <tr>
            <td>Directories metrics</td>
            <td><a href="/metrics/directories?endpoint=host">/metrics/directories</a></td>
            <td>endpoint</td>
            <td>Provides only directories related metrics.</td>
        </tr>
    </tbody>
</table>
</body>
</html>`

	fmt.Fprintf(w, "%s", msg)
}
