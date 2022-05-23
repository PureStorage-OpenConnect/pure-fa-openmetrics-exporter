from prometheus_client.core import GaugeMetricFamily

class ArrayEventsMetrics():
    """
    Base class for FlashArray Prometheus events metrics
    """
    def __init__(self, fa_client):
        self.alerts = fa_client.alerts()
        self.open_events = GaugeMetricFamily('purefa_alerts_open',
                                             'Open alert events',
                                             labels=['severity',
                                                     'component_type',
                                                     'component_name'])
       
    def _build_metrics(self):
        cnt_e = 0
        for a in self.alerts:
            self.open_events.add_metric([a.severity,
                                         a.component_type,
                                         a.component_name], 1.0)
            cnt_e += 1
        if cnt_e == 0 :
            self.open_events = None

    def get_metrics(self):
        self._build_metrics()
        yield self.open_events
