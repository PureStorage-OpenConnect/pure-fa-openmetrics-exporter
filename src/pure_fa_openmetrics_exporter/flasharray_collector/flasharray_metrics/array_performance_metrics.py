from prometheus_client.core import GaugeMetricFamily

class ArrayPerformanceMetrics():
    """
    Base class for FlashArray OpenMetrics array performance metrics
    """

    def __init__(self, fa_client):
        self.latency = None
        self.iops = None
        self.array_perf = fa_client.arrays_performance()


    def _performance(self):
        """
        Create array performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                           'purefa_array_performance_latency_usec',
                           'FlashArray array latency',
                           labels=['dimension'])

        self.iops = GaugeMetricFamily('purefa_array_performance_iops',
                                      'FlashArray IOPS',
                                      labels=['dimension'])

        self.bandwidth = GaugeMetricFamily(
                             'purefa_array_performance_bandwidth_bytes',
                             'FlashArray bandwidth',
                             labels=['dimension'])

        self.avg_bsz = GaugeMetricFamily(
                           'purefa_array_performance_average_block_bytes',
                           'FlashArray array average block size',
                           labels=['dimension'])

        self.qdepth = GaugeMetricFamily('purefa_array_performance_qdepth',
                                        'FlashArray queue depth',
                                        labels=[])

    def _latency(self) -> None:
        """
        Create array latency performance metrics of gauge type.
        Metrics values can be iterated over.
        """
        for k in ['local_queue_usec_per_op',
                  'queue_usec_per_read_op',
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
                  'usec_per_mirrored_write_op',
                  'usec_per_other_op']:
            val = self.fa.get_array()['performance'][k]
            val = val if val is not None else 0
            self.latency.add_metric([k], val)

    def _bandwidth(self) -> None:
        """
        Create array bandwidth performance metrics of gauge type.
        Metrics values can be iterated over.
        """
        for k in ['read_bytes_per_sec',
                  'write_bytes_per_sec',
                  'mirrored_write_bytes_per_sec']:
            val = self.fa.get_array()['performance'][k]
            val = val if val is not None else 0
            self.bandwidth.add_metric([k], val)

    def _iops(self) -> None:
        """
        Create array iops performance metrics of gauge type.
        Metrics values can be iterated over.
        """
        for k in ['reads_per_sec',
                  'writes_per_sec',
                  'mirrored_writes_per_sec',
                  'others_per_sec']:
            val = self.fa.get_array()['performance'][k]
            val = val if val is not None else 0
            self.iops.add_metric([k], val)

    def _avg_block_size(self) -> None:
        """
        Create array average block size performance metrics of gauge type.
        Metrics values can be iterated over.
        """
        for k in ['bytes_per_read',
                  'bytes_per_write',
                  'bytes_per_op']:
            val = self.fa.get_array()['performance'][k]
            val = val if val is not None else 0
            self.avg_bsz.add_metric([k], val)

    def _qdepth(self) -> None:
        """
        Create array queue depth performance metric of gauge type.
        Metrics values can be iterated over.
        """
        val = self.fa.get_array()['performance']['queue_depth']
        self.qdepth.add_metric([], val if val is not None else '0')

    def get_metrics(self) -> None:
        self._performance()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
        yield self.qdepth
