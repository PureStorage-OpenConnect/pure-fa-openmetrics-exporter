from prometheus_client.core import GaugeMetricFamily

PURE_NAA = 'naa.624a9370'

class HostConnectionsMetrics():
    """
    Base class for mapping FlashArray hosts to connected volumes
    This is a helper metric that allows to correlate hosts to
    volumes in Prometheus queries.
    """

    def __init__(self, fa_client):
        self.connections = fa_client.connections()
        self.host_connections = GaugeMetricFamily(
                                'purefa_host_connections_info',
                                'FlashArray host volumes connections',
                                labels=['hostname', 'naaid'])

    def _build_metrics(self):
        cnt = 0
        for hc in self.connections:
            naaid = PURE_NAA + hc['volume_serial']
            self.host_connections.add_metric([hc['host'], naaid], 1)
            cnt += 1
        if cnt == 0:
            self.host_connections = None


    def get_metrics(self):
        self._build_metrics()
        yield self.host_connections
