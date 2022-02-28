from prometheus_client.core import GaugeMetricFamily


class PodPerformanceMetrics():
    """
    Base class for FlashArray Prometheus pod performance metrics
    """

    def __init__(self, fa):
        self.fa = fa

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

    def _latency(self) -> None:
        """
        Create pods latency metrics of gauge type, with pod name and
        dimension as label.
        """
        for p in self.fa.get_pods():
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
                val = p['performance'][k]
                val = val if val is not None else 0
                self.latency.add_metric([p['name'], k], val)

    def _bandwidth(self) -> None:
        """
        Create pods bandwidth metrics of gauge type, with pod name and
        dimension as label.
        """
        for p in self.fa.get_pods():
            for k in ['read_bytes_per_sec',
                      'write_bytes_per_sec',
                      'mirrored_write_bytes_per_sec']:
                val = p['performance'][k]
                val = val if val is not None else 0
                self.bandwidth.add_metric([p['name'], k], val)

    def _iops(self) -> None:
        """
        Create IOPS bandwidth metrics of gauge type, with pod name and
        dimension as label.
        """
        for p in self.fa.get_pods():
            for k in ['reads_per_sec',
                      'writes_per_sec',
                      'mirrored_writes_per_sec']:
                val = p['performance'][k]
                val = val if val is not None else 0
                self.iops.add_metric([p['name'], k], val)

    def _avg_block_size(self) -> None:
        """
        Create array average block size performance metrics of gauge type.
        Metrics values can be iterated over.
        """
        for p in self.fa.get_pods():
            for k in ['bytes_per_read',
                      'bytes_per_write',
                      'bytes_per_op',
                      'bytes_per_mirrored_write']:
                val = p['performance'][k]
                val = val if val is not None else 0
                self.avg_bsz.add_metric([p['name'], k], val)

    def get_metrics(self) -> None:
        self._latency()
        self._bandwidth()
        self._iops()
        self._avg_block_size()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
