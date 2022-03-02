from prometheus_client.core import GaugeMetricFamily

class DirectorySpaceMetrics():
    """
    Base class for FlashArray Prometheus directory space metrics
    """

    def __init__(self, fa_client):
        self.data_reduction = None
        self.size = None
        self.used = None
        self.directories = fa_client.directories()
        
    def _space(self):
        """
        Create metrics of gauge type for directories space indicators.
        """
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_directory_space_data_reduction_ratio',
                                  'FlashArray directories data reduction ratio',
                                  labels=['name', 'filesystem', 'path'],
                                  unit='ratio')

        self.size = GaugeMetricFamily(
                                  'purefa_directory_space_size_bytes',
                                  'FlashArray directories size',
                                  labels=['name', 'filesystem', 'path'])

        self.used = GaugeMetricFamily(
                                  'purefa_directory_space_used_bytes',
                                  'FlashArray used space',
                                  labels=['name', 'filesystem', 'path',
                                          'space'])

        for dir in self.directories:
            d = dir['directory']
            self.data_reduction.add_metric([d.directory_name,
                                            d.file_system.name,
                                            d.path], d.space.data_reduction or 0)

            self.size.add_metric([d.directory_name,
                                  d.file_system.name,
                                  d.path], d.space.virtual or 0)
            self.used.add_metric([d.directory_name,
                                  d.file_system.name,
                                  d.path,
                                  'snapshots'], d.space.snapshots or 0)
            self.used.add_metric([d.directory_name,
                                  d.file_system.name,
                                  d.path,
                                  'total_physical'], d.space.total_physical or 0)
            self.used.add_metric([d.directory_name,
                                  d.file_system.name,
                                  d.path,
                                  'unique'], d.space.unique or 0)

    def get_metrics(self):
        self._space()
        yield self.data_reduction
        yield self.size
        yield self.used
