package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type PodsSpaceCollector struct {
	ReductionDesc *prometheus.Desc
	SpaceDesc     *prometheus.Desc
	Client        *client.FAClient
}

func (c *PodsSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *PodsSpaceCollector) Collect(ch chan<- prometheus.Metric) {
	pods := c.Client.GetPods()
	if len(pods.Items) == 0 {
		return
	}
	for _, h := range pods.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			h.Space.DataReduction,
			h.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.Shared,
			h.Name, "shared",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.Snapshots,
			h.Name, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.System,
			h.Name, "system",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.ThinProvisioning,
			h.Name, "thin_provisioning",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.TotalPhysical,
			h.Name, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.TotalProvisioned,
			h.Name, "total_provisioned",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.TotalReduction,
			h.Name, "total_reduction",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.Unique,
			h.Name, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.Virtual,
			h.Name, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.Replication,
			h.Name, "replication",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.SharedEffective,
			h.Name, "shared_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.SnapshotsEffective,
			h.Name, "snapshots_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.UniqueEffective,
			h.Name, "unique_effective",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			h.Space.TotalEffective,
			h.Name, "total_effective",
		)
        }
}

func NewPodsSpaceCollector(fa *client.FAClient) *PodsSpaceCollector {
	return &PodsSpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefa_pod_space_data_reduction_ratio",
			"FlashArray pod space data reduction",
			[]string{"name"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefa_pod_space_bytes",
			"FlashArray pod space in bytes",
			[]string{"name", "space"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
