package collectors

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
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
