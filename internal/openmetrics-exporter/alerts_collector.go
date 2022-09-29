package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type AlertsCollector struct {
	AlertsDesc *prometheus.Desc
	Client     *client.FAClient
}

func (c *AlertsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *AlertsCollector) Collect(ch chan<- prometheus.Metric) {
	alerts := c.Client.GetAlerts("state='open'")
	if len(alerts.Items) == 0 {
		return
	}
	for _, alert := range alerts.Items {
		ch <- prometheus.MustNewConstMetric(
			c.AlertsDesc,
			prometheus.GaugeValue,
			1.0,
			alert.Severity, alert.ComponentType, alert.ComponentName,
		)
	}
}

func NewAlertsCollector(fa *client.FAClient) *AlertsCollector {
	return &AlertsCollector{
		AlertsDesc: prometheus.NewDesc(
			"purefa_alerts_open",
			"FlashArray open alert events",
			[]string{"severity", "component_type", "component_name"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
