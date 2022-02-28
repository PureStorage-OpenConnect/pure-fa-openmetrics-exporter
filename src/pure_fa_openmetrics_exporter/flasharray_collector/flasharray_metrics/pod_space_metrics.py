from prometheus_client.core import GaugeMetricFamily


class PodSpaceMetrics():
    """
    Base class for FlashArray Prometheus pod space metrics
    """

    def __init__(self, fa):
        self.fa = fa

        self.data_reduction = GaugeMetricFamily(
                                  'purefa_pod_space_datareduction_ratio',
                                  'FlashArray pod data reduction ratio',
                                  labels=['name'],
                                  unit='ratio')

        self.size = GaugeMetricFamily(
                                  'purefa_pod_space_size_bytes',
                                  'FlashArray pod size',
                                  labels=['name'])

        self.used = GaugeMetricFamily(
                                  'purefa_pod_space_used_bytes',
                                  'FlashArray pod used space',
                                  labels=['name', 'space'])

    def _data_reduction(self) -> None:
        """
        Create metrics of gauge type for pod data reduction
        """
        for p in self.fa.get_pods():
            val = p['space']['data_reduction']
            val = val if val is not None else 0
            self.data_reduction.add_metric([p['name']], val)


    def _size(self) -> None:
        """
        Create metrics of gauge type for pod size.
        """
        for p in self.fa.get_pods():
            val = p['space']['virtual']
            val = val if val is not None else 0
            self.size.add_metric([p['name']], val)

    def _used(self) -> None:
        for p in self.fa.get_pods():
            for s in ['replication',
                      'shared',
                      'snapshots',
                      'system',
                      'thin_provisioning',
                      'total_physical',
                      'total_provisioned',
                      'total_reduction',
                      'unique']:
                val = p['space'][s]
                val = val if val is not None else 0
                self.used.add_metric([p['name'], s], val)


    def get_metrics(self) -> None:
        self._data_reduction()
        self._size()
        self._used()
        yield self.data_reduction
        yield self.size
        yield self.used
