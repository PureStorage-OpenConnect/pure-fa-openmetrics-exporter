package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type ControllersCollector struct {
	ControllersModeSinceDesc *prometheus.Desc
	ControllersInfoDesc      *prometheus.Desc
	Client                   *client.FAClient
}

func (c *ControllersCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *ControllersCollector) Collect(ch chan<- prometheus.Metric) {
	cl := c.Client.GetControllers()
	if len(cl.Items) == 0 {
		return
	}
	for _, ctl := range cl.Items {
		if ctl.ModeSince != 0 {
			ch <- prometheus.MustNewConstMetric(
				c.ControllersModeSinceDesc,
				prometheus.GaugeValue,
				// OpenMetrics timestamps MUST be in seconds, divide as an float64 to keep precision
				(float64(ctl.ModeSince) / 1000),
				ctl.Mode, ctl.Model, ctl.Name, ctl.Status, ctl.Type, ctl.Version,
			)
		}
		ch <- prometheus.MustNewConstMetric(
			c.ControllersInfoDesc,
			prometheus.GaugeValue,
			1,
			ctl.Mode, ctl.Model, ctl.Name, ctl.Status, ctl.Type, ctl.Version,
		)

	}
}

func NewControllersCollector(fa *client.FAClient) *ControllersCollector {
	return &ControllersCollector{
		ControllersModeSinceDesc: prometheus.NewDesc(
			"purefa_hw_controller_mode_since_timestamp_seconds",
			"FlashArray controller mode since change timestamp in seconds since UNIX epoch",
			[]string{"mode", "model", "name", "status", "type", "version"},
			prometheus.Labels{},
		),
		ControllersInfoDesc: prometheus.NewDesc(
			"purefa_hw_controller_info",
			"FlashArray controller info",
			[]string{"mode", "model", "name", "status", "type", "version"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
