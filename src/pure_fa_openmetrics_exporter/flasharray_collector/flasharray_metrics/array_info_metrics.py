from prometheus_client.core import InfoMetricFamily

class ArrayInfoMetrics():
    """
    Base class for FlashArray OpenMetrics array info
    """
    def __init__(self, fa_client):
        self.info = None
        self.array = fa_client.arrays()[0]

    def _array(self):
        """Assemble a simple information metric defining the scraped system."""

        self.array_info = InfoMetricFamily(
                                      'purefa',
                                      'FlashArray system information',
                                      value={'array_name': self.array.name,
                                            'system_id': self.array.id,
                                            'os': self.array.os,
                                            'version': self.array.version
                                            })

    def get_metrics(self):
        self._array()
        yield self.array_info
