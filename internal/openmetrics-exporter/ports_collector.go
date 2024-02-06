package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type PortsCollector struct {
	PortiSCSIInfoDesc  *prometheus.Desc
	PortNVMeoFInfoDesc *prometheus.Desc
	PortFCInfoDesc     *prometheus.Desc
	Client             *client.FAClient
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
		if h.Iqn != "" {
			ch <- prometheus.MustNewConstMetric(
				c.PortiSCSIInfoDesc,
				prometheus.GaugeValue,
				1,
				h.Name, h.Iqn, h.Portal,
			)
		}
		if h.Nqn != "" {
			ch <- prometheus.MustNewConstMetric(
				c.PortNVMeoFInfoDesc,
				prometheus.GaugeValue,
				1,
				h.Name, h.Nqn, h.Portal,
			)
		}
		if h.Wwn != "" {
			ch <- prometheus.MustNewConstMetric(
				c.PortFCInfoDesc,
				prometheus.GaugeValue,
				1,
				h.Name, h.Wwn,
			)
		}
	}
}

func NewPortsCollector(fa *client.FAClient) *PortsCollector {
	return &PortsCollector{
		PortiSCSIInfoDesc: prometheus.NewDesc(
			"purefa_ports_iscsi_info",
			"FlashArray port iscsi info",
			[]string{"name", "iqn", "portal"},
			prometheus.Labels{},
		),
		PortNVMeoFInfoDesc: prometheus.NewDesc(
			"purefa_ports_nvmeof_info",
			"FlashArray ports nvmeof info",
			[]string{"name", "nqn", "portal"},
			prometheus.Labels{},
		),
		PortFCInfoDesc: prometheus.NewDesc(
			"purefa_ports_fc_info",
			"FlashArray ports fc info",
			[]string{"name", "wwn"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
