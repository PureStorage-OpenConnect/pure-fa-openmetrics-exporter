from prometheus_client.core import GaugeMetricFamily

space_used_kpis = ['snapshots',
                   'total_physical',
                   'unique']

class DirectorySpaceMetrics():
    """
    Base class for FlashArray Prometheus directory space metrics
    """

    def __init__(self, fa_client):
        self.directories = fa_client.directories()
        self.data_reduction = GaugeMetricFamily(
                                 'purefa_directory_space_data_reduction',
                                 'FlashArray directories data reduction ratio',
                                 labels=['name', 'filesystem', 'path'],
                                 unit='ratio')
        self.size = GaugeMetricFamily(
                         'purefa_directory_space_size',
                         'FlashArray directories size',
                         labels=['name', 'filesystem', 'path'],
                         unit='bytes')
        self.used = GaugeMetricFamily(
                         'purefa_directory_space_used',
                         'FlashArray used space',
                         labels=['name', 'filesystem', 'path', 'space'],
                         unit='bytes')

    def _build_metrics(self):
        cnt_dr = 0
        cnt_s = 0
        cnt_u = 0
        for dir in self.directories:
            d = dir['directory']
            v = getattr(d.space, 'data_reduction')
            if v is not None:
                self.data_reduction.add_metric([d.directory_name,
                                                d.file_system.name,
                                                d.path], v)
                cnt_dr += 1

            v = getattr(d.space, 'virtual')
            if v is not None:
                self.size.add_metric([d.directory_name,
                                      d.file_system.name,
                                      d.path], v)
                cnt_s += 1
            for k in space_used_kpis:
                v = getattr(d.space, k)
                if v is None:
                    continue
                cnt_u += 1
                self.used.add_metric([d.directory_name,
                                      d.file_system.name,
                                      d.path,
                                      k], v)
        if cnt_dr == 0:
            self.data_reduction = None
        if cnt_s == 0:
            self.size = None
        if cnt_u == 0:
            self.used = None

    def get_metrics(self):
        self._build_metrics()
        yield self.data_reduction
        yield self.size
        yield self.used
