package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type ControllersCollector struct {
	ControllersDesc *prometheus.Desc
	Client          *client.FAClient
}

func (c *ControllersCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *ControllersCollector) Collect(ch chan<- prometheus.Metric) {
	cl := c.Client.GetControllers()
	if len(cl.Items) == 0 {
		return
	}
	for _, d := range cl.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ControllersDesc,
			prometheus.GaugeValue,
			// OpenMetrics timestamps MUST be in seconds, divide as an float64 to keep precision
			(float64(d.ModeSince) / 1000),
			d.Mode, d.Model, d.Name, d.Status, d.Type, d.Version,
		)
	}
}

func NewControllersCollector(fa *client.FAClient) *ControllersCollector {
	return &ControllersCollector{
		ControllersDesc: prometheus.NewDesc(
			"purefa_hw_controller_mode_since_timestamp_seconds",
			"FlashArray controller mode since",
			[]string{"mode", "model", "name", "status", "type", "version"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
