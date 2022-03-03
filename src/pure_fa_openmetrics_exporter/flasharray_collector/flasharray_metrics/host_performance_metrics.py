from prometheus_client.core import GaugeMetricFamily

class HostPerformanceMetrics():
    """
    Base class for FlashArray Prometheus host performance metrics
    """

    def __init__(self, fa_client):
        self.latency = None
        self.bandwidth = None
        self.iops = None
        self.hosts = fa_client.hosts()

    def _performance(self):
        self.latency = GaugeMetricFamily(
                           'purefa_host_performance_latency_usec',
                           'FlashArray host IO latency',
                           labels=['name', 'hostgroup', 'dimension'])

        self.bandwidth = GaugeMetricFamily(
                             'purefa_host_performance_bandwidth_bytes',
                             'FlashArray host bandwidth',
                             labels=['name', 'hostgroup', 'dimension'])

        self.iops = GaugeMetricFamily('purefa_host_performance_iops',
                                      'FlashArray host IOPS',
                                      labels=['name', 'hostgroup', 'dimension'])

        for h in self.hosts:
            host = h['host']
            if not host.is_local:
                continue
            hg = ''
            if hasattr(host.host_group, 'name'):
                hg = host.host_group.name
            perf = h['performance']
            self.latency.add_metric([host.name, hg, 'queue_usec_per_read_op'],
                      perf.queue_usec_per_read_op or 0)
            self.latency.add_metric([host.name, hg, 'queue_usec_per_write_op'],
                      perf.queue_usec_per_write_op or 0)
            self.latency.add_metric([host.name, hg, 'queue_usec_per_mirrored_write_op'],
                      perf.queue_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([host.name, hg, 'san_usec_per_read_op'],
                      perf.san_usec_per_read_op or 0)
            self.latency.add_metric([host.name, hg, 'san_usec_per_write_op'],
                      perf.san_usec_per_write_op or 0)
            self.latency.add_metric([host.name, hg, 'san_usec_per_mirrored_write_op'],
                      perf.san_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([host.name, hg, 'service_usec_per_mirrored_write_op'],
                      perf.service_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([host.name, hg, 'service_usec_per_read_op'],
                      perf.service_usec_per_read_op or 0)
            self.latency.add_metric([host.name, hg, 'service_usec_per_write_op'],
                      perf.service_usec_per_write_op or 0)
            self.latency.add_metric([host.name, hg, 'usec_per_read_op'],
                      perf.usec_per_read_op or 0)
            self.latency.add_metric([host.name, hg, 'usec_per_write_op'],
                      perf.usec_per_write_op or 0)
            self.latency.add_metric([host.name, hg, 'usec_per_mirrored_write_op'],
                      perf.usec_per_mirrored_write_op or 0)

            self.bandwidth.add_metric([host.name, hg, 'read_bytes_per_sec'],
                      perf.read_bytes_per_sec or 0)
            self.bandwidth.add_metric([host.name, hg, 'write_bytes_per_sec'],
                      perf.write_bytes_per_sec or 0)
            self.bandwidth.add_metric([host.name, hg, 'mirrored_write_bytes_per_sec'],
                      perf.mirrored_write_bytes_per_sec or 0)

            self.iops.add_metric([host.name, hg, 'reads_per_sec'],
                      perf.reads_per_sec or 0)
            self.iops.add_metric([host.name, hg, 'writes_per_sec'],
                      perf.writes_per_sec or 0)
            self.iops.add_metric([host.name, hg, 'mirrored_writes_per_sec'],
                      perf.mirrored_writes_per_sec or 0)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.bandwidth
        yield self.iops
