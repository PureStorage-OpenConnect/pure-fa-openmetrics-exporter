from prometheus_client.core import GaugeMetricFamily

PURE_NAA = 'naa.624a9370'

class VolumePerformanceMetrics():
    """
    Base class for FlashArray Prometheus volume performance metrics
    """
    def __init__(self, fa_client):
        self.latency = None
        self.bandwidth = None
        self.iops = None
        self.avg_bsz = None
        self.volumes = fa_client.volumes()

    def _performance(self):
        self.latency = GaugeMetricFamily(
                           'purefa_volume_performance_latency_usec',
                           'FlashArray volume IO latency',
                           labels = ['name',
                                     'naaid',
                                     'pod',
                                     'vgroup',
                                     'dimension'])

        self.bandwidth = GaugeMetricFamily(
                             'purefa_volume_performance_bandwidth_bytes',
                             'FlashArray volume bandwidth',
                             labels = ['name',
                                       'naaid',
                                       'pod',
                                       'vgroup',
                                       'dimension'])

        self.iops = GaugeMetricFamily(
                        'purefa_volume_performance_iops',
                        'FlashArray volume IOPS',
                        labels = ['name',
                                  'naaid',
                                  'pod',
                                  'vgroup',
                                  'dimension'])

        self.avg_bsz = GaugeMetricFamily(
                           'purefa_volume_performance_avg_block_bytes',
                           'FlashArray volume avg block size',
                           labels = ['name',
                                     'naaid',
                                     'pod',
                                     'vgroup',
                                     'dimension'])

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
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'queue_usec_per_mirrored_write_op'],
                                    perf.queue_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'queue_usec_per_read_op'],
                                    perf.queue_usec_per_read_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'queue_usec_per_write_op'],
                                    perf.queue_usec_per_write_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'san_usec_per_mirrored_write_op'],
                                    perf.san_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'san_usec_per_read_op'],
                                    perf.san_usec_per_read_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'san_usec_per_write_op'],
                                    perf.san_usec_per_write_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'service_usec_per_mirrored_write_op'],
                                    perf.service_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'service_usec_per_read_op'],
                                    perf.service_usec_per_read_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'service_usec_per_write_op'],
                                    perf.service_usec_per_write_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'usec_per_mirrored_write_op'],
                                    perf.usec_per_mirrored_write_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'usec_per_read_op'],
                                    perf.usec_per_read_op or 0)
            self.latency.add_metric([vol.name, naaid, pod, vg,
                                    'usec_per_write_op'],
                                    perf.usec_per_write_op or 0)

            self.bandwidth.add_metric([vol.name, naaid, pod, vg,
                                      'read_bytes_per_sec'],
                                      perf.read_bytes_per_sec or 0)
            self.bandwidth.add_metric([vol.name, naaid, pod, vg,
                                      'write_bytes_per_sec'],
                                      perf.write_bytes_per_sec or 0)
            self.bandwidth.add_metric([vol.name, naaid, pod, vg,
                                      'mirrored_write_bytes_per_sec'],
                                      perf.mirrored_write_bytes_per_sec or 0)

            self.iops.add_metric([vol.name, naaid, pod, vg,
                                 'reads_per_sec'],
                                 perf.reads_per_sec or 0)
            self.iops.add_metric([vol.name, naaid, pod, vg,
                                 'writes_per_sec'],
                                 perf.writes_per_sec or 0)
            self.iops.add_metric([vol.name, naaid, pod, vg,
                                 'mirrored_writes_per_sec'],
                                 perf.mirrored_writes_per_sec or 0)
            self.avg_bsz.add_metric([vol.name, naaid, pod, vg,
                                    'bytes_per_read'],
                                    perf.bytes_per_read or 0)
            self.avg_bsz.add_metric([vol.name, naaid, pod, vg,
                                    'bytes_per_write'],
                                    perf.bytes_per_write or 0)
            self.avg_bsz.add_metric([vol.name, naaid, pod, vg,
                                    'bytes_per_op'],
                                    perf.bytes_per_op or 0)
            self.avg_bsz.add_metric([vol.name, naaid, pod, vg,
                                    'bytes_per_mirrored_write'],
                                    perf.bytes_per_mirrored_write or 0)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
