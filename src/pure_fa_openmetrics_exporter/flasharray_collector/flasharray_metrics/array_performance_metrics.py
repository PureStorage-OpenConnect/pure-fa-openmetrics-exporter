from prometheus_client.core import GaugeMetricFamily

class ArrayPerformanceMetrics():
    """
    Base class for FlashArray OpenMetrics array performance metrics
    """

    def __init__(self, fa_client):
        self.latency = None
        self.iops = None
        self.bandwidth = None
        self.avg_bsz = None
        self.array_perf = fa_client.arrays_performance()


    def _performance(self):
        """
        Create array performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                           'purefa_array_performance_latency_usec',
                           'FlashArray array latency',
                           labels=['protocol', 'dimension'])

        self.iops = GaugeMetricFamily('purefa_array_performance_iops',
                                      'FlashArray IOPS',
                                      labels=['protocol', 'dimension'])

        self.bandwidth = GaugeMetricFamily(
                             'purefa_array_performance_bandwidth_bytes',
                             'FlashArray bandwidth',
                             labels=['protocol', 'dimension'])

        self.avg_bsz = GaugeMetricFamily(
                           'purefa_array_performance_average_block_bytes',
                           'FlashArray array average block size',
                           labels=['protocol', 'dimension'])

        for p in self.array_perf:

            self.latency.add_metric([p, 'local_queue_usec_per_op'],
                                self.array_perf[p].local_queue_usec_per_op or 0)
            self.latency.add_metric([p, 'queue_usec_per_read_op'],
                                self.array_perf[p].queue_usec_per_read_op or 0)
            self.latency.add_metric([p, 'queue_usec_per_write_op'],
                                self.array_perf[p].queue_usec_per_write_op or 0)
            self.latency.add_metric([p, 'queue_usec_per_mirrored_write_op'],
                                self.array_perf[p].queue_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([p, 'san_usec_per_read_op'],
                                self.array_perf[p].san_usec_per_read_op or 0)
            self.latency.add_metric([p, 'san_usec_per_write_op'],
                                self.array_perf[p].san_usec_per_write_op or 0)
            self.latency.add_metric([p, 'san_usec_per_mirrored_write_op'],
                                self.array_perf[p].san_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([p, 'service_usec_per_mirrored_write_op'],
                                self.array_perf[p].service_usec_per_mirrored_write_op or 0)
            self.latency.add_metric([p, 'service_usec_per_read_op'],
                                self.array_perf[p].service_usec_per_read_op or 0)
            self.latency.add_metric([p, 'service_usec_per_read_op_cache_reduction'],
                                self.array_perf[p].service_usec_per_read_op_cache_reduction or 0)
            self.latency.add_metric([p, 'service_usec_per_write_op'],
                                self.array_perf[p].service_usec_per_write_op or 0)
            self.latency.add_metric([p, 'usec_per_read_op'],
                                self.array_perf[p].usec_per_read_op or 0)
            self.latency.add_metric([p, 'usec_per_write_op'],
                                self.array_perf[p].usec_per_write_op or 0)
            self.latency.add_metric([p, 'usec_per_mirrored_write_op'],
                                self.array_perf[p].usec_per_mirrored_write_op or 0)
            self.latency.add_metric([p, 'usec_per_other_op'],
                                self.array_perf[p].usec_per_other_op or 0)

            self.bandwidth.add_metric([p, 'read_bytes_per_sec'],
                                  self.array_perf[p].read_bytes_per_sec or 0)
            self.bandwidth.add_metric([p, 'write_bytes_per_sec'],
                                  self.array_perf[p].write_bytes_per_sec or 0)
            self.bandwidth.add_metric([p, 'mirrored_write_bytes_per_sec'],
                                  self.array_perf[p].mirrored_write_bytes_per_sec or 0)

            self.iops.add_metric([p, 'reads_per_sec'],
                             self.array_perf[p].reads_per_sec or 0)
            self.iops.add_metric([p, 'writes_per_sec'],
                             self.array_perf[p].writes_per_sec or 0)
            self.iops.add_metric([p, 'mirrored_writes_per_sec'],
                             self.array_perf[p].mirrored_writes_per_sec or 0)
            self.iops.add_metric([p, 'others_per_sec'],
                             self.array_perf[p].others_per_sec or 0)

            self.avg_bsz.add_metric([p, 'bytes_per_read'],
                                self.array_perf[p].bytes_per_read or 0)
            self.avg_bsz.add_metric([p, 'bytes_per_write'],
                                self.array_perf[p].bytes_per_write or 0)
            self.avg_bsz.add_metric([p, 'bytes_per_op'],
                                self.array_perf[p].bytes_per_op or 0)
        

    def get_metrics(self) -> None:
        self._performance()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
