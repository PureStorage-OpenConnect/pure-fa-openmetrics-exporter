package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	config "purestorage/fa-openmetrics-exporter/internal/config"
	collectors "purestorage/fa-openmetrics-exporter/internal/openmetrics-exporter"
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v3"
)

var version string = "1.0.5"
var debug bool = false
var arraytokens config.FlashArrayList

func FileExists(args []string) error {
	_, err := os.Stat(args[0])
	return err
}

func main() {

	parser := argparse.NewParser("pure-fa-om-exporter", "Pure Storage FA OpenMetrics exporter")
	host := parser.String("a", "address", &argparse.Options{Required: false, Help: "IP address for this exporter to bind to", Default: "0.0.0.0"})
	port := parser.Int("p", "port", &argparse.Options{Required: false, Help: "Port for this exporter to listen", Default: 9490})
	d := parser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Enable debug", Default: false})
	at := parser.File("t", "tokens", os.O_RDONLY, 0600, &argparse.Options{Required: false, Validate: FileExists, Help: "API token(s) map file"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatalf("Error in token file: %v", err)
	}
	if !isNilFile(*at) {
		defer at.Close()
		buf := make([]byte, 1024)
		arrlist := ""
		for {
			n, err := at.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Reading token file: %v", err)
			}
			if n > 0 {
				arrlist = arrlist + string(buf[:n])
			}
		}
		buf = []byte(arrlist)
		err := yaml.Unmarshal(buf, &arraytokens)
		if err != nil {
			log.Fatalf("Unmarshalling token file: %v", err)
		}
	}
	debug = *d
	addr := fmt.Sprintf("%s:%d", *host, *port)
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
	endpoint := params.Get("target")
	if endpoint == "" {
		http.Error(w, "Target parameter is missing", http.StatusBadRequest)
		return
	}
	apiver := params.Get("api-version")
	if apiver == "" {
		apiver = "latest"
	}
	authHeader := r.Header.Get("Authorization")
	authFields := strings.Fields(authHeader)

	apitoken := arraytokens.GetApiToken(endpoint)
	if len(authFields) == 2 && strings.ToLower(authFields[0]) == "bearer" {
		apitoken = authFields[1]
	}
	if apitoken == "" {
		http.Error(w, "Target authorization token is missing", http.StatusBadRequest)
		return
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

func isNilFile(f os.File) bool {
	var tf os.File
	return f == tf
}
