package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type VolumesPerformanceCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	NAAids          map[string]string
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
			c.NAAids[vp.Name], vp.Name, "queue_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.QueueUsecPerReadOp,
			c.NAAids[vp.Name], vp.Name, "queue_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.QueueUsecPerWriteOp,
			c.NAAids[vp.Name], vp.Name, "queue_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.SanUsecPerMirroredWriteOp,
			c.NAAids[vp.Name], vp.Name, "san_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.SanUsecPerReadOp,
			c.NAAids[vp.Name], vp.Name, "san_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.SanUsecPerWriteOp,
			c.NAAids[vp.Name], vp.Name, "san_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerMirroredWriteOp,
			c.NAAids[vp.Name], vp.Name, "service_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerReadOp,
			c.NAAids[vp.Name], vp.Name, "service_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerWriteOp,
			c.NAAids[vp.Name], vp.Name, "service_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.UsecPerMirroredWriteOp,
			c.NAAids[vp.Name], vp.Name, "usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.UsecPerReadOp,
			c.NAAids[vp.Name], vp.Name, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.UsecPerWriteOp,
			c.NAAids[vp.Name], vp.Name, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			vp.ServiceUsecPerReadOpCacheReduction,
			c.NAAids[vp.Name], vp.Name, "service_usec_per_read_op_cache_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(*vp.QosRateLimitUsecPerMirroredWriteOp),
			c.NAAids[vp.Name], vp.Name, "qos_rate_limit_usec_per_mirrored_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(*vp.QosRateLimitUsecPerReadOp),
			c.NAAids[vp.Name], vp.Name, "qos_rate_limit_usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(*vp.QosRateLimitUsecPerWriteOp),
			c.NAAids[vp.Name], vp.Name, "qos_rate_limit_usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vp.MirroredWriteBytesPerSec,
			c.NAAids[vp.Name], vp.Name, "mirrored_write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vp.ReadBytesPerSec,
			c.NAAids[vp.Name], vp.Name, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			vp.WriteBytesPerSec,
			c.NAAids[vp.Name], vp.Name, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vp.MirroredWritesPerSec,
			c.NAAids[vp.Name], vp.Name, "mirrored_writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vp.ReadsPerSec,
			c.NAAids[vp.Name], vp.Name, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			vp.WritesPerSec,
			c.NAAids[vp.Name], vp.Name, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerMirroredWrite,
			c.NAAids[vp.Name], vp.Name, "bytes_per_mirrored_write",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerOp,
			c.NAAids[vp.Name], vp.Name, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerRead,
			c.NAAids[vp.Name], vp.Name, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			vp.BytesPerWrite,
			c.NAAids[vp.Name], vp.Name, "bytes_per_write",
		)
	}
}

func NewVolumesPerformanceCollector(fa *client.FAClient, volumes *client.VolumesList) *VolumesPerformanceCollector {
	purenaa := "naa.624a9370"
	naaid := make(map[string]string)
	for _, v := range volumes.Items {
		naaid[v.Name] = purenaa + v.Serial
	}

	return &VolumesPerformanceCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefa_volume_performance_latency_usec",
			"FlashArray volume latency",
			[]string{"naa_id", "name", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_volume_performance_throughput_iops",
			"FlashArray volume throughput",
			[]string{"naa_id", "name", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefa_volume_performance_bandwidth_bytes",
			"FlashArray volume throughput",
			[]string{"naa_id", "name", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefa_volume_performance_average_bytes",
			"FlashArray volume average operations size",
			[]string{"naa_id", "name", "dimension"},
			prometheus.Labels{},
		),
		NAAids: naaid,
		Client: fa,
	}
}
