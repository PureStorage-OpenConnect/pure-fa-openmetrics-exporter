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

        arr = self.array['array']
        self.data_reduction.add_metric([], arr.space.data_reduction or 0)
        self.capacity.add_metric([], arr.capacity or 0)

        self.used.add_metric(['shared'], arr.space.shared or 0)
        self.used.add_metric(['snapshots'], arr.space.snapshots or 0)
        self.used.add_metric(['system'], arr.space.system or 0)
        self.used.add_metric(['total_physical'], arr.space.total_physical or 0)
        self.used.add_metric(['unique'], arr.space.unique or 0)
        self.used.add_metric(['virtual'], arr.space.virtual or 0)

    def get_metrics(self):
        self._space()
        yield self.data_reduction
        yield self.capacity
        yield self.used
