package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type HostsSpaceCollector struct {
	ReductionDesc    *prometheus.Desc
	SpaceDesc        *prometheus.Desc
	ConnectivityDesc *prometheus.Desc
	Client           *client.FAClient
}

func (c *HostsSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *HostsSpaceCollector) Collect(ch chan<- prometheus.Metric) {
	hosts := c.Client.GetHosts()
	if len(hosts.Items) == 0 {
		return
	}
	for _, h := range hosts.Items {
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
		ch <- prometheus.MustNewConstMetric(
			c.ConnectivityDesc,
			prometheus.GaugeValue,
			1.0,
			h.Name, h.PortConnectivity.Details, h.PortConnectivity.Status,
		)
	}
}

func NewHostsSpaceCollector(fa *client.FAClient) *HostsSpaceCollector {
	return &HostsSpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefa_host_space_data_reduction_ratio",
			"FlashArray host space data reduction",
			[]string{"host"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefa_host_space_bytes",
			"FlashArray host space in bytes",
			[]string{"host", "space"},
			prometheus.Labels{},
		),
		ConnectivityDesc: prometheus.NewDesc(
			"purefa_host_connectivity_info",
			"FlashArray host connectivity information",
			[]string{"host", "details", "status"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
