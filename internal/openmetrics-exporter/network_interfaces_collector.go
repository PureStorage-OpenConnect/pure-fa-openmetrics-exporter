package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type NetworkInterfacesCollector struct {
	NetworkInterfaceEthInfoDesc *prometheus.Desc
	NetworkInterfaceFcInfoDesc  *prometheus.Desc
	Client                      *client.FAClient
}

func (c *NetworkInterfacesCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *NetworkInterfacesCollector) Collect(ch chan<- prometheus.Metric) {
	nwl := c.Client.GetNetworkInterfaces()
	if len(nwl.Items) == 0 {
		return
	}
	for _, h := range nwl.Items {
		//		if h.InterfaceType == "eth" {
		ch <- prometheus.MustNewConstMetric(
			c.NetworkInterfaceEthInfoDesc,
			prometheus.GaugeValue,
			1,
			h.Name, strconv.FormatBool(h.Enabled), strings.Join(h.Services, ", "), h.InterfaceType, h.Eth.Subtype,
		)
		//		}
		//		if h.InterfaceType == "fc" {
		//			ch <- prometheus.MustNewConstMetric(
		//				c.NetworkInterfaceFcInfoDesc,
		//				prometheus.GaugeValue,
		//				1,
		//				h.Name, strconv.FormatBool(h.Enabled), strings.Join(h.Services, ", "), h.InterfaceType,
		//			)
		//		}
	}
}

func NewNetworkInterfacesCollector(fa *client.FAClient) *NetworkInterfacesCollector {
	return &NetworkInterfacesCollector{
		NetworkInterfaceEthInfoDesc: prometheus.NewDesc(
			//			"purefa_network_interface_eth_info",
			"purefa_network_interface_info",
			"FlashArray network interface ethernet info",
			[]string{"name", "enabled", "services", "interfacetype", "ethernetinterfacesubtype"},
			prometheus.Labels{},
		),
		//		NetworkInterfaceFcInfoDesc: prometheus.NewDesc(
		//			"purefa_network_interface_fc_info",
		//			"FlashArray network interface fc info",
		//			[]string{"name", "enabled", "services", "interfacetype"},
		//			prometheus.Labels{},
		//		),
		Client: fa,
	}
}
