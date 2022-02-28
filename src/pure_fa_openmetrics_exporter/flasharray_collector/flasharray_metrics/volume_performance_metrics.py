from prometheus_client.core import GaugeMetricFamily


class VolumePerformanceMetrics():
    """
    Base class for FlashArray Prometheus volume performance metrics
    """

    def __init__(self, fa):
        self.fa = fa

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

    def _latency(self) -> None:
        """
        Create volumes latency metrics of gauge type.
        """
        for v in self.fa.get_volumes():
            for k in ['queue_usec_per_mirrored_write_op',
                      'queue_usec_per_read_op',
                      'queue_usec_per_write_op',
                      'san_usec_per_mirrored_write_op',
                      'san_usec_per_read_op',
                      'san_usec_per_write_op',
                      'service_usec_per_mirrored_write_op',
                      'service_usec_per_read_op',
                      'service_usec_per_read_op_cache_reduction',
                      'service_usec_per_write_op',
                      'usec_per_mirrored_write_op',
                      'usec_per_read_op',
                      'usec_per_write_op']:
                pod = v['pod']['name']
                pod = pod  if pod is not None else ''
                vg = v['volume_group']['name']
                vg = vg if vg is not None else ''
                val = v['performance'][k]
                val = val if val is not None else 0
                self.latency.add_metric([v['name'], v['naaid'], pod, vg, k],
                                        val)


    def _bandwidth(self) -> None:
        """
        Create volumes bandwidth metrics of gauge type.
        """
        for v in self.fa.get_volumes():
            for k in ['read_bytes_per_sec',
                      'write_bytes_per_sec',
                      'mirrored_write_bytes_per_sec']:
                pod = v['pod']['name']
                pod = pod  if pod is not None else ''
                vg = v['volume_group']['name']
                vg = vg if vg is not None else ''
                val = v['performance'][k]
                val = val if val is not None else 0
                self.bandwidth.add_metric([v['name'], v['naaid'], pod, vg, k],
                                          val)

    def _iops(self) -> None:
        """
        Create volumes IOPS bandwidth metrics of gauge type.
        """
        for v in self.fa.get_volumes():
            for k in ['reads_per_sec',
                      'writes_per_sec',
                      'mirrored_writes_per_sec']:
                pod = v['pod']['name']
                pod = pod  if pod is not None else ''
                vg = v['volume_group']['name']
                vg = vg if vg is not None else ''
                val = v['performance'][k]
                val = val if val is not None else 0
                self.iops.add_metric([v['name'], v['naaid'], pod, vg, k], val)

    def _avg_block_size(self) -> None:
        """
        Create volumes average block size performance metrics of gauge type.
        Metrics values can be iterated over.
        """
        for v in self.fa.get_volumes():
            for k in ['bytes_per_read',
                      'bytes_per_write',
                      'bytes_per_op',
                      'bytes_per_mirrored_write']:
                pod = v['pod']['name']
                pod = pod  if pod is not None else ''
                vg = v['volume_group']['name']
                vg = vg if vg is not None else ''
                val = v['performance'][k]
                val = val if val is not None else 0
                self.avg_bsz.add_metric([v['name'], v['naaid'], pod, vg, k],
                                        val)

    def get_metrics(self) -> None:
        self._latency()
        self._bandwidth()
        self._iops()
        self._avg_block_size()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
