package main

import (
	"context"
	"crypto/tls"
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

var version string = "development"
var debug bool = false
var secure bool = false
var arraytokens config.FlashArrayList

func fileExists(args []string) error {
	_, err := os.Stat(args[0])
	return err
}

func isFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {

	parser := argparse.NewParser("pure-fa-om-exporter", "Pure Storage FA OpenMetrics exporter")
	host := parser.String("a", "address", &argparse.Options{Required: false, Help: "IP address for this exporter to bind to", Default: "0.0.0.0"})
	port := parser.Int("p", "port", &argparse.Options{Required: false, Help: "Port for this exporter to listen", Default: 9490})
	d := parser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Enable debug", Default: false})
	s := parser.Flag("s", "secure", &argparse.Options{Required: false, Help: "Enable TLS verification when connecting to array", Default: false})
	at := parser.File("t", "tokens", os.O_RDONLY, 0600, &argparse.Options{Required: false, Validate: fileExists, Help: "API token(s) map file"})
	cert := parser.String("c", "cert", &argparse.Options{Required: false, Help: "SSL/TLS certificate file. Required only for Exporter TLS"})
	key := parser.String("k", "key", &argparse.Options{Required: false, Help: "SSL/TLS private key file. Required only for Exporter TLS"})
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
	if (len(*cert) > 0 && len(*key) == 0) || (len(*cert) == 0 && len(*key) > 0) {
		log.Fatal("Both certificate and key must be specified to enable TLS")
	}
	if len(*cert) > 0 && len(*key) > 0 {
		if !isFile(*cert) {
			log.Fatal("TLS cert file not found")
		} else if !isFile(*key) {
			log.Fatal("TLS key file not found")
		}
	}
	debug = *d
	secure = *s
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Start Pure FlashArray exporter %s on %s", version, addr)

	if isFile(*cert) && isFile(*key) {

		cfg := &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			},
		}

		srv := &http.Server{
			TLSConfig: cfg,
			Addr:      addr,
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			index(w, r)
		})
		http.HandleFunc("/metrics/volumes", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			metricsHandler(w, r)
		})
		http.HandleFunc("/metrics/hosts", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			metricsHandler(w, r)
		})
		http.HandleFunc("/metrics/pods", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			metricsHandler(w, r)
		})
		http.HandleFunc("/metrics/directories", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			metricsHandler(w, r)
		})
		http.HandleFunc("/metrics/array", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			metricsHandler(w, r)
		})
		http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			metricsHandler(w, r)
		})
		log.Fatal(srv.ListenAndServeTLS(*cert, *key))
	} else {
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
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if debug {
		log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.Header.Get("User-Agent"))
	}
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
		log.Printf("[ERROR] %s %s %s HTTP REQUEST ERROR: Endpoint parameter is missing\n", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "Endpoint parameter is missing", http.StatusBadRequest)
		return
	}
	apiver := params.Get("api-version")
	if apiver == "" {
		apiver = "latest"
	}
	authHeader := r.Header.Get("Authorization")
	authFields := strings.Fields(authHeader)

	address, apitoken := arraytokens.GetArrayParams(endpoint)
	if len(authFields) == 2 && strings.ToLower(authFields[0]) == "bearer" {
		apitoken = authFields[1]
		address = endpoint
	}
	if apitoken == "" {
		log.Printf("[ERROR] %s %s %s HTTP REQUEST ERROR: Target authorization token is missing\n", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "Target authorization token is missing", http.StatusBadRequest)
		return
	}

	uagent := r.Header.Get("User-Agent")
	rid := r.Header.Get("X-Request-ID")

	registry := prometheus.NewRegistry()
	faclient := client.NewRestClient(address, apitoken, apiver, uagent, rid, debug, secure)
	if faclient.Error != nil {
		log.Printf("[ERROR] %s %s %s %s FACLIENT ERROR: %s\n", r.RemoteAddr, r.Method, r.URL, r.Header.Get("User-Agent"), faclient.Error.Error())
		http.Error(w, faclient.Error.Error(), http.StatusBadRequest)
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
