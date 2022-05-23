from prometheus_client.core import InfoMetricFamily

class ArrayInfoMetrics():
    """
    Base class for FlashArray OpenMetrics array info
    """
    def __init__(self, fa_client):
        self.array = fa_client.arrays()[0]
        self.array_info = InfoMetricFamily('purefa',
                                           'FlashArray system information',
                                           labels = [])

    def _build_metrics(self):
        array = self.array['array']
        self.array_info.add_metric([], {'array_name': array.name,
                                        'system_id': array.id,
                                        'os' : array.os,
                                        'version': array.version})

    def get_metrics(self):
        self._build_metrics()
        yield self.array_info
