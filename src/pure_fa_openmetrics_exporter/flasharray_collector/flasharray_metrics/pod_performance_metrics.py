from prometheus_client.core import GaugeMetricFamily

class PodPerformanceMetrics():
    """
    Base class for FlashArray Prometheus pod performance metrics
    """

    def __init__(self, fa_client):
        self.latency = None
        self.bandwidth = None
        self.iops = None
        self.avg_bsz = None
        self.pods = fa_client.pods()

    def _performance(self):
        self.latency = GaugeMetricFamily('purefa_pod_performance_latency_usec',
                                         'FlashArray pod IO latency',
                                         labels=['name', 'dimension'])

        self.bandwidth = GaugeMetricFamily('purefa_pod_performance_bandwidth_bytes',
                                           'FlashArray pod bandwidth',
                                           labels=['name', 'dimension'])

        self.iops = GaugeMetricFamily('purefa_pod_performance_iops',
                                      'FlashArray pod IOPS',
                                      labels=['name', 'dimension'])

        self.avg_bsz = GaugeMetricFamily(
                           'purefa_pod_performance_avg_block_bytes',
                           'FlashArray avg block size',
                           labels=['name', 'dimension'])

        for p in self.pods:
            pod = p['pod']
            perf = p['performance']
            self.latency.add_metric([pod.name,
                                    'queue_usec_per_mirrored_write_op'],
                                    perf.queue_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([pod.name,
                                    'queue_usec_per_read_op'],
                                    perf.queue_usec_per_read_op or 0)
            self.latency.add_metric([pod.name,
                                    'queue_usec_per_write_op'],
                                    perf.queue_usec_per_write_op or 0)
            self.latency.add_metric([pod.name,
                                    'san_usec_per_mirrored_write_op'],
                                    perf.san_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([pod.name,
                                    'san_usec_per_read_op'],
                                    perf.san_usec_per_read_op or 0)
            self.latency.add_metric([pod.name,
                                    'san_usec_per_write_op'],
                                    perf.san_usec_per_write_op or 0)
            self.latency.add_metric([pod.name,
                                    'service_usec_per_mirrored_write_op'],
                                    perf.service_usec_per_mirrored_write_op or 0) 
            self.latency.add_metric([pod.name,
                                    'service_usec_per_read_op'],
                                    perf.service_usec_per_read_op or 0) 
            self.latency.add_metric([pod.name,
                                    'service_usec_per_write_op'],
                                    perf.service_usec_per_write_op or 0) 
            self.latency.add_metric([pod.name,
                                    'usec_per_mirrored_write_op'], 
                                    perf.usec_per_mirrored_write_op or 0) 
            self.latency.add_metric([pod.name,
                                    'usec_per_read_op'],
                                    perf.usec_per_read_op or 0) 
            self.latency.add_metric([pod.name,
                                    'usec_per_write_op'],
                                    perf.usec_per_write_op or 0)

            self.bandwidth.add_metric([pod.name,
                                      'read_bytes_per_sec'],
                                      perf.read_bytes_per_sec or 0)
            self.bandwidth.add_metric([pod.name,
                                      'write_bytes_per_sec'],
                                      perf.write_bytes_per_sec or 0)
            self.bandwidth.add_metric([pod.name,
                                      'mirrored_write_bytes_per_sec'],
                                      perf.mirrored_write_bytes_per_sec or 0)

            self.iops.add_metric([pod.name,
                                 'reads_per_sec'],
                                 perf.reads_per_sec or 0)
            self.iops.add_metric([pod.name,
                                 'writes_per_sec'],
                                 perf.writes_per_sec or 0)
            self.iops.add_metric([pod.name,
                                 'mirrored_writes_per_sec'],
                                 perf.mirrored_writes_per_sec or 0)

            self.avg_bsz.add_metric([pod.name,
                                    'bytes_per_read'],
                                    perf.bytes_per_read or 0)
            self.avg_bsz.add_metric([pod.name,
                                    'bytes_per_write'],
                                    perf.bytes_per_write or 0)
            self.avg_bsz.add_metric([pod.name,
                                    'bytes_per_op'],
                                    perf.bytes_per_op or 0)
            self.avg_bsz.add_metric([pod.name,
                                    'bytes_per_mirrored_write'],
                                    perf.bytes_per_mirrored_write or 0)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
