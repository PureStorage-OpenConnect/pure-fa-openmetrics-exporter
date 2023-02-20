package collectors


import (
	"testing"

	"github.com/prometheus/client_model/go"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/google/go-cmp/cmp"
	"fmt"
)

func metricsCheck(t *testing.T, c prometheus.Collector, want map[string]bool) {
        chM := make(chan prometheus.Metric)
        go func() {
                c.Collect(chM)
                close(chM)
        }()
        var buff io_prometheus_client.Metric
        metrics := make(map[string]bool)
        for m := range chM {
                m.Write(&buff)
                metrics[buff.String()] = true
        }
        if diff := cmp.Diff(want, metrics); diff != "" {
                t.Errorf("Mismatch (-want +got):\n%s", diff)
        }
}

func metricsCheckA(t *testing.T, c prometheus.Collector, want map[string]bool) {
        chM := make(chan prometheus.Metric)
        go func() {
                c.Collect(chM)
                close(chM)
        }()
        var buff io_prometheus_client.Metric
        metrics := make(map[string]bool)
        for m := range chM {
                m.Write(&buff)
                metrics[buff.String()] = true
		fmt.Println(buff.String())
        }
}
