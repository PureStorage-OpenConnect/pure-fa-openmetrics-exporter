package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type VolumesPerformanceCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FAClient
}

func (c *VolumesPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *VolumesPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	volumes := c.Client.GetVolumesPerformance()
	if len(volumes.Items) == 0 {
		return
	}
	for _, vp := range volumes.Items {
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.QueueUsecPerMirroredWriteOp,
			vp.Name, "queue_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.QueueUsecPerReadOp,
			vp.Name, "queue_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.QueueUsecPerWriteOp,
			vp.Name, "queue_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.SanUsecPerMirroredWriteOp,
			vp.Name, "san_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.SanUsecPerReadOp,
			vp.Name, "san_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.SanUsecPerWriteOp,
			vp.Name, "san_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerMirroredWriteOp,
			vp.Name, "service_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerReadOp,
			vp.Name, "service_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerWriteOp,
			vp.Name, "service_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.UsecPerMirroredWriteOp,
			vp.Name, "usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.UsecPerReadOp,
			vp.Name, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.UsecPerWriteOp,
			vp.Name, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerReadOpCacheReduction,
			vp.Name, "service_usec_per_read_op_cache_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vp.MirroredWriteBytesPerSec,
			vp.Name, "mirrored_write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vp.ReadBytesPerSec,
			vp.Name, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vp.WriteBytesPerSec,
			vp.Name, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vp.MirroredWritesPerSec,
			vp.Name, "mirrored_writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vp.ReadsPerSec,
			vp.Name, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vp.WritesPerSec,
			vp.Name, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerMirroredWrite,
			vp.Name, "bytes_per_mirrored_write",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerOp,
			vp.Name, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerRead,
			vp.Name, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerWrite,
			vp.Name, "bytes_per_write",
		)
	}
}

func NewVolumesPerformanceCollector(fa *client.FAClient) *VolumesPerformanceCollector {
	return &VolumesPerformanceCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefa_volume_performance_latency_usec",
			"FlashArray volume latency",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_volume_performance_throughput_iops",
			"FlashArray volume throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefa_volume_performance_bandwidth_bytes",
			"FlashArray volume throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefa_volume_performance_average_bytes",
			"FlashArray volume average operations size",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
