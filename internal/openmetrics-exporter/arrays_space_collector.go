package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type ArraySpaceCollector struct {
	ReductionDesc *prometheus.Desc
	SpaceDesc     *prometheus.Desc
	Client        *client.FAClient
}

func (c *ArraySpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *ArraySpaceCollector) Collect(ch chan<- prometheus.Metric) {
	arrays := c.Client.GetArrays()
	if len(arrays.Items) == 0 {
		return
	}
	a := arrays.Items[0]
	ch <- prometheus.MustNewConstMetric(
		c.ReductionDesc,
		prometheus.GaugeValue,
		a.Space.DataReduction,
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Capacity, "capacity",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.Shared, "shared",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.Snapshots, "snapshots",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.System, "system",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.ThinProvisioning, "thin_provisioning",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.TotalPhysical, "total_physical",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.TotalProvisioned, "total_provisioned",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.TotalReduction, "total_reduction",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.Unique, "unique",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.Virtual, "virtual",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.Replication, "replication",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.SharedEffective, "shared_effective",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.SnapshotsEffective, "snapshots_effective",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.UniqueEffective, "unique_effective",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Space.TotalEffective, "total_effective",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Capacity-a.Space.System-a.Space.Replication-a.Space.Shared-a.Space.Snapshots-a.Space.Unique, "empty",
	)
}

func NewArraySpaceCollector(fa *client.FAClient) *ArraySpaceCollector {
	return &ArraySpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefa_array_space_data_reduction_ratio",
			"FlashArray array space data reduction",
			[]string{},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefa_array_space_bytes",
			"FlashArray array space in bytes",
			[]string{"space"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
