from prometheus_client.core import GaugeMetricFamily


class VolumeSpaceMetrics():
    """
    Base class for FlashArray Prometheus volume space metrics
    """

    def __init__(self, fa):
        self.fa = fa

        self.data_reduction = GaugeMetricFamily(
                                  'purefa_volume_space_datareduction_ratio',
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

    def _data_reduction(self) -> None:
        """
        Create metrics of gauge type for volume data reduction
        """
        for v in self.fa.get_volumes():
            pod = v['pod']['name']
            pod = pod  if pod is not None else ''
            vg = v['volume_group']['name'] 
            vg = vg if vg is not None else ''
            val = v['space']['data_reduction']
            val = val if val is not None else 0
            self.data_reduction.add_metric([v['name'], v['naaid'], pod, vg], val)

    def _size(self) -> None:
        """
        Create metrics of gauge type for volume size
        """
        for v in self.fa.get_volumes():
            pod = v['pod']['name']
            pod = pod  if pod is not None else ''
            vg = v['volume_group']['name'] 
            vg = vg if vg is not None else ''
            val = v['provisioned']
            val = val if val is not None else 0
            self.size.add_metric([v['name'], v['naaid'], pod, vg], val)

    def _used(self) -> None:
        """
        Create metrics of gauge type for volume used space
        """
        for v in self.fa.get_volumes():
            for s in ['shared',
                      'snapshots',
                      'snapshots_effective',
                      'system',
                      'thin_provisioning',
                      'total_effective',
                      'total_physical',
                      'total_provisioned',
                      'total_reduction',
                      'unique',
                      'virtual']:
                pod = v['pod']['name']
                pod = pod  if pod is not None else ''
                vg = v['volume_group']['name'] 
                vg = vg if vg is not None else ''
                val = v['space'][s]
                val = val if val is not None else 0
                self.used.add_metric([v['name'], v['naaid'], pod, vg, s],
                                          val)

    def get_metrics(self) -> None:
        self._data_reduction()
        self._size()
        self._used()
        yield self.data_reduction
        yield self.size
        yield self.used
