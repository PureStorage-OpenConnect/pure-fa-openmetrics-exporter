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
			"from_remote", "continuos", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ResyncBytesPerSec.FromRemoteBytesPerSec,
			"from_remote", "resync", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.SyncBytesPerSec.FromRemoteBytesPerSec,
			"from_remote", "sync", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.PeriodicBytesPerSec.FromRemoteBytesPerSec,
			"from_remote", "periodic", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ContinuousBytesPerSec.ToRemoteBytesPerSec,
			"to_remote", "continuos", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ResyncBytesPerSec.ToRemoteBytesPerSec,
			"to_remote", "resync", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.SyncBytesPerSec.ToRemoteBytesPerSec,
			"to_remote", "sync", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.PeriodicBytesPerSec.ToRemoteBytesPerSec,
			"to_remote", "periodic", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ContinuousBytesPerSec.TotalBytesPerSec,
			"total", "continuos", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.ResyncBytesPerSec.TotalBytesPerSec,
			"total", "resync", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.SyncBytesPerSec.TotalBytesPerSec,
			"total", "sync", p.Pod.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
		        p.PeriodicBytesPerSec.TotalBytesPerSec,
			"total", "periodic", p.Pod.Name,
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
