from prometheus_client.core import GaugeMetricFamily

space_used_kpis = ['replication',
                   'shared',
                   'snapshots',
                   'total_physical',
                   'total_provisioned',
                   'unique']

class PodSpaceMetrics():
    """
    Base class for FlashArray Prometheus pod space metrics
    """

    def __init__(self, fa_client):
        self.pods = fa_client.pods()
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_pod_space_data_reduction',
                                  'FlashArray pod data reduction ratio',
                                  labels=['name'],
                                  unit='ratio')
        self.size = GaugeMetricFamily(
                                  'purefa_pod_space_size',
                                  'FlashArray pod size',
                                  labels=['name'],
                                  unit='bytes')
        self.used = GaugeMetricFamily(
                                  'purefa_pod_space_used',
                                  'FlashArray pod used space',
                                  labels=['name', 'space'],
                                  unit='bytes')

    def _build_metrics(self):
        cnt_dr = 0
        cnt_sz = 0
        cnt_u = 0
        for p in self.pods:
            pod = p['pod']
            dr = getattr(pod.space, 'data_reduction')
            if dr is not None:
                cnt_dr += 1
                self.data_reduction.add_metric([pod.name], dr)
            sz = getattr(pod.space, 'virtual')
            if sz is not None:
                cnt_sz += 1
                self.size.add_metric([pod.name], sz)
            for k in space_used_kpis:
                u = getattr(pod.space, k)
                if u is not None:
                    cnt_u += 1
                    self.used.add_metric([pod.name, k], u)
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
