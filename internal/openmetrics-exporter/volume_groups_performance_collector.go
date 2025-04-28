package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type VolumeGroupsPerformanceCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FAClient
}

func (c *VolumeGroupsPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *VolumeGroupsPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	volumeGroupsPerformance := c.Client.GetVolumeGroupsPerformance()
	if len(volumeGroupsPerformance.Items) == 0 {
		return
	}
	for _, vgp := range volumeGroupsPerformance.Items {
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.QueueUsecPerMirroredWriteOp,
			vgp.Name, "queue_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.QueueUsecPerReadOp,
			vgp.Name, "queue_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.QueueUsecPerWriteOp,
			vgp.Name, "queue_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.SanUsecPerMirroredWriteOp,
			vgp.Name, "san_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.SanUsecPerReadOp,
			vgp.Name, "san_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.SanUsecPerWriteOp,
			vgp.Name, "san_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.ServiceUsecPerMirroredWriteOp,
			vgp.Name, "service_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.ServiceUsecPerReadOp,
			vgp.Name, "service_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.ServiceUsecPerWriteOp,
			vgp.Name, "service_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.UsecPerMirroredWriteOp,
			vgp.Name, "usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.UsecPerReadOp,
			vgp.Name, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.UsecPerWriteOp,
			vgp.Name, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.ServiceUsecPerReadOpCacheReduction,
			vgp.Name, "service_usec_per_read_op_cache_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.QosRateLimitUsecPerMirroredWriteOp,
			vgp.Name, "qos_rate_limit_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.QosRateLimitUsecPerReadOp,
			vgp.Name, "qos_rate_limit_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vgp.QosRateLimitUsecPerWriteOp,
			vgp.Name, "qos_rate_limit_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vgp.MirroredWriteBytesPerSec,
			vgp.Name, "mirrored_write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vgp.ReadBytesPerSec,
			vgp.Name, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vgp.WriteBytesPerSec,
			vgp.Name, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vgp.MirroredWritesPerSec,
			vgp.Name, "mirrored_writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vgp.ReadsPerSec,
			vgp.Name, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vgp.WritesPerSec,
			vgp.Name, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vgp.BytesPerMirroredWrite,
			vgp.Name, "bytes_per_mirrored_write",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vgp.BytesPerOp,
			vgp.Name, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vgp.BytesPerRead,
			vgp.Name, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vgp.BytesPerWrite,
			vgp.Name, "bytes_per_write",
		)
	}
}

func NewVolumeGroupsPerformanceCollector(fa *client.FAClient) *VolumeGroupsPerformanceCollector {
	return &VolumeGroupsPerformanceCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefa_volume_group_performance_latency_usec",
			"FlashArray volume group latency",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_volume_group_performance_throughput_iops",
			"FlashArray volume group throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefa_volume_group_performance_bandwidth_bytes",
			"FlashArray volume group throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefa_volume_group_performance_average_bytes",
			"FlashArray volume group average operations size",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
