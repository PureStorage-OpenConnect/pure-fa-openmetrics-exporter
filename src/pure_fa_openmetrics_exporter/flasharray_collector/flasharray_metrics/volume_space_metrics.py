from prometheus_client.core import GaugeMetricFamily

PURE_NAA = 'naa.624a9370'

space_used_kpis = ['snapshots',
                   'snapshots_effective',
                   'total_effective',
                   'total_physical',
                   'total_provisioned',
                   'unique',
                   'unique_effective',
                   'virtual']

class VolumeSpaceMetrics():
    """
    Base class for FlashArray Prometheus volume space metrics
    """
    def __init__(self, fa_client):
        self.volumes = fa_client.volumes()
        self.data_reduction = GaugeMetricFamily(
                              'purefa_volume_space_data_reduction',
                              'FlashArray volumes data reduction ratio',
                              labels=['name', 'naaid', 'pod', 'vgroup'],
                              unit='ratio')
        self.size = GaugeMetricFamily(
                              'purefa_volume_space_size_bytes',
                              'FlashArray volumes size',
                               labels=['name', 'naaid', 'pod', 'vgroup'],
                               unit='bytes')
        self.used = GaugeMetricFamily(
                              'purefa_volume_space_used',
                              'FlashArray used space',
                               labels=['name', 'naaid', 'pod', 'vgroup',
                                       'space'],
                               unit='bytes')

    def _build_metrics(self):
        cnt_dr = 0
        cnt_sz = 0
        cnt_u = 0
        for v in self.volumes:
            vol = v['volume']
            pod = ''
            vg = ''
            if hasattr(vol.pod, 'name'):
                pod = vol.pod.name
            if hasattr(vol.volume_group, 'name'):
                vg = vol.volume_group.name
            naaid = PURE_NAA + vol.serial
            dr = getattr(vol.space, 'data_reduction')
            if dr is not None:
                cnt_dr += 1
                self.data_reduction.add_metric([vol.name, naaid, pod, vg], dr)
            sz = getattr(vol.space, 'total_provisioned')
            if sz is not None:
                cnt_sz += 1
                self.size.add_metric([vol.name, naaid, pod, vg], sz)
            for k in space_used_kpis:
                u = getattr(vol.space, k)
                if u is not None:
                    cnt_u += 1
                    self.used.add_metric([vol.name, naaid, pod, vg, k], u)
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
