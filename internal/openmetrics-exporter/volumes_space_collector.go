package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
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
		if v.Space.DataReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.ReductionDesc,
				prometheus.GaugeValue,
				*v.Space.DataReduction,
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name,
			)
		}
		if v.Space.Shared != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.Shared),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "shared",
			)
		}
		if v.Space.Snapshots != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.Snapshots),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "snapshots",
			)
		}
		if v.Space.System != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.System),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "system",
			)
		}
		if v.Space.ThinProvisioning != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				*v.Space.ThinProvisioning,
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "thin_provisioning",
			)
		}
		if v.Space.TotalPhysical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.TotalPhysical),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_physical",
			)
		}
		if v.Space.TotalProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.TotalProvisioned),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_provisioned",
			)
		}
		if v.Space.TotalReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				*v.Space.TotalReduction,
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_reduction",
			)
		}
		if v.Space.Unique != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.Unique),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "unique",
			)
		}
		if v.Space.Virtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.Virtual),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "virtual",
			)
		}
		if v.Space.Replication != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.Replication),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "replication",
			)
		}
		if v.Space.SharedEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.SharedEffective),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "shared_effective",
			)
		}
		if v.Space.SnapshotsEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.SnapshotsEffective),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "snapshots_effective",
			)
		}
		if v.Space.UniqueEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.UniqueEffective),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "unique_effective",
			)
		}
		if v.Space.TotalEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*v.Space.TotalEffective),
				purenaa+v.Serial, v.Name, v.Pod.Name, v.VolumeGroup.Name, "total_effective",
			)
		}
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
