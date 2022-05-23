from prometheus_client.core import GaugeMetricFamily

protocols = ['all', 'nfs', 'smb']

performance_latency_kpis = ['local_queue_usec_per_op',
                            'queue_usec_per_read_op',
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
                            'usec_per_mirrored_write_op',
                            'usec_per_other_op']

performance_bandwidth_kpis = ['read_bytes_per_sec',
                              'write_bytes_per_sec',
                              'mirrored_write_bytes_per_sec']

performance_iops_kpis = ['reads_per_sec',
                         'writes_per_sec',
                         'mirrored_writes_per_sec',
                         'others_per_sec']

performance_avg_size_kpis = ['bytes_per_read',
                             'bytes_per_write',
                             'bytes_per_op']

class ArrayPerformanceMetrics():
    """
    Base class for FlashArray OpenMetrics array performance metrics
    """

    def __init__(self, fa_client):
        self.array = fa_client.arrays()[0]
        self.latency = GaugeMetricFamily(
                            'purefa_array_performance_latency',
                            'FlashArray array latency',
                            labels=['protocol', 'dimension'],
                            unit='usec')
        self.bandwidth = GaugeMetricFamily(
                              'purefa_array_performance_bandwidth',
                              'FlashArray bandwidth',
                              labels=['protocol', 'dimension'],
                                  unit='bytes')
        self.iops = GaugeMetricFamily(
                         'purefa_array_performance_iops',
                         'FlashArray IOPS',
                         labels=['protocol', 'dimension'])
        self.avg_size = GaugeMetricFamily(
                             'purefa_array_performance_average_block',
                             'FlashArray array average block size',
                             labels=['protocol', 'dimension'],
                             unit='bytes')

    def _build_metrics(self):
        cnt_l = 0
        cnt_b = 0
        cnt_i = 0
        cnt_a = 0
        for p in protocols:
            for k in performance_latency_kpis:
                v = getattr(self.array['performance'][p], k)
                if v is None:
                    continue
                cnt_l += 1
                self.latency.add_metric([p, k], v)
            for k in performance_bandwidth_kpis:
                v = getattr(self.array['performance'][p], k)
                if v is None:
                    continue
                cnt_b += 1
                self.bandwidth.add_metric([p, k], v)
            for k in performance_iops_kpis:
                v = getattr(self.array['performance'][p], k)
                if v is None:
                    continue
                cnt_i += 1
                self.iops.add_metric([p, k], v)
            for k in performance_avg_size_kpis:
                v = getattr(self.array['performance'][p], k)
                if v is None:
                    continue
                cnt_a += 1
                self.avg_size.add_metric([p, k], v)
        if cnt_l == 0:
            self.latency = None
        if cnt_b == 0:
            self.bandwidth = None
        if cnt_i == 0:
            self.iops = None
        if cnt_a == 0:
            self.avg_size = None

    def get_metrics(self):
        self._build_metrics()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_size
