package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type DirectoriesPerformanceCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FAClient
}

func (c *DirectoriesPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *DirectoriesPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	dirs := c.Client.GetDirectoriesPerformance()
	if len(dirs.Items) == 0 {
		return
	}
	for _, dp := range dirs.Items {
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(dp.UsecPerOtherOp),
			dp.Name, "usec_per_other_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(dp.UsecPerReadOp),
			dp.Name, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			float64(dp.UsecPerWriteOp),
			dp.Name, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			float64(dp.ReadBytesPerSec),
			dp.Name, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			float64(dp.WriteBytesPerSec),
			dp.Name, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			float64(dp.OthersPerSec),
			dp.Name, "others_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			float64(dp.ReadsPerSec),
			dp.Name, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			float64(dp.WritesPerSec),
			dp.Name, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			float64(dp.BytesPerOp),
			dp.Name, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			float64(dp.BytesPerRead),
			dp.Name, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			float64(dp.BytesPerWrite),
			dp.Name, "bytes_per_write",
		)
        }
}

func NewDirectoriesPerformanceCollector(fa *client.FAClient) *DirectoriesPerformanceCollector {
	return &DirectoriesPerformanceCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefa_directory_performance_latency_usec",
			"FlashArray directory latency",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefa_directory_performance_throughput_iops",
			"FlashArray directory throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefa_directory_performance_bandwidth_bytes",
			"FlashArray directory bandwidth",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefa_directory_performance_average_bytes",
			"FlashArray directory average operations size",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
