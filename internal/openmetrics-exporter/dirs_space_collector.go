package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
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

		if d.Space.DataReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.ReductionDesc,
				prometheus.GaugeValue,
				*d.Space.DataReduction,
				d.Name,
			)
		}
		if d.Space.Shared != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.Shared),
				d.Name, "shared",
			)
		}
		if d.Space.Snapshots != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.Snapshots),
				d.Name, "snapshots",
			)
		}
		if d.Space.System != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.System),
				d.Name, "system",
			)
		}
		if d.Space.ThinProvisioning != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				*d.Space.ThinProvisioning,
				d.Name, "thin_provisioning",
			)
		}
		if d.Space.TotalPhysical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.TotalPhysical),
				d.Name, "total_physical",
			)
		}
		if d.Space.TotalProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.TotalProvisioned),
				d.Name, "total_provisioned",
			)
		}
		if d.Space.TotalReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				*d.Space.TotalReduction,
				d.Name, "total_reduction",
			)
		}
		if d.Space.Unique != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.Unique),
				d.Name, "unique",
			)
		}
		if d.Space.Virtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.Virtual),
				d.Name, "virtual",
			)
		}
		if d.Space.Replication != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.Replication),
				d.Name, "replication",
			)
		}
		if d.Space.SharedEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.SharedEffective),
				d.Name, "shared_effective",
			)
		}
		if d.Space.SnapshotsEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.SnapshotsEffective),
				d.Name, "snapshots_effective",
			)
		}
		if d.Space.UniqueEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.UniqueEffective),
				d.Name, "unique_effective",
			)
		}
		if d.Space.TotalEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*d.Space.TotalEffective),
				d.Name, "total_effective",
			)
		}
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
