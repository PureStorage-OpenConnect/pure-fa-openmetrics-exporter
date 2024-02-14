package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type PortsCollector struct {
	PortInfoDesc *prometheus.Desc
	//	PortiSCSIInfoDesc   *prometheus.Desc
	//	PortNVMeTCPInfoDesc *prometheus.Desc
	//	PortNVMeFCInfoDesc  *prometheus.Desc
	//	PortFCInfoDesc      *prometheus.Desc
	Client *client.FAClient
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
		//		if h.Iqn != "" {
		ch <- prometheus.MustNewConstMetric(
			c.PortInfoDesc,
			prometheus.GaugeValue,
			1,
			h.Iqn, h.Name, h.Nqn, h.Portal, h.Wwn,
		)
		//		}
		//		if h.Nqn != "" && h.Portal != "" {
		//			ch <- prometheus.MustNewConstMetric(
		//				c.PortNVMeTCPInfoDesc,
		//				prometheus.GaugeValue,
		//				1,
		//				h.Name, h.Nqn, h.Portal,
		//			)
		//		}
		//		if h.Nqn != "" && h.Wwn != "" {
		//			ch <- prometheus.MustNewConstMetric(
		//				c.PortNVMeFCInfoDesc,
		//				prometheus.GaugeValue,
		//				1,
		//				h.Name, h.Nqn, h.Wwn,
		//			)
		//		}
		//		if h.Wwn != "" {
		//			ch <- prometheus.MustNewConstMetric(
		//				c.PortFCInfoDesc,
		//				prometheus.GaugeValue,
		//				1,
		//				h.Name, h.Wwn,
		//			)
		//		}
	}
}

func NewPortsCollector(fa *client.FAClient) *PortsCollector {
	return &PortsCollector{
		PortInfoDesc: prometheus.NewDesc(
			"purefa_port_info",
			"FlashArray port info",
			[]string{"iqn", "name", "nqn", "portal", "wwn"},
			prometheus.Labels{},
		),
		//PortiSCSIInfoDesc: prometheus.NewDesc(
		//	"purefa_ports_iscsi_info",
		//	"FlashArray port iscsi info",
		//	[]string{"name", "iqn", "portal"},
		//	prometheus.Labels{},
		//),
		// PortNVMeTCPInfoDesc: prometheus.NewDesc(
		//	"purefa_ports_nvmetcp_info",
		//	"FlashArray ports NVMe/TCP info",
		//	[]string{"name", "nqn", "portal"},
		//	prometheus.Labels{},
		//),
		//PortNVMeFCInfoDesc: prometheus.NewDesc(
		//	"purefa_ports_nvmefc_info",
		//	"FlashArray ports NVMe/FC info",
		//	[]string{"name", "nqn", "wwn"},
		//	prometheus.Labels{},
		//),
		//PortFCInfoDesc: prometheus.NewDesc(
		//	"purefa_ports_fc_info",
		//	"FlashArray ports fc info",
		//	[]string{"name", "wwn"},
		//	prometheus.Labels{},
		//),
		Client: fa,
	}
}
