from prometheus_client.core import GaugeMetricFamily

class ArraySpaceMetrics():
    """
    Base class for FlashArray Prometheus array space metrics
    """
    def __init__(self, fa_client):
        self.data_reduction = None
        self.capacity = None
        self.used = None
        self.array = fa_client.arrays()[0]

    def _space(self):
        """
        Create metrics of gauge type for array space indicators.
        """
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_array_space_data_reduction_ratio',
                                  'FlashArray overall data reduction',
                                  labels=[],
                                  unit='ratio')

        self.capacity = GaugeMetricFamily(
                                  'purefa_array_space_capacity_bytes',
                                  'FlashArray overall space capacity',
                                  labels=[])

        self.used = GaugeMetricFamily(
                                  'purefa_array_space_used_bytes',
                                  'FlashArray overall used space',
                                  labels=['space'])

        self.data_reduction.add_metric([], self.array.space.data_reduction or 0)
        self.capacity.add_metric([], self.array.capacity or 0)

        self.used.add_metric(['shared'], self.array.space.shared or 0)
        self.used.add_metric(['snapshots'], self.array.space.snapshots or 0)
        self.used.add_metric(['system'], self.array.space.system or 0)
        self.used.add_metric(['total_physical'], self.array.space.total_physical or 0)
        self.used.add_metric(['unique'], self.array.space.unique or 0)
        self.used.add_metric(['virtual'], self.array.space.virtual or 0)

    def get_metrics(self) -> None:
        self._space()
        yield self.data_reduction
        yield self.capacity
        yield self.used
