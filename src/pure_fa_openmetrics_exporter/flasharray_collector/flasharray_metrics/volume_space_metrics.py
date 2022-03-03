from prometheus_client.core import GaugeMetricFamily

PURE_NAA = 'naa.624a9370'

class VolumeSpaceMetrics():
    """
    Base class for FlashArray Prometheus volume space metrics
    """
    def __init__(self, fa_client):
        self.data_reduction = None
        self.size = None
        self.used = None
        self.volumes = fa_client.volumes()

    def _space(self):
        self.data_reduction = GaugeMetricFamily(
                                  'purefa_volume_space_data_reduction_ratio',
                                  'FlashArray volumes data reduction ratio',
                                  labels=['name', 'naaid', 'pod', 'vgroup'],
                                  unit='ratio')

        self.size = GaugeMetricFamily(
                                  'purefa_volume_space_size_bytes',
                                  'FlashArray volumes size',
                                   labels=['name', 'naaid', 'pod', 'vgroup'])

        self.used = GaugeMetricFamily(
                                  'purefa_volume_space_used_bytes',
                                  'FlashArray used space',
                                   labels=['name', 'naaid', 'pod', 'vgroup',
                                           'space'])
        for v in self.volumes:
            vol = v['volume']
            pod = ''
            vg = ''
            if hasattr(vol.pod, 'name'):
                pod = vol.pod.name
            if hasattr(vol.volume_group, 'name'):
                vg = vol.volume_group.name
            naaid = PURE_NAA + vol.serial
            self.data_reduction.add_metric([vol.name, naaid, pod, vg],
                                           vol.space.data_reduction or 0)
            self.size.add_metric([vol.name, naaid, pod, vg],
                                 vol.provisioned or 0)

            self.used.add_metric([vol.name, naaid, pod, vg, 'snapshots'],
                                 vol.space.snapshots or 0)
            self.used.add_metric([vol.name, naaid, pod, vg, 'snapshots_effective'],
                                 vol.space.snapshots_effective or 0)
            self.used.add_metric([vol.name, naaid, pod, vg, 'total_effective'],
                                 vol.space.total_effective or 0)
            self.used.add_metric([vol.name, naaid, pod, vg, 'total_physical'],
                                 vol.space.total_physical or 0)
            self.used.add_metric([vol.name, naaid, pod, vg, 'total_provisioned'],
                                 vol.space.total_provisioned or 0)
            self.used.add_metric([vol.name, naaid, pod, vg, 'unique'],
                                 vol.space.unique or 0)
            self.used.add_metric([vol.name, naaid, pod, vg, 'unique_effective'],

                                 vol.space.unique_effective or 0)
            self.used.add_metric([vol.name, naaid, pod, vg, 'virtual'],
                                 vol.space.virtual or 0)

    def get_metrics(self):
        self._space()
        yield self.data_reduction
        yield self.size
        yield self.used
