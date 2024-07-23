package collectors

import (
	client "purestorage/fa-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type PodReplicaLinksLagCollector struct {
	AvgLagDesc *prometheus.Desc
	MaxLagDesc *prometheus.Desc
	Client     *client.FAClient
}

func (c *PodReplicaLinksLagCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *PodReplicaLinksLagCollector) Collect(ch chan<- prometheus.Metric) {
	podrl := c.Client.GetPodReplicaLinksLag()
	if len(podrl.Items) == 0 {
		return
	}
	for _, p := range podrl.Items {
		ch <- prometheus.MustNewConstMetric(
			c.AvgLagDesc,
			prometheus.GaugeValue,
			p.Lag.Avg,
			p.Remotes[0].Name, p.LocalPod.Name, p.RemotePod.Name, p.Direction, p.Status,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MaxLagDesc,
			prometheus.GaugeValue,
			p.Lag.Max,
			p.Remotes[0].Name, p.LocalPod.Name, p.RemotePod.Name, p.Direction, p.Status,
		)
	}
}

func NewPodReplicaLinksLagCollector(fa *client.FAClient) *PodReplicaLinksLagCollector {
	return &PodReplicaLinksLagCollector{
		AvgLagDesc: prometheus.NewDesc(
			"purefa_pod_replica_links_lag_average_msec",
			"FlashArray pod links average lag in milliseconds",
			[]string{"remote", "local_pod", "remote_pod", "direction", "status"},
			prometheus.Labels{},
		),
		MaxLagDesc: prometheus.NewDesc(
			"purefa_pod_replica_links_lag_max_msec",
			"FlashArray pod links max lag in milliseconds",
			[]string{"remote", "local_pod", "remote_pod", "direction", "status"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
