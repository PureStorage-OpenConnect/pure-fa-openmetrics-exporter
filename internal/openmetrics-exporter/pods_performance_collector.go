package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type PodsPerformanceCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FAClient
}

func (c *PodsPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *PodsPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	pods := c.Client.GetPodsPerformance()
	if len(pods.Items) == 0 {
		return
	}
	for _, hp := range pods.Items {
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.QueueUsecPerMirroredWriteOp,
			hp.Name, "queue_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.QueueUsecPerReadOp,
			hp.Name, "queue_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.QueueUsecPerWriteOp,
			hp.Name, "queue_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.SanUsecPerMirroredWriteOp,
			hp.Name, "san_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.SanUsecPerReadOp,
			hp.Name, "san_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.SanUsecPerWriteOp,
			hp.Name, "san_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.ServiceUsecPerMirroredWriteOp,
			hp.Name, "service_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.ServiceUsecPerReadOp,
			hp.Name, "service_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.ServiceUsecPerWriteOp,
			hp.Name, "service_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.UsecPerMirroredWriteOp,
			hp.Name, "usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.UsecPerReadOp,
			hp.Name, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.UsecPerWriteOp,
			hp.Name, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.ServiceUsecPerReadOpCacheReduction,
			hp.Name, "service_usec_per_read_op_cache_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			hp.UsecPerOtherOp,
			hp.Name, "usec_per_other_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			hp.MirroredWriteBytesPerSec,
			hp.Name, "mirrored_write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			hp.ReadBytesPerSec,
			hp.Name, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			hp.WriteBytesPerSec,
			hp.Name, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			hp.MirroredWritesPerSec,
			hp.Name, "mirrored_writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			hp.ReadsPerSec,
			hp.Name, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			hp.WritesPerSec,
			hp.Name, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			hp.OthersPerSec,
			hp.Name, "others_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			hp.BytesPerMirroredWrite,
			hp.Name, "bytes_per_mirrored_write",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			hp.BytesPerOp,
			hp.Name, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			hp.BytesPerRead,
			hp.Name, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			hp.BytesPerWrite,
			hp.Name, "bytes_per_write",
		)
	}
}

func NewPodsPerformanceCollector(fa *client.FAClient) *PodsPerformanceCollector {
	return &PodsPerformanceCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefa_pod_performance_latency_usec",
			"FlashArray pod latency",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_pod_performance_throughput_iops",
			"FlashArray pod throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefa_pod_performance_bandwidth_bytes",
			"FlashArray pod throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefa_pod_performance_average_bytes",
			"FlashArray pod average operations size",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
