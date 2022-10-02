package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type VolumesSpaceCollector struct {
	ReductionDesc *prometheus.Desc
	SpaceDesc     *prometheus.Desc
	Volumes       *client.VolumesList
}

func (c *VolumesSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *VolumesSpaceCollector) Collect(ch chan<- prometheus.Metric) {
	purenaa := "naa.624a9370"
	volumes := c.Volumes
	if len(volumes.Items) == 0 {
		return
	}
	for _, v := range volumes.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			v.Space.DataReduction,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Shared,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "shared",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Snapshots,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.System,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "system",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.ThinProvisioning,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "thin_provisioning",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalPhysical,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalProvisioned,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalReduction,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Unique,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Virtual,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.Replication,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "replication",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.SharedEffective,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "shared_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.SnapshotsEffective,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "snapshots_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.UniqueEffective,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "unique_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			v.Space.TotalEffective,
			purenaa + v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_effective",
		)
        }
}

func NewVolumesSpaceCollector(volumes *client.VolumesList) *VolumesSpaceCollector {
	return &VolumesSpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefa_volume_space_data_reduction_ratio",
			"FlashArray volume space data reduction",
			[]string{"naa_id", "name", "pod", "volume_group"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefa_volume_space_bytes",
			"FlashArray volume space in bytes",
			[]string{"naa_id", "name", "pod", "volume_group", "space"},
			prometheus.Labels{},
		),
		Volumes: volumes,
	}
}
