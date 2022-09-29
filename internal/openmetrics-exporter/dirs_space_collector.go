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
	for _, ds := range dirs.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			ds.Space.DataReduction,
			ds.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.Shared,
			ds.Name, "shared",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.Snapshots,
			ds.Name, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.System,
			ds.Name, "system",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.ThinProvisioning,
			ds.Name, "thin_provisioning",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.TotalPhysical,
			ds.Name, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.TotalProvisioned,
			ds.Name, "total_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.Unique,
			ds.Name, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.Virtual,
			ds.Name, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.Replication,
			ds.Name, "replication",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.SharedEffective,
			ds.Name, "shared_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.SnapshotsEffective,
			ds.Name, "snapshots_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.UniqueEffective,
			ds.Name, "unique_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			ds.Space.TotalEffective,
			ds.Name, "total_effective",
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
