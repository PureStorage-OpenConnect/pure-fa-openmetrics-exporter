from prometheus_client.core import GaugeMetricFamily


class ArraySpaceMetrics():
    """
    Base class for FlashArray Prometheus array space metrics
    """
    def __init__(self, fa):
        self.fa = fa
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_array_space_datareduction_ratio',
                                  'FlashArray overall data reduction',
                                  labels=[],
                                  unit='ratio')

        self.capacity = GaugeMetricFamily(
                                  'purefa_array_space_capacity_bytes',
                                  'FlashArray overall space capacity',
                                  labels=[])

        self.provisioned = GaugeMetricFamily(
                                  'purefa_array_space_provisioned_bytes',
                                  'FlashArray overall provisioned space',
                                  labels=[])

        self.used = GaugeMetricFamily(
                                  'purefa_array_space_used_bytes',
                                  'FlashArray overall used space',
                                  labels=['space'])

    def _data_reduction(self) -> None:
        """
        Create metrics of gauge type for array data reduction.
        Metrics values can be iterated over.
        """
        val = self.fa.get_array()['space']['data_reduction']
        val = val if val is not None else 0
        self.data_reduction.add_metric([], val)

    def _capacity(self) -> None:
        """
        Create metrics of gauge type for array capacity indicators.
        Metrics values can be iterated over.
        """
        val = self.fa.get_array()['space']['capacity']
        val = val if val is not None else 0
        self.capacity.add_metric([], val)

    def _provisioned(self) -> None:
        """
        Create metrics of gauge type for array provisioned space indicators.
        Metrics values can be iterated over.
        """
        val = self.fa.get_array()['space']['total_provisioned']
        val = val if val is not None else 0
        self.provisioned.add_metric([], val) 

    def _used(self) -> None:
        """
        Create metrics of gauge type for array used space indicators.
        Metrics values can be iterated over.
        """
        for k in ['replication',
                  'shared',
                  'shared_effective',
                  'snapshots',
                  'snapshots_effective',
                  'system',
                  'thin_provisioning', 
                  'total_physical',
                  'total_effective',
                  'unique',
                  'unique_effective',
                  'virtual']:
                  
            val = self.fa.get_array()['space'][k]
            val = val if val is not None else 0
            self.used.add_metric([k], val)

    def get_metrics(self) -> None:
        self._data_reduction()
        self._capacity()
        self._provisioned()
        self._used()
        yield self.data_reduction
        yield self.capacity
        yield self.provisioned
        yield self.used
