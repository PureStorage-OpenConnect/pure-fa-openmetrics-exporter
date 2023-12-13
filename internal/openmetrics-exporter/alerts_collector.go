package collectors

import (
	"fmt"
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
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
	al := make(map[string]float64)
	for _, alert := range alerts.Items {
		al[fmt.Sprintf("%s,%d,%s,%d,%s,%s,%s,%s",
			alert.Category,
			alert.Code,
			alert.ComponentType,
			alert.Created,
			alert.Issue,
			alert.Name,
			alert.Severity,
			alert.Summary,
		)] += 1
	}
	for a, n := range al {
		alert := strings.Split(a, ",")
		ch <- prometheus.MustNewConstMetric(
			c.AlertsDesc,
			prometheus.GaugeValue,
			n,
			alert[0], alert[1], alert[2], alert[3], alert[4], alert[5], alert[6], alert[7],
		)
	}
}

func NewAlertsCollector(fa *client.FAClient) *AlertsCollector {
	return &AlertsCollector{
		AlertsDesc: prometheus.NewDesc(
			"purefa_alerts_open",
			"FlashArray open alert events",
			[]string{"category", "code", "component_type", "created", "issue", "name", "severity", "summary"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
