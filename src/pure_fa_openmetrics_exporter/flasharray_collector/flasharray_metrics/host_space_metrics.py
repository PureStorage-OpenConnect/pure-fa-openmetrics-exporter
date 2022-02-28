from prometheus_client.core import GaugeMetricFamily


class HostSpaceMetrics():
    """
    Base class for FlashArray Prometheus host space metrics
    """

    def __init__(self, fa):
        self.fa = fa

        self.data_reduction = GaugeMetricFamily(
                                  'purefa_host_space_datareduction_ratio',
                                  'FlashArray host data reduction ratio',
                                  labels=['name', 'hostgroup'],
                                  unit='ratio')

        self.size = GaugeMetricFamily(
                                   'purefa_host_space_size_bytes',
                                   'FlashArray host size',
                                   labels=['name', 'hostgroup'])

        self.used = GaugeMetricFamily(
                                   'purefa_host_space_used_bytes',
                                   'FlashArray host used space',
                                   labels=['name', 'hostgroup', 'space'])

    def _data_reduction(self) -> None:
        """
        Create metrics of gauge type for host data reduction
        """
        for h in self.fa.get_hosts():
            hg_name = h['host_group']['name']
            hg_name = hg_name if hg_name is not None else ''
            val = h['space']['data_reduction']
            val = val if val is not None else 0
            self.data_reduction.add_metric([h['name'], hg_name], val)

    def _size(self) -> None:
        """
        Create metrics of gauge type for host size
        """
        for h in self.fa.get_hosts():
            hg_name = h['host_group']['name']
            hg_name = hg_name if hg_name is not None else ''
            val = h['space']['virtual']
            val = val if val is not None else 0
            self.size.add_metric([h['name'], hg_name], val)

    def _used(self) -> None:
        """
        Create metrics of gauge type for host used space
        """
        for h in self.fa.get_hosts():
            hg_name = h['host_group']['name']
            hg_name = hg_name if hg_name is not None else ''
            for s in ['shared',
                      'snapshots',
                      'system',
                      'thin_provisioning',
                      'total_physical',
                      'total_provisioned',
                      'total_reduction',
                      'unique']:
                val = h['space'][s]
                val = val if val is not None else 0
                self.used.add_metric([h['name'], hg_name, s], val)

    def get_metrics(self) -> None:
        self._data_reduction()
        self._size()
        self._used()
        yield self.data_reduction
        yield self.size
        yield self.used
