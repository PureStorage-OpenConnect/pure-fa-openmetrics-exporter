from prometheus_client.core import GaugeMetricFamily


class DirectorySpaceMetrics():
    """
    Base class for FlashArray Prometheus directory space metrics
    """

    def __init__(self, fa):
        self.fa = fa
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_directory_space_datareduction_ratio',
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

    def _data_reduction(self) -> None:
        """
        Create metrics of gauge type for directory data reduction
        Metrics values can be iterated over.
        """
        for dir in self.fa.get_directories():
            val = dir['space']['data_reduction']
            val = val if val is not None else 0
            self.data_reduction.add_metric([dir['directory_name'],
                                            dir['file_system']['name'],
                                            dir['path']], val)

    def _size(self) -> None:
        """
        Create metrics of gauge type for directory size
        Metrics values can be iterated over.
        """
        for dir in self.fa.get_directories():
            val = dir['space']['virtual']
            val = val if val is not None else 0
            self.size.add_metric([dir['directory_name'],
                                  dir['file_system']['name'],
                                  dir['path']], val)

    def _used(self) -> None:
        """
        Create metrics of gauge type for directory used space
        Metrics values can be iterated over.
        """
        for dir in self.fa.get_directories():
            for s in ['shared',
                      'snapshots',
                      'system',
                      'thin_provisioning',
                      'total_physical',
                      'total_provisioned',
                      'total_reduction',
                      'unique']:
                val = dir['space'][s]
                val = val if val is not None else 0
                self.used.add_metric([dir['directory_name'],
                                           dir['file_system']['name'],
                                           dir['path'],
                                           s], val)

    def get_metrics(self) -> None:
        self._data_reduction()
        self._size()
        self._used()
        yield self.data_reduction
        yield self.size
        yield self.used
