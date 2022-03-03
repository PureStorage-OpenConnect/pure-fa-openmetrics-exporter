from prometheus_client.core import GaugeMetricFamily

class PodSpaceMetrics():
    """
    Base class for FlashArray Prometheus pod space metrics
    """

    def __init__(self, fa_client):
        self.data_reduction = None
        self.size = None
        self.used = None
        self.pods = fa_client.pods()

    def _space(self):
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_pod_space_data_reduction_ratio',
                                  'FlashArray pod data reduction ratio',
                                  labels=['name'],
                                  unit='ratio')

        self.size = GaugeMetricFamily(
                                  'purefa_pod_space_size_bytes',
                                  'FlashArray pod size',
                                  labels=['name'])

        self.used = GaugeMetricFamily(
                                  'purefa_pod_space_used_bytes',
                                  'FlashArray pod used space',
                                  labels=['name', 'space'])

        for p in self.pods:
            pod = p['pod']
            self.data_reduction.add_metric([pod.name], 
                      pod.space.data_reduction or 0)
            self.size.add_metric([pod.name], 
                      pod.space.virtual or 0)
            self.used.add_metric([pod.name, 'replication'],
                      pod.space.replication or 0)
            self.used.add_metric([pod.name, 'shared'],
                      pod.space.shared or 0)
            self.used.add_metric([pod.name, 'snapshots'],
                      pod.space.snapshots or 0)
            self.used.add_metric([pod.name, 'total_physical'],
                      pod.space.total_physical or 0)
            self.used.add_metric([pod.name, 'total_provisioned'],
                      pod.space.total_provisioned or 0)
            self.used.add_metric([pod.name, 'unique'],
                      pod.space.unique or 0)

    def get_metrics(self):
        self._space()
        yield self.data_reduction
        yield self.size
        yield self.used
