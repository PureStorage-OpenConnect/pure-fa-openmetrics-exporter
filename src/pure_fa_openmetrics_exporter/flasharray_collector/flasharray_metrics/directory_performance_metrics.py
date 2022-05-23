from prometheus_client.core import GaugeMetricFamily


performance_latency_kpis = ['usec_per_read_op',
                            'usec_per_write_op',
                            'usec_per_other_op']

performance_bandwidth_kpis = ['write_bytes_per_sec',
                              'read_bytes_per_sec']

performance_iops_kpis = ['reads_per_sec',
                         'writes_per_sec',
                         'others_per_sec']

performance_avg_size_kpis = ['bytes_per_op',
                             'bytes_per_read',
                             'bytes_per_write']

class DirectoryPerformanceMetrics():
    """
    Base class for FlashArray Prometheus directory performance metrics
    """

    def __init__(self, fa_client):
        self.directories = fa_client.directories()
        self.latency = GaugeMetricFamily(
                           'purefa_directory_performance_latency',
                           'FlashArray directory IO latency',
                           labels=['name', 'filesystem', 'path', 'dimension'],
                           unit='usec')
        self.bandwidth = GaugeMetricFamily(
                             'purefa_directory_performance_bandwidth',
                             'FlashArray directory bandwidth',
                             labels=['name', 'filesystem', 'path', 'dimension'],
                              unit='bytes')
        self.iops = GaugeMetricFamily(
                        'purefa_directory_performance_iops',
                        'FlashArray directory IOPS',
                        labels=['name', 'filesystem', 'path', 'dimension'])
        self.avg_size = GaugeMetricFamily(
                            'purefa_directory_performance_avg_block',
                            'FlashArray directory avg block size',
                            labels=['name', 'filesystem', 'path', 'dimension'],
                            unit='bytes')

    def _build_metrics(self):
        cnt_l = 0
        cnt_b = 0
        cnt_i = 0
        cnt_a = 0
        for dir in self.directories:
            d = dir['directory']
            p = dir['performance']
            for k in performance_latency_kpis:
                v = getattr(p, k)
                if v is None:
                    continue
                cnt_l += 1
                self.latency.add_metric([d.directory_name,
                                         d.file_system.name,
                                         d.path,
                                         k], v)
            for k in performance_bandwidth_kpis:
                v = getattr(p, k)
                if v is None:
                    continue
                cnt_b += 1
                self.bandwidth.add_metric([d.directory_name,
                                           d.file_system.name,
                                           d.path,
                                           k], v)
            for k in performance_iops_kpis:
                v = getattr(p, k)
                if v is None:
                    continue
                cnt_i += 1
                self.iops.add_metric([d.directory_name,
                                      d.file_system.name,
                                      d.path,
                                      k], v)
            for k in performance_avg_size_kpis:
                v = getattr(p, k)
                if v is None:
                    continue
                cnt_a += 1
                self.avg_size.add_metric([d.directory_name,
                                          d.file_system.name,
                                          d.path,
                                          k], v)
        if cnt_l == 0 :
            self.latency = None
        if cnt_b == 0 :
            self.bandwidth = None
        if cnt_i == 0 :
            self.iops = None
        if cnt_a == 0:
            self.avg_size

    def get_metrics(self):
        self._build_metrics()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_size
