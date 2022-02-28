from prometheus_client.core import GaugeMetricFamily


class DirectoryPerformanceMetrics():
    """
    Base class for FlashArray Prometheus directory performance metrics
    """

    def __init__(self, fa):
        self.fa = fa
        self.latency = GaugeMetricFamily(
                           'purefa_directory_performance_latency_usec',
                           'FlashArray directory IO latency',
                           labels=['name',
                                   'filesystem',
                                   'path',
                                   'dimension'])

        self.bandwidth = GaugeMetricFamily(
                             'purefa_directory_performance_bandwidth_bytes',
                             'FlashArray directory bandwidth',
                             labels=['name',
                                     'filesystem',
                                     'path',
                                     'dimension'])

        self.iops = GaugeMetricFamily(
                        'purefa_directory_performance_iops',
                        'FlashArray directory IOPS',
                        labels=['name',
                                'filesystem',
                                'path',
                                'dimension'])

        self.avg_bsz = GaugeMetricFamily(
                           'purefa_directory_performance_avg_block_bytes',
                           'FlashArray directory avg block size',
                           labels=['name',
                                   'filesystem',
                                   'path',
                                   'dimension'])

    def _latency(self) -> None:
        """
        Create directories latency metrics of gauge type.
        Metrics values can be iterated over.
        """
        for dir in self.fa.get_directories():
            for p in ['usec_per_read_op',
                      'usec_per_write_op',
                      'usec_per_other_op']:
                val = dir['performance'][p]
                val = val if val is not None else 0
                self.latency.add_metric([dir['directory_name'],
                                           dir['file_system']['name'],
                                           dir['path'],
                                           p], val)

    def _bandwidth(self) -> None:
        """
        Create directories bandwidth metrics of gauge type.
        Metrics values can be iterated over.
        """
        for dir in self.fa.get_directories():
            for p in ['write_bytes_per_sec',
                      'read_bytes_per_sec']:
                val = dir['performance'][p]
                val = val if val is not None else 0
                self.bandwidth.add_metric([dir['directory_name'],
                                           dir['file_system']['name'],
                                           dir['path'],
                                           p], val)

    def _iops(self) -> None:
        """
        Create directories IOPS bandwidth metrics of gauge type.
        Metrics values can be iterated over.
        """
        for dir in self.fa.get_directories():
            for p in ['reads_per_sec',
                      'writes_per_sec',
                      'others_per_sec']:
                val = dir['performance'][p]
                val = val if val is not None else 0
                self.iops.add_metric([dir['directory_name'],
                                      dir['file_system']['name'],
                                      dir['path'],
                                      p], val)

    def _avg_block_size(self) -> None:
        """
        Create directories average block size performance metrics of gauge type.
        Metrics values can be iterated over.
        """
        for dir in self.fa.get_directories():
            for p in ['bytes_per_op',
                      'bytes_per_read',
                      'bytes_per_write']:
                val = dir['performance'][p]
                val = val if val is not None else 0
                self.avg_bsz.add_metric([dir['directory_name'],
                                         dir['file_system']['name'],
                                         dir['path'],
                                         p], val)

    def get_metrics(self) -> None:
        self._latency()
        self._bandwidth()
        self._iops()
        self._avg_block_size()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
