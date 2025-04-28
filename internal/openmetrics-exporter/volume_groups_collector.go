package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type VolumeGroupsCollector struct {
	QoSBandwidthLimitDesc *prometheus.Desc
	QoSIPOSLimitDesc      *prometheus.Desc
	Client                *client.FAClient
}

func (c *VolumeGroupsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *VolumeGroupsCollector) Collect(ch chan<- prometheus.Metric) {
	volumeGroups := c.Client.GetVolumeGroups()
	if len(volumeGroups.Items) == 0 {
		return
	}
	for _, vg := range volumeGroups.Items {
		if vg.QoS.BandwidthLimit != nil {
			ch <- prometheus.MustNewConstMetric(
				c.QoSBandwidthLimitDesc,
				prometheus.GaugeValue,
				float64(*vg.QoS.BandwidthLimit),
				vg.Name,
			)
		}
		if vg.QoS.IopsLimit != nil {
			ch <- prometheus.MustNewConstMetric(
				c.QoSIPOSLimitDesc,
				prometheus.GaugeValue,
				float64(*vg.QoS.IopsLimit),
				vg.Name,
			)
		}
	}
}

func NewVolumeGroupsCollector(fa *client.FAClient) *VolumeGroupsCollector {
	return &VolumeGroupsCollector{
		QoSBandwidthLimitDesc: prometheus.NewDesc(
			"purefa_volume_group_qos_bandwidth_bytes_per_sec_limit",
			"FlashArray volume group maximum QoS bandwidth limit in bytes per second",
			[]string{"name"},
			prometheus.Labels{},
		),
		QoSIPOSLimitDesc: prometheus.NewDesc(
			"purefa_volume_group_qos_iops_limit",
			"FlashArray volume group QoS IOPs limit",
			[]string{"name"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
