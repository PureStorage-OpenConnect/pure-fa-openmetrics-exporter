from prometheus_client.core import GaugeMetricFamily

space_used_kpis = ['snapshots',
                   'total_physical',
                   'total_provisioned',
                   'unique']

class HostSpaceMetrics():
    """
    Base class for FlashArray Prometheus host space metrics
    """

    def __init__(self, fa_client):
        self.hosts = fa_client.hosts()
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_host_space_data_reduction',
                                  'FlashArray host data reduction ratio',
                                  labels=['name', 'hostgroup'],
                                  unit='ratio')

        self.size = GaugeMetricFamily('purefa_host_space_size',
                                      'FlashArray host size',
                                      labels=['name', 'hostgroup'],
                                      unit='bytes')

        self.used = GaugeMetricFamily('purefa_host_space_used',
                                      'FlashArray host used space',
                                      labels=['name', 'hostgroup', 'space'],
                                      unit='bytes')

    def _build_metrics(self):
        cnt_dr = 0
        cnt_sz = 0
        cnt_u = 0
        for h in self.hosts:
            host = h['host']
            if not host.is_local:
                continue
            hg = ''
            if hasattr(host.host_group, 'name'):
                hg = host.host_group.name
            dr = getattr(host.space, 'data_reduction')
            if dr is not None:
                cnt_dr += 1
                self.data_reduction.add_metric([host.name, hg], dr)
            sz = getattr(host.space, 'virtual')
            if sz is not None:
                cnt_sz += 1
                self.size.add_metric([host.name, hg], sz)
            for k in space_used_kpis:
                u = getattr(host.space, k)
                if u is not None:
                    cnt_u += 1
                    self.used.add_metric([host.name, hg, k], u)
        if cnt_dr == 0:
            self.data_reduction = None
        if cnt_sz == 0:
            self.size = None
        if cnt_u == 0:
            self.used = None

    def get_metrics(self):
        self._build_metrics()
        yield self.data_reduction
        yield self.size
        yield self.used
