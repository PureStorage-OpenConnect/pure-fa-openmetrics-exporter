package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type DirectoriesSpaceCollector struct {
	ReductionDesc *prometheus.Desc
	SpaceDesc     *prometheus.Desc
	Client        *client.FAClient
}

func (c *DirectoriesSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *DirectoriesSpaceCollector) Collect(ch chan<- prometheus.Metric) {
	dirs := c.Client.GetDirectories()
	if len(dirs.Items) == 0 {
		return
	}
	for _, d := range dirs.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			d.Space.DataReduction,
			d.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.Shared,
			d.Name, "shared",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.Snapshots,
			d.Name, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.System,
			d.Name, "system",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.ThinProvisioning,
			d.Name, "thin_provisioning",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.TotalPhysical,
			d.Name, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.TotalProvisioned,
			d.Name, "total_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.TotalReduction,
			d.Name, "total_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.Unique,
			d.Name, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.Virtual,
			d.Name, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.Replication,
			d.Name, "replication",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.SharedEffective,
			d.Name, "shared_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.SnapshotsEffective,
			d.Name, "snapshots_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.UniqueEffective,
			d.Name, "unique_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			d.Space.TotalEffective,
			d.Name, "total_effective",
		)
        }
}

func NewDirectoriesSpaceCollector(fa *client.FAClient) *DirectoriesSpaceCollector {
	return &DirectoriesSpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefa_directory_space_data_reduction_ratio",
			"FlashArray directory space data reduction",
			[]string{"name"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefa_directory_space_bytes",
			"FlashArray directory space in bytes",
			[]string{"name", "space"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
