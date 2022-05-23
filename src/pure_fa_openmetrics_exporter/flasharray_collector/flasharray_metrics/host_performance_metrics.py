from prometheus_client.core import GaugeMetricFamily

performance_latency_kpis = ['queue_usec_per_read_op',
                            'queue_usec_per_write_op',
                            'queue_usec_per_mirrored_write_op',
                            'san_usec_per_read_op',
                            'san_usec_per_write_op',
                            'san_usec_per_mirrored_write_op',
                            'service_usec_per_mirrored_write_op',
                            'service_usec_per_read_op',
                            'service_usec_per_write_op',
                            'usec_per_read_op',
                            'usec_per_write_op',
                            'usec_per_mirrored_write_op']

performance_bandwidth_kpis = ['read_bytes_per_sec',
                              'write_bytes_per_sec',
                              'mirrored_write_bytes_per_sec']

performance_iops_kpis = ['reads_per_sec',
                         'writes_per_sec',
                         'mirrored_writes_per_sec']

class HostPerformanceMetrics():
    """
    Base class for FlashArray Prometheus host performance metrics
    """

    def __init__(self, fa_client):
        self.hosts = fa_client.hosts()
        self.latency = GaugeMetricFamily(
                            'purefa_host_performance_latency',
                            'FlashArray host IO latency',
                            labels=['name', 'hostgroup', 'dimension'],
                            unit='usec')
        self.bandwidth = GaugeMetricFamily(
                              'purefa_host_performance_bandwidth',
                              'FlashArray host bandwidth',
                              labels=['name', 'hostgroup', 'dimension'],
                              unit='bytes')
        self.iops = GaugeMetricFamily('purefa_host_performance',
                         'FlashArray host IOPS',
                         labels=['name', 'hostgroup', 'dimension'],
                         unit='iops')

    def _performance(self):
        cnt_l = 0
        cnt_b = 0
        cnt_i = 0
        for h in self.hosts:
            host = h['host']
            if not host.is_local:
                continue
            hg = ''
            if hasattr(host.host_group, 'name'):
                hg = host.host_group.name
            perf = h['performance']
            for k in performance_latency_kpis:
                v = getattr(perf, k)
                if v is None:
                    continue
                cnt_l += 1
                self.latency.add_metric([host.name, hg, k], v)
            for k in performance_bandwidth_kpis:
                v = getattr(perf, k)
                if v is None:
                    continue
                cnt_b += 1
                self.bandwidth.add_metric([host.name, hg, k], v)
            for k in performance_iops_kpis:
                v = getattr(perf, k)
                if v is None:
                    continue
                cnt_i += 1
                self.iops.add_metric([host.name, hg, k], v)
        if cnt_l == 0:
            self.latecny = None
        if cnt_b == 0:
            self.bandwidth = None
        if cnt_i == 0:
            self.iops = None

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.bandwidth
        yield self.iops
