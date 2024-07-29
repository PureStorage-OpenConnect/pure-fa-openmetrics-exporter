package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type ArraySpaceCollector struct {
	ReductionDesc   *prometheus.Desc
	SpaceDesc       *prometheus.Desc
	UtilizationDesc *prometheus.Desc
	Client          *client.FAClient
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

	if a.Space.DataReduction != nil {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			*a.Space.DataReduction,
		)
	}
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		a.Capacity, "capacity",
	)
	if a.Space.Shared != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.Shared), "shared",
		)
	}
	if a.Space.Snapshots != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.Snapshots), "snapshots",
		)
	}
	if a.Space.System != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.System), "system",
		)
	}
	if a.Space.ThinProvisioning != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			*a.Space.ThinProvisioning, "thin_provisioning",
		)
	}
	if a.Space.TotalPhysical != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.TotalPhysical), "total_physical",
		)
	}
	if a.Space.TotalProvisioned != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.TotalProvisioned), "total_provisioned",
		)
	}
	if a.Space.TotalReduction != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			*a.Space.TotalReduction, "total_reduction",
		)
	}
	if a.Space.Unique != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.Unique), "unique",
		)
	}
	if a.Space.Virtual != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.Virtual), "virtual",
		)
	}
	if a.Space.Replication != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.Replication), "replication",
		)
	}
	if a.Space.SharedEffective != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.SharedEffective), "shared_effective",
		)
	}
	if a.Space.SnapshotsEffective != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.SnapshotsEffective), "snapshots_effective",
		)
	}
	if a.Space.UniqueEffective != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.UniqueEffective), "unique_effective",
		)
	}
	if a.Space.TotalEffective != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(*a.Space.TotalEffective), "total_effective",
		)
	}
	if a.Space.System != nil {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			a.Capacity-(float64(*a.Space.System)+float64(*a.Space.Replication)+float64(*a.Space.Shared)+float64(*a.Space.Snapshots)+float64(*a.Space.Unique)), "empty",
		)
		ch <- prometheus.MustNewConstMetric(
			c.UtilizationDesc,
			prometheus.GaugeValue,
			(float64(*a.Space.System)+float64(*a.Space.Replication)+float64(*a.Space.Shared)+float64(*a.Space.Snapshots)+float64(*a.Space.Unique))/a.Capacity*100,
		)
	} else {
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			a.Capacity-(float64(*a.Space.Replication)+float64(*a.Space.Shared)+float64(*a.Space.Snapshots)+float64(*a.Space.Unique)), "empty",
		)
		ch <- prometheus.MustNewConstMetric(
			c.UtilizationDesc,
			prometheus.GaugeValue,
			(float64(*a.Space.Replication)+float64(*a.Space.Shared)+float64(*a.Space.Snapshots)+float64(*a.Space.Unique))/a.Capacity*100,
		)
	}
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
		UtilizationDesc: prometheus.NewDesc(
			"purefa_array_space_utilization",
			"FlashArray array space utilization in percent",
			[]string{},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
