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
