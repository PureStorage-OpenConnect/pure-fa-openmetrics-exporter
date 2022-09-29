package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type ArraysCollector struct {
	ArraysDesc *prometheus.Desc
	Client     *client.FAClient
}

func (c *ArraysCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.ArraysDesc
}

func (c *ArraysCollector) Collect(ch chan<- prometheus.Metric) {
	arrays := c.Client.GetArrays()
	if len(arrays.Items) == 0 {
		return
	}
	array := arrays.Items[0]

	ch <- prometheus.MustNewConstMetric(
		c.ArraysDesc,
		prometheus.GaugeValue,
		1.0,
		array.Name, array.Id, array.Os, array.Version,
	)
}

func NewArraysCollector(fa *client.FAClient) *ArraysCollector {
	return &ArraysCollector{
		ArraysDesc: prometheus.NewDesc(
			"purefa_info",
			"FlashArray system information",
			[]string{"array_name", "system_id", "os", "version"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
