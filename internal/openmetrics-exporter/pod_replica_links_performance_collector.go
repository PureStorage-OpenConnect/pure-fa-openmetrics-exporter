package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type PodReplicaLinksPerformanceCollector struct {
	BandwidthDesc   *prometheus.Desc
	Client          *client.FAClient
}

func (c *PodReplicaLinksPerformanceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *PodReplicaLinksPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	podrl := c.Client.GetPodReplicaLinksPerformance()
	if len(podrl.Items) == 0 {
		return
	}
	for _, p := range podrl.Items {
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			p.BytesPerSecFromRemote,
			p.Remotes[0].Name,
			p.LocalPod.Name,
			p.RemotePod.Name,
			p.Direction, "bytes_per_sec_from_remote",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			p.BytesPerSecToRemote,
			p.Remotes[0].Name,
			p.LocalPod.Name,
			p.RemotePod.Name,
			p.Direction, "bytes_per_sec_to_remote",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			p.BytesPerSecTotal,
			p.Remotes[0].Name,
			p.LocalPod.Name,
			p.RemotePod.Name,
			p.Direction, "bytes_per_sec_total",
		)
	}
}

func NewPodReplicaLinksPerformanceCollector(fa *client.FAClient) *PodReplicaLinksPerformanceCollector {
	return &PodReplicaLinksPerformanceCollector{
		BandwidthDesc: prometheus.NewDesc(
			"purefa_pod_replica_links_performance_bandwidth_bytes",
			"FlashArray pod links bandwidth",
			[]string{"remote", "local_pod", "remote_pod", "direction", "dimension"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
