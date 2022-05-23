from prometheus_client.core import GaugeMetricFamily

PURE_NAA = 'naa.624a9370'
performance_latency_kpis = ['queue_usec_per_mirrored_write_op',
                            'queue_usec_per_read_op',
                            'queue_usec_per_write_op',
                            'san_usec_per_mirrored_write_op',
                            'san_usec_per_read_op',
                            'san_usec_per_write_op',
                            'service_usec_per_mirrored_write_op',
                            'service_usec_per_read_op',
                            'service_usec_per_write_op',
                            'usec_per_mirrored_write_op',
                            'usec_per_read_op',
                            'usec_per_write_op']

performance_bandwidth_kpis = ['read_bytes_per_sec',
                              'write_bytes_per_sec',
                              'mirrored_write_bytes_per_sec']
            
performance_iops_kpis = ['reads_per_sec',
                         'writes_per_sec',
                         'mirrored_writes_per_sec']

performance_avg_size_kpis = ['bytes_per_read',
                             'bytes_per_write',
                             'bytes_per_op',
                             'bytes_per_mirrored_write']

class VolumePerformanceMetrics():
    """
    Base class for FlashArray Prometheus volume performance metrics
    """
    def __init__(self, fa_client):
        self.volumes = fa_client.volumes()
        self.latency = GaugeMetricFamily(
                           'purefa_volume_performance_latency',
                           'FlashArray volume IO latency',
                           labels = ['name',
                                     'naaid',
                                     'pod',
                                     'vgroup',
                                     'dimension'],
                           unit='usec')
        self.bandwidth = GaugeMetricFamily(
                             'purefa_volume_performance_bandwidth',
                             'FlashArray volume bandwidth',
                             labels = ['name',
                                       'naaid',
                                       'pod',
                                       'vgroup',
                                       'dimension'],
                             unit='bytes')
        self.iops = GaugeMetricFamily(
                        'purefa_volume_performance',
                        'FlashArray volume IOPS',
                        labels = ['name',
                                  'naaid',
                                  'pod',
                                  'vgroup',
                                  'dimension'],
                        unit='iops')
        self.avg_bsz = GaugeMetricFamily(
                           'purefa_volume_performance_avg_block',
                           'FlashArray volume avg block size',
                           labels = ['name',
                                     'naaid',
                                     'pod',
                                     'vgroup',
                                     'dimension'],
                           unit='bytes')

    def _build_metrics(self):
        cnt_l = 0
        cnt_b = 0
        cnt_i = 0
        cnt_a = 0
        for v in self.volumes:
            vol = v['volume']
            pod = ''
            vg = ''
            if hasattr(vol.pod, 'name'):
                pod = vol.pod.name
            if hasattr(vol.volume_group, 'name'):
                vg = vol.volume_group.name
            naaid = PURE_NAA + vol.serial
            perf = v['performance']
            for k in performance_latency_kpis:
                v = getattr(perf, k)
                if v is None:
                    continue
                cnt_l += 1
                self.latency.add_metric([vol.name, naaid, pod, vg, k], v)
            for k in performance_bandwidth_kpis:
                v = getattr(perf, k)
                if v is None:
                    continue
                cnt_b += 1
                self.bandwidth.add_metric([vol.name, naaid, pod, vg, k], v)
            for k in performance_iops_kpis:
                v = getattr(perf, k)
                if v is None:
                    continue
                cnt_i += 1
                self.iops.add_metric([vol.name, naaid, pod, vg, k], v)
            for k in performance_avg_size_kpis:
                v = getattr(perf, k)
                if v is None:
                    continue
                cnt_a += 1
                self.avg_bsz.add_metric([vol.name, naaid, pod, vg, k], v)
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
        yield self.avg_bsz
