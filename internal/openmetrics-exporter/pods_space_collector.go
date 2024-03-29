package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type PodsSpaceCollector struct {
	ReductionDesc *prometheus.Desc
	SpaceDesc     *prometheus.Desc
	MediatorDesc  *prometheus.Desc
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
		var s float64
		if h.Space.DataReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.ReductionDesc,
				prometheus.GaugeValue,
				*h.Space.DataReduction,
				h.Name,
			)
		}
		if h.Space.Shared != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.Shared),
				h.Name, "shared",
			)
		}
		if h.Space.Snapshots != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.Snapshots),
				h.Name, "snapshots",
			)
		}
		if h.Space.System != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.System),
				h.Name, "system",
			)
		}
		if h.Space.ThinProvisioning != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				*h.Space.ThinProvisioning,
				h.Name, "thin_provisioning",
			)
		}
		if h.Space.TotalPhysical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.TotalPhysical),
				h.Name, "total_physical",
			)
		}
		if h.Space.TotalProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.TotalProvisioned),
				h.Name, "total_provisioned",
			)
		}
		if h.Space.TotalReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				*h.Space.TotalReduction,
				h.Name, "total_reduction",
			)
		}
		if h.Space.Unique != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.Unique),
				h.Name, "unique",
			)
		}
		if h.Space.Virtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.Virtual),
				h.Name, "virtual",
			)
		}
		if h.Space.Replication != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.Replication),
				h.Name, "replication",
			)
		}
		if h.Space.SharedEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.SharedEffective),
				h.Name, "shared_effective",
			)
		}
		if h.Space.SnapshotsEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.SnapshotsEffective),
				h.Name, "snapshots_effective",
			)
		}
		if h.Space.UniqueEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.UniqueEffective),
				h.Name, "unique_effective",
			)
		}
		if h.Space.TotalEffective != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*h.Space.TotalEffective),
				h.Name, "total_effective",
			)
		}
		for _, a := range h.Arrays {
			if a.MediatorStatus == "online" {
				s = 1
			} else {
				s = 0
			}
			ch <- prometheus.MustNewConstMetric(
				c.MediatorDesc,
				prometheus.GaugeValue,
				s,
				a.Name, h.Mediator, h.Name, a.MediatorStatus,
			)
		}
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
		MediatorDesc: prometheus.NewDesc(
			"purefa_pod_mediator_status",
			"FlashArray pod mediator status",
			[]string{"array", "mediator", "pod", "status"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
