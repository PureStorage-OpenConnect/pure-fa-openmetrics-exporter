package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type NetworkInterfacesPerformanceCollector struct {
	BandwidthDesc  *prometheus.Desc
	ThroughputDesc *prometheus.Desc
	ErrorsDesc     *prometheus.Desc
	Client         *client.FAClient
}

func (c *NetworkInterfacesPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *NetworkInterfacesPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	nwl := c.Client.GetNetworkInterfacesPerformance()
	if len(nwl.Items) == 0 {
		return
	}
	for _, n := range nwl.Items {
		if n.InterfaceType == "eth" {
			ch <- prometheus.MustNewConstMetric(
				c.BandwidthDesc,
				prometheus.GaugeValue,
				n.Eth.ReceivedBytesPerSec,
				n.Name, "received_bytes_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.BandwidthDesc,
				prometheus.GaugeValue,
				n.Eth.TransmittedBytesPerSec,
				n.Name, "transmitted_bytes_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ThroughputDesc,
				prometheus.GaugeValue,
				n.Eth.ReceivedPacketsPerSec,
				n.Name, "received_packets_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ThroughputDesc,
				prometheus.GaugeValue,
				n.Eth.TransmittedPacketsPerSec,
				n.Name, "transmitted_packets_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Eth.OtherErrorsPerSec,
				n.Name, "other_errors_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Eth.ReceivedCrcErrorsPerSec,
				n.Name, "received_crc_errors_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Eth.ReceivedFrameErrorsPerSec,
				n.Name, "received_frame_errors_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Eth.TotalErrorsPerSec,
				n.Name, "total_errors_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Eth.TransmittedCarrierErrorsPerSec,
				n.Name, "transmitted_carrier_errors_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Eth.TransmittedDroppedErrorsPerSec,
				n.Name, "transmitted_dropped_errors_per_sec", n.InterfaceType,
			)
		}
		if n.InterfaceType == "fc" {
			ch <- prometheus.MustNewConstMetric(
				c.BandwidthDesc,
				prometheus.GaugeValue,
				n.Fc.ReceivedBytesPerSec,
				n.Name, "received_bytes_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.BandwidthDesc,
				prometheus.GaugeValue,
				n.Fc.TransmittedBytesPerSec,
				n.Name, "transmitted_bytes_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ThroughputDesc,
				prometheus.GaugeValue,
				n.Fc.ReceivedFramesPerSec,
				n.Name, "received_frames_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ThroughputDesc,
				prometheus.GaugeValue,
				n.Fc.TransmittedFramesPerSec,
				n.Name, "transmitted_frames_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Fc.ReceivedCrcErrorsPerSec,
				n.Name, "received_crc_errors_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Fc.ReceivedLinkFailuresPerSec,
				n.Name, "received_link_failures_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Fc.ReceivedLossOfSignalPerSec,
				n.Name, "received_loss_of_signal_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Fc.ReceivedLossOfSyncPerSec,
				n.Name, "received_loss_of_sync_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Fc.TotalErrorsPerSec,
				n.Name, "total_errors_per_sec", n.InterfaceType,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ErrorsDesc,
				prometheus.GaugeValue,
				n.Fc.TransmittedInvalidWordsPerSec,
				n.Name, "transmitted_invalid_words_per_sec", n.InterfaceType,
			)
		}
	}
}

func NewNetworkInterfacesPerformanceCollector(fa *client.FAClient) *NetworkInterfacesPerformanceCollector {
	return &NetworkInterfacesPerformanceCollector{
		BandwidthDesc: prometheus.NewDesc(
			"purefa_network_interface_performance_bandwidth_bytes",
			"FlashArray network interface bandwidth",
			[]string{"name", "dimension", "type"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_network_interface_performance_throughput_pkts",
			"FlashArray network interface throughput",
			[]string{"name", "dimension", "type"},
			prometheus.Labels{},
		),
		ErrorsDesc: prometheus.NewDesc(
			"purefa_network_interface_performance_errors",
			"FlashArray network interface errors",
			[]string{"name", "dimension", "type"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
