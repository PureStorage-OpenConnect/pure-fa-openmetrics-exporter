package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	collectors "purestorage/fa-openmetrics-exporter/internal/openmetrics-exporter"
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var version string = "1.0.2"
var debug bool = false
var allow_secret_parameter bool = false

func main() {

	host := flag.String("host", "0.0.0.0", "Address of the exporter")
	port := flag.Int("port", 9490, "Port of the exporter")
	d := flag.Bool("debug", false, "Debug")
	sp := flag.Bool("secret_parameter", false, "allows api key to be provided as GET parameter")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	debug = *d
	allow_secret_parameter = *sp
	log.Printf("Start Pure FlashArray exporter v%s on %s", version, addr)

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
	http.HandleFunc("/metrics/array", func(w http.ResponseWriter, r *http.Request) {
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
		case "array":
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

	apitoken := ""

	if allow_secret_parameter {
		apitoken = params.Get("api_token")
	}

	// if a get parameter is not supplied, use header
	if apitoken == "" {
		authHeader := r.Header.Get("Authorization")
		authFields := strings.Fields(authHeader)
		if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
			http.Error(w, "Target authorization token is missing", http.StatusBadRequest)
			return
		}
		apitoken = authFields[1]
	}

	registry := prometheus.NewRegistry()
	faclient := client.NewRestClient(endpoint, apitoken, apiver, debug)
	if faclient.Error != nil {
		http.Error(w, "Error connecting to FlashArray. Check your management endpoint and/or api token are correct.", http.StatusBadRequest)
		return
	}
	collectors.Collector(context.TODO(), metrics, registry, faclient)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
	faclient.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	msg := `<html>
<body>
<h1>Pure Storage FlashArray OpenMetrics Exporter</h1>
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
            <td>Array metrics</td>
            <td><a href="/metrics/array?endpoint=host">/metrics/array</a></td>
            <td>endpoint</td>
            <td>Provides only array related metrics.</td>
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
