from prometheus_client.core import GaugeMetricFamily

space_used_kpis = ['shared',
                   'snapshots',
                   'system',
                   'total_physical',
                   'unique',
                   'virtual']

class ArraySpaceMetrics():
    """
    Base class for FlashArray array space metrics.
    """

    def __init__(self, fa_client):
        self.array = fa_client.arrays()[0]['array']
        self.data_reduction = GaugeMetricFamily(
                                       'purefa_array_space_data_reduction',
                                       'FlashArray overall data reduction',
                                       labels=[],
                                       unit='ratio')
        self.capacity = GaugeMetricFamily(
                                 'purefa_array_space_capacity',
                                 'FlashArray overall space capacity',
                                 labels=[],
                                 unit='bytes')
        self.used = GaugeMetricFamily(
                         'purefa_array_space_used',
                         'FlashArray overall used space',
                         labels=['space'],
                         unit='bytes')

    def _build_metrics(self):
        v = getattr(self.array.space, 'data_reduction')
        if v is None:
            self.data_reduction = None
        else:
            self.data_reduction.add_metric([], v)
        v = getattr(self.array, 'capacity')
        if v is None:
            self.capacity = None
        else:
            self.capacity.add_metric([], v)
        cnt_u = 0
        for k in space_used_kpis:
            v = getattr(self.array.space, k)
            if v is None:
                continue
            cnt_u += 1
            self.used.add_metric([k], v)
        if cnt_u == 0:
            self.used = None

    def get_metrics(self):
        self._build_metrics()
        yield self.data_reduction
        yield self.capacity
        yield self.used
