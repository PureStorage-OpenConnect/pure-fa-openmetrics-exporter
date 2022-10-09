package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type HostConnectionsCollector struct {
	HostConnectionsDesc *prometheus.Desc
	Client              *client.FAClient
}

func (c *HostConnectionsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *HostConnectionsCollector) Collect(ch chan<- prometheus.Metric) {
	hconns := c.Client.GetConnections()
	if len(hconns.Items) == 0 {
		return
	}
	for _, hc := range hconns.Items {
		ch <- prometheus.MustNewConstMetric(
			c.HostConnectionsDesc,
			prometheus.GaugeValue,
			1.0,
			hc.Host.Name, hc.HostGroup.Name, hc.Volume.Name,
		)
	}
}

func NewHostConnectionsCollector(fa *client.FAClient) *HostConnectionsCollector {
	return &HostConnectionsCollector{
		HostConnectionsDesc: prometheus.NewDesc(
			"purefa_host_connections_info",
			"FlashArray host volumes connections",
			[]string{"host", "hostgroup", "volume"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
