package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type HostsPerformanceCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FAClient
}

func (c *HostsPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *HostsPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	hosts := c.Client.GetHostsPerformance()
	if len(hosts.Items) == 0 {
		return
	}
	for _, hp := range hosts.Items {
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.QueueUsecPerMirroredWriteOp),
			hp.Name, "queue_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.QueueUsecPerReadOp),
			hp.Name, "queue_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.QueueUsecPerWriteOp),
			hp.Name, "queue_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.SanUsecPerMirroredWriteOp),
			hp.Name, "san_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.SanUsecPerReadOp),
			hp.Name, "san_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.SanUsecPerWriteOp),
			hp.Name, "san_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.ServiceUsecPerMirroredWriteOp),
			hp.Name, "service_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.ServiceUsecPerReadOp),
			hp.Name, "service_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.ServiceUsecPerWriteOp),
			hp.Name, "service_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.UsecPerMirroredWriteOp),
			hp.Name, "usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.UsecPerReadOp),
			hp.Name, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.UsecPerWriteOp),
			hp.Name, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(hp.ServiceUsecPerReadOpCacheReduction),
			hp.Name, "service_usec_per_read_op_cache_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			float64(hp.MirroredWriteBytesPerSec),
			hp.Name, "mirrored_write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			float64(hp.ReadBytesPerSec),
			hp.Name, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			float64(hp.WriteBytesPerSec),
			hp.Name, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			float64(hp.MirroredWritesPerSec),
			hp.Name, "mirrored_writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			float64(hp.ReadsPerSec),
			hp.Name, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			float64(hp.WritesPerSec),
			hp.Name, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			float64(hp.BytesPerMirroredWrite),
			hp.Name, "bytes_per_mirrored_write",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			float64(hp.BytesPerOp),
			hp.Name, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			float64(hp.BytesPerRead),
			hp.Name, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			float64(hp.BytesPerWrite),
			hp.Name, "bytes_per_write",
		)
	}
}

func NewHostsPerformanceCollector(fa *client.FAClient) *HostsPerformanceCollector {
	return &HostsPerformanceCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefa_host_performance_latency_usec",
			"FlashArray host latency in microseconds",
			[]string{"host", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_host_performance_throughput_iops",
			"FlashArray host throughput in iops",
			[]string{"host", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefa_host_performance_bandwidth_bytes",
			"FlashArray host bandwidth in bytes per second",
			[]string{"host", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefa_host_performance_average_bytes",
			"FlashArray host average operations size in bytes",
			[]string{"host", "dimension"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
