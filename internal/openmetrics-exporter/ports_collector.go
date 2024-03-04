package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type PortsCollector struct {
	PortInfoDesc *prometheus.Desc
	Client       *client.FAClient
}

func (c *PortsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *PortsCollector) Collect(ch chan<- prometheus.Metric) {
	hwl := c.Client.GetPorts()
	if len(hwl.Items) == 0 {
		return
	}
	for _, h := range hwl.Items {
		ch <- prometheus.MustNewConstMetric(
			c.PortInfoDesc,
			prometheus.GaugeValue,
			1,
			h.Iqn, h.Name, h.Nqn, h.Portal, h.Wwn,
		)
	}
}

func NewPortsCollector(fa *client.FAClient) *PortsCollector {
	return &PortsCollector{
		PortInfoDesc: prometheus.NewDesc(
			"purefa_network_port_info",
			"FlashArray network port info",
			[]string{"iqn", "name", "nqn", "portal", "wwn"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
