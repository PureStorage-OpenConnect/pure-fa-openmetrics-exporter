from prometheus_client.core import GaugeMetricFamily

PURE_NAA = 'naa.624a9370'

class HostConnectionsMetrics():
    """
    Base class for mapping FlashArray hosts to connected volumes
    This is a helper metric that allows to correlate hosts to
    volumes in Prometheus queries.
    """

    def __init__(self, fa_client):
        self.host_connections = None
        self.connections = fa_client.connections()

    def _map_host_vol(self):
        self.host_connections = GaugeMetricFamily(
                                'purefa_host_connections_info',
                                'FlashArray host volumes connections',
                                labels=['hostname', 'naaid'])
        for hc in self.connections:
            naaid = PURE_NAA + hc['volume_serial']
            self.host_connections.add_metric([hc['host'], naaid], 1)


    def get_metrics(self):
        self._map_host_vol()
        yield self.host_connections
