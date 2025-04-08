package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type ArrayPerformanceCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	QueueDepthDesc  *prometheus.Desc
	Client          *client.FAClient
}

func (c *ArrayPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *ArrayPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	arrays := c.Client.GetArraysPerformance()
	if len(arrays.Items) == 0 {
		return
	}
	ap := arrays.Items[0]
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.QueueUsecPerMirroredWriteOp,
		"queue_usec_per_mirrored_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.QueueUsecPerReadOp,
		"queue_usec_per_read_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.QueueUsecPerWriteOp,
		"queue_usec_per_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.SanUsecPerMirroredWriteOp,
		"san_usec_per_mirrored_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.SanUsecPerReadOp,
		"san_usec_per_read_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.SanUsecPerWriteOp,
		"san_usec_per_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.ServiceUsecPerMirroredWriteOp,
		"service_usec_per_mirrored_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.ServiceUsecPerReadOp,
		"service_usec_per_read_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.ServiceUsecPerWriteOp,
		"service_usec_per_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.QosRateLimitUsecPerMirroredWriteOp,
		"qos_rate_limit_usec_per_mirrored_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.QosRateLimitUsecPerReadOp,
		"qos_rate_limit_usec_per_read_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.QosRateLimitUsecPerWriteOp,
		"qos_rate_limit_usec_per_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.UsecPerMirroredWriteOp,
		"usec_per_mirrored_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.UsecPerReadOp,
		"usec_per_read_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.UsecPerWriteOp,
		"usec_per_write_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.ServiceUsecPerReadOpCacheReduction,
		"service_usec_per_read_op_cache_reduction",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.LocalQueueUsecPerOp,
		"local_queue_usec_per_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		ap.UsecPerOtherOp,
		"usec_per_other_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.BandwidthDesc,
		prometheus.GaugeValue,
		ap.MirroredWriteBytesPerSec,
		"mirrored_write_bytes_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.BandwidthDesc,
		prometheus.GaugeValue,
		ap.ReadBytesPerSec,
		"read_bytes_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.BandwidthDesc,
		prometheus.GaugeValue,
		ap.WriteBytesPerSec,
		"write_bytes_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		ap.MirroredWritesPerSec,
		"mirrored_writes_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		ap.ReadsPerSec,
		"reads_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		ap.WritesPerSec,
		"writes_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		ap.OthersPerSec,
		"others_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.AverageSizeDesc,
		prometheus.GaugeValue,
		ap.BytesPerMirroredWrite,
		"bytes_per_mirrored_write",
	)
	ch <- prometheus.MustNewConstMetric(
		c.AverageSizeDesc,
		prometheus.GaugeValue,
		ap.BytesPerOp,
		"bytes_per_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.AverageSizeDesc,
		prometheus.GaugeValue,
		ap.BytesPerRead,
		"bytes_per_read",
	)
	ch <- prometheus.MustNewConstMetric(
		c.AverageSizeDesc,
		prometheus.GaugeValue,
		ap.BytesPerWrite,
		"bytes_per_write",
	)
	ch <- prometheus.MustNewConstMetric(
		c.QueueDepthDesc,
		prometheus.GaugeValue,
		ap.QueueDepth,
	)
}

func NewArraysPerformanceCollector(fa *client.FAClient) *ArrayPerformanceCollector {
	return &ArrayPerformanceCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefa_array_performance_latency_usec",
			"FlashArray array latency",
			[]string{"dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_array_performance_throughput_iops",
			"FlashArray array throughput",
			[]string{"dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefa_array_performance_bandwidth_bytes",
			"FlashArray array bandwidth",
			[]string{"dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefa_array_performance_average_bytes",
			"FlashArray array average operations size",
			[]string{"dimension"},
			prometheus.Labels{},
		),
		QueueDepthDesc: prometheus.NewDesc(
			"purefa_array_performance_queue_depth_ops",
			"FlashArray array queue depth size",
			[]string{},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
