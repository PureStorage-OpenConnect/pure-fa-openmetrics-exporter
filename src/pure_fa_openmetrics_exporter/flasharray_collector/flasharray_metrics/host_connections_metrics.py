from prometheus_client.core import GaugeMetricFamily

class HostConnectionsMetrics():
    """
    Base class for mapping FlashArray hosts to connected volumes
    This is a helper metric that allows to correlate hosts to
    volumes in Prometheus queries.
    """

    def __init__(self, fa):
        self.fa = fa
        self.map_host_vol = GaugeMetricFamily(
                                'purefa_host_connections_info',
                                'FlashArray host volumes connections',
                                labels=['name', 'naaid'])

    def _map_host_vol(self):
        for hc in self.fa.get_host_connections():
            self.map_host_vol.add_metric([hc['host']['name'],
                                          hc['volume']['naaid']], 1)


    def get_metrics(self):
        self._map_host_vol()
        yield self.map_host_vol
