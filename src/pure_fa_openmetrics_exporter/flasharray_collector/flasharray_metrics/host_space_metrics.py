from prometheus_client.core import GaugeMetricFamily


class HostSpaceMetrics():
    """
    Base class for FlashArray Prometheus host space metrics
    """

    def __init__(self, fa_client):
        self.data_reduction = None
        self.size = None
        self.used = None
        self.hosts = fa_client.hosts()

    def _space(self):
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_host_space_data_reduction_ratio',
                                  'FlashArray host data reduction ratio',
                                  labels=['name', 'hostgroup'],
                                  unit='ratio')

        self.size = GaugeMetricFamily(
                                   'purefa_host_space_size_bytes',
                                   'FlashArray host size',
                                   labels=['name', 'hostgroup'])

        self.used = GaugeMetricFamily(
                                   'purefa_host_space_used_bytes',
                                   'FlashArray host used space',
                                   labels=['name', 'hostgroup', 'space'])

        for h in self.hosts:
            host = h['host']
            if not host.is_local:
                continue
            hg = ''
            if hasattr(host.host_group, 'name'):
                hg = host.host_group.name
            self.data_reduction.add_metric([host.name, hg], 
                                            host.space.data_reduction or 0)

            self.size.add_metric([host.name, hg], host.space.virtual or 0)
            self.used.add_metric([host.name, hg, 'snapshots'],
                                 host.space.snapshots or 0)
            self.used.add_metric([host.name, hg, 'total_physical'],
                                 host.space.total_physical or 0)
            self.used.add_metric([host.name, hg, 'total_provisioned'],
                                 host.space.total_provisioned or 0)
            self.used.add_metric([host.name, hg, 'unique'],
                                 host.space.unique or 0)

    def get_metrics(self):
        self._space()
        yield self.data_reduction
        yield self.size
        yield self.used
