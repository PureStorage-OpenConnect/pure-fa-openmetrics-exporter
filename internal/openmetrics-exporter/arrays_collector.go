package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
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

	subscriptions := c.Client.GetSubscriptions()
	if len(subscriptions.Items) == 0 {
		return
	}
	subscription := subscriptions.Items[0]

	ch <- prometheus.MustNewConstMetric(
		c.ArraysDesc,
		prometheus.GaugeValue,
		1.0,
		array.Name, array.Os, subscription.Service, array.Id, array.Version,
	)
}

func NewArraysCollector(fa *client.FAClient) *ArraysCollector {
	return &ArraysCollector{
		ArraysDesc: prometheus.NewDesc(
			"purefa_info",
			"FlashArray system information",
			[]string{"array_name", "os", "subscription_type", "system_id", "version"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
