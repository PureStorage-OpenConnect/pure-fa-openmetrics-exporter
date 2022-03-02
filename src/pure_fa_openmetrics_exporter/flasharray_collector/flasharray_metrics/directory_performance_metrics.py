from prometheus_client.core import GaugeMetricFamily

class DirectoryPerformanceMetrics():
    """
    Base class for FlashArray Prometheus directory performance metrics
    """

    def __init__(self, fa_client):
        self.latency = None
        self.iops = None
        self.bandwidth = None
        self.avg_bsz = None
        self.directories = fa_client.directories()


    def _performance(self):
        """
        Create directories performance metrics of gauge type.
        """

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

        for dir in self.directories:
            d = dir['directory']
            p = dir['performance']
            self.latency.add_metric([d.directory_name,
                                     d.file_system.name,
                                     d.path,
                                     'usec_per_read_op'], p.usec_per_read_op or 0)
            self.latency.add_metric([d.directory_name,
                                     d.file_system.name,
                                     d.path,
                                     'usec_per_write_op'], p.usec_per_write_op or 0)
            self.latency.add_metric([d.directory_name,
                                     d.file_system.name,
                                     d.path,
                                     'usec_per_other_op'], p.usec_per_other_op or 0)
            self.bandwidth.add_metric([d.directory_name,
                                       d.file_system.name,
                                       d.path,
                                       'write_bytes_per_sec'], p.write_bytes_per_sec or 0)
            self.bandwidth.add_metric([d.directory_name,
                                       d.file_system.name,
                                       d.path,
                                       'read_bytes_per_sec'], p.read_bytes_per_sec or 0)

            self.iops.add_metric([d.directory_name,
                                  d.file_system.name,
                                  d.path,
                                  'reads_per_sec'], p.reads_per_sec or 0)

            self.iops.add_metric([d.directory_name,
                                  d.file_system.name,
                                  d.path,
                                  'writes_per_sec'], p.writes_per_sec or 0)
            self.iops.add_metric([d.directory_name,
                                  d.file_system.name,
                                  d.path,
                                  'others_per_sec'], p.others_per_sec or 0)

            self.avg_bsz.add_metric([d.directory_name,
                                     d.file_system.name,
                                     d.path,
                                     'bytes_per_op'], p.bytes_per_op or 0)
            self.avg_bsz.add_metric([d.directory_name,
                                     d.file_system.name,
                                     d.path,
                                     'bytes_per_read'], p.bytes_per_read or 0)
            self.avg_bsz.add_metric([d.directory_name,
                                     d.file_system.name,
                                     d.path,
                                     'bytes_per_write'], p.bytes_per_write or 0)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.bandwidth
        yield self.iops
        yield self.avg_bsz
