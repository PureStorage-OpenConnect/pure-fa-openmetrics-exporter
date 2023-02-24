package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type PodsPerformanceReplicationCollector struct {
	BandwidthDesc   *prometheus.Desc
	Client          *client.FAClient
}

func (c *PodsPerformanceReplicationCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *PodsPerformanceReplicationCollector) Collect(ch chan<- prometheus.Metric) {
	pods := c.Client.GetPodsPerformanceReplication()
	if len(pods.Items) == 0 {
		return
	}
	for _, p := range pods.Items {
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ContinuousBytesPerSec.FromRemoteBytesPerSec,
			"continuous", "from_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ResyncBytesPerSec.FromRemoteBytesPerSec,
			"resync", "from_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.SyncBytesPerSec.FromRemoteBytesPerSec,
			"sync", "from_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.PeriodicBytesPerSec.FromRemoteBytesPerSec,
			"periodic", "from_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ContinuousBytesPerSec.ToRemoteBytesPerSec,
			"continuous", "to_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ResyncBytesPerSec.ToRemoteBytesPerSec,
			"resync", "to_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.SyncBytesPerSec.ToRemoteBytesPerSec,
			"sync", "to_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.PeriodicBytesPerSec.ToRemoteBytesPerSec,
			"periodic", "to_remote", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ContinuousBytesPerSec.TotalBytesPerSec,
			"continuous", "total", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ResyncBytesPerSec.TotalBytesPerSec,
			"resync", "total", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.SyncBytesPerSec.TotalBytesPerSec,
			"sync", "total", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.PeriodicBytesPerSec.TotalBytesPerSec,
			"periodic", "total", p.Pod.Name,
		)
	}
}

func NewPodsPerformanceReplicationCollector(fa *client.FAClient) *PodsPerformanceReplicationCollector {
	return &PodsPerformanceReplicationCollector{
		BandwidthDesc: prometheus.NewDesc(
			"purefa_pod_performance_replication_bandwidth_bytes",
			"FlashArray pod replication bandwidth",
			[]string{"dimension", "direction", "name"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
