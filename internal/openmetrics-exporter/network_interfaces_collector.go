package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type NetworkInterfacesCollector struct {
	NetworkInterfaceInfoDesc *prometheus.Desc
	Client                   *client.FAClient
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
		ch <- prometheus.MustNewConstMetric(
			c.NetworkInterfaceInfoDesc,
			prometheus.GaugeValue,
			float64(h.Speed),
			h.Name, strconv.FormatBool(h.Enabled), strings.Join(h.Services, ", "), h.InterfaceType, h.Eth.Subtype,
		)
	}
}

func NewNetworkInterfacesCollector(fa *client.FAClient) *NetworkInterfacesCollector {
	return &NetworkInterfacesCollector{
		NetworkInterfaceInfoDesc: prometheus.NewDesc(
			"purefa_network_interface_speed_bandwidth_bytes",
			"FlashArray network interface speed in bytes per second",
			[]string{"name", "enabled", "services", "type", "ethsubtype"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
