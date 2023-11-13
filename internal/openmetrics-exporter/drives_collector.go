package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type DriveCollector struct {
	CapacityDesc     *prometheus.Desc
	Client		 *client.FAClient
}

func (c *DriveCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *DriveCollector) Collect(ch chan<- prometheus.Metric) {
	drl := c.Client.GetDrives()
	if len(drl.Items) == 0 {
		return
	}
	for _, d := range drl.Items {
		ch <- prometheus.MustNewConstMetric(
			c.CapacityDesc,
			prometheus.GaugeValue,
			d.Capacity,
			d.Name, d.Type, d.Status, d.Protocol,
		)
		}
}

func NewDriveCollector(fa *client.FAClient) *DriveCollector {
	return &DriveCollector{
		CapacityDesc: prometheus.NewDesc(
			"purefa_drive_capacity_bytes",
			"FlashArray drive capacity in bytes",
			[]string{"component_name", "component_type", "component_status", "component_protocol"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
