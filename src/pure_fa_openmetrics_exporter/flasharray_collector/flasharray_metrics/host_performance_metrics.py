from prometheus_client.core import GaugeMetricFamily

class HostPerformanceMetrics():
    """
    Base class for FlashArray Prometheus host performance metrics
    """

    def __init__(self, fa):
        self.fa = fa
        self.latency = GaugeMetricFamily(
                           'purefa_host_performance_latency_usec',
                           'FlashArray host IO latency',
                           labels=['name', 'dimension'])

        self.bandwidth = GaugeMetricFamily(
                             'purefa_host_performance_bandwidth_bytes',
                             'FlashArray host bandwidth',
                             labels=['name', 'dimension'])

        self.iops = GaugeMetricFamily('purefa_host_performance_iops',
                                      'FlashArray host IOPS',
                                      labels=['name', 'dimension'])

    def _latency(self) -> None:
        """
        Create hosts latency metrics of gauge type.
        """
        for h in self.fa.get_hosts():
            for k in ['queue_usec_per_read_op',
                      'queue_usec_per_write_op',
                      'queue_usec_per_mirrored_write_op',
                      'san_usec_per_read_op',
                      'san_usec_per_write_op',
                      'san_usec_per_mirrored_write_op',
                      'service_usec_per_mirrored_write_op',
                      'service_usec_per_read_op',
                      'service_usec_per_read_op_cache_reduction',
                      'service_usec_per_write_op',
                      'usec_per_read_op',
                      'usec_per_write_op',
                      'usec_per_mirrored_write_op']:
                val = h['performance'][k]
                val = val if val is not None else 0
                self.latency.add_metric([h['name'], k], val)

    def _bandwidth(self) -> None:
        """
        Create hosts bandwidth metrics of gauge type.
        """
        for h in self.fa.get_hosts():
            for k in ['read_bytes_per_sec',
                      'write_bytes_per_sec',
                      'mirrored_write_bytes_per_sec']:
                val = h['performance'][k]
                val = val if val is not None else 0
                self.bandwidth.add_metric([h['name'], k], val)

    def _iops(self) -> None:
        """
        Create hosts IOPS metrics of gauge type.
        """
        for h in self.fa.get_hosts():
            for k in ['reads_per_sec',
                      'writes_per_sec',
                      'mirrored_writes_per_sec']:
                val = h['performance'][k]
                val = val if val is not None else 0
                self.iops.add_metric([h['name'], k], val)

    def get_metrics(self) -> None:
        self._latency()
        self._bandwidth()
        self._iops()
        yield self.latency
        yield self.bandwidth
        yield self.iops
