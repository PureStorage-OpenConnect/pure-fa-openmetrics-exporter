from prometheus_client.core import GaugeMetricFamily

class PodStatusMetrics():
    """
    Base class for FlashArray Prometheus pod stattus metrics
    """

    def __init__(self, fa_client):
        self.status = None
        self.mediator_status = None
        self.pods = fa_client.pods()


    def _status(self):
        """
        Create pods status metrics of gauge type, with pod name, array id and
        array name as label.
        Metrics values can be iterated over.
        """
        self.status = GaugeMetricFamily('purefa_pod_status',
                                        'FlashArray pod status',
                                        labels=['name', 'array_name'])

        self.mediator_status = GaugeMetricFamily(
                                   'purefa_pod_mediator_status',
                                   'FlashArray pod mediatorstatus',
                                   labels=['name', 'array_name'])

        for p in self.pods:
            pod = p['pod']
            arrays = pod.arrays
            self.status.add_metric([pod.name, arrays[0].name],
                                   1 if arrays[0].status == 'online' else 0)
            self.mediator_status.add_metric([pod.name, arrays[0].name],
                          1 if arrays[0].mediator_status == 'online' else 0)
            if len(arrays) == 1:
                continue
            self.status.add_metric([pod.name, arrays[1].name], 
                                   1 if arrays[1].status == 'online' else 0)
            self.mediator_status.add_metric([pod.name, arrays[1].name],
                          1 if arrays[1].mediator_status == 'online' else 0)

    def get_metrics(self):
        self._status()
        yield self.status
        yield self.mediator_status
