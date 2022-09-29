package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
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
	array := arrays.Items[0]

	ch <- prometheus.MustNewConstMetric(
		c.ReductionDesc,
		prometheus.GaugeValue,
		array.Space.DataReduction,
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.Shared,
		"shared",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.Snapshots,
		"snapshots",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.System,
		"system",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.ThinProvisioning,
		"thin_provisioning",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.TotalPhysical,
		"total_physical",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.TotalProvisioned,
		"total_provisioned",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.Unique,
		"unique",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.Virtual,
		"virtual",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.Replication,
		"replication",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.SharedEffective,
		"shared_effective",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.SnapshotsEffective,
		"snapshots_effective",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.UniqueEffective,
		"unique_effective",
	)
	ch <- prometheus.MustNewConstMetric(
		c.SpaceDesc,
		prometheus.GaugeValue,
		array.Space.TotalEffective,
		"total_effective",
	)
}

func NewArraysSpaceCollector(fa *client.FAClient) *ArraySpaceCollector {
	return &ArraySpaceCollector{
                ReductionDesc: prometheus.NewDesc(
                        "purefa_array_space_data_reduction_ratio",
                        "FlashArray space data reduction",
                        []string{},
                        prometheus.Labels{},
                ),
                SpaceDesc: prometheus.NewDesc(
                        "purefa_array_space_bytes",
                        "FlashArray space in bytes",
                        []string{"space"},
                        prometheus.Labels{},
                ),
		Client: fa,
	}
}
