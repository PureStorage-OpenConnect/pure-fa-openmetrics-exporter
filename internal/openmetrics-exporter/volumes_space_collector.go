package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type VolumesSpaceCollector struct {
	ReductionDesc *prometheus.Desc
	SpaceDesc     *prometheus.Desc
	Client        *client.FAClient
}

func (c *VolumesSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *VolumesSpaceCollector) Collect(ch chan<- prometheus.Metric) {
	volumes := c.Client.GetVolumes()
	if len(volumes.Items) == 0 {
		return
	}
	for _, v := range volumes.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			v.Space.DataReduction,
			v.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Shared,
			v.Name, "shared",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Snapshots,
			v.Name, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.System,
			v.Name, "system",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.ThinProvisioning,
			v.Name, "thin_provisioning",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalPhysical,
			v.Name, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalProvisioned,
			v.Name, "total_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalReduction,
			v.Name, "total_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Unique,
			v.Name, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Virtual,
			v.Name, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Replication,
			v.Name, "replication",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.SharedEffective,
			v.Name, "shared_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.SnapshotsEffective,
			v.Name, "snapshots_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.UniqueEffective,
			v.Name, "unique_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalEffective,
			v.Name, "total_effective",
		)
        }
}

func NewVolumesSpaceCollector(fa *client.FAClient) *VolumesSpaceCollector {
	return &VolumesSpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefa_volume_space_data_reduction_ratio",
			"FlashArray volume space data reduction",
			[]string{"name"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefa_volume_space_bytes",
			"FlashArray volume space in bytes",
			[]string{"name", "space"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
