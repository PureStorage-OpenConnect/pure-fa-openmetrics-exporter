from .flasharray_metrics.array_info_metrics import ArrayInfoMetrics
from .flasharray_metrics.array_hardware_metrics import ArrayHardwareMetrics
from .flasharray_metrics.array_events_metrics import ArrayEventsMetrics
from .flasharray_metrics.array_space_metrics import ArraySpaceMetrics
from .flasharray_metrics.array_performance_metrics import ArrayPerformanceMetrics
from .flasharray_metrics.volume_space_metrics import VolumeSpaceMetrics
from .flasharray_metrics.volume_performance_metrics import VolumePerformanceMetrics
from .flasharray_metrics.host_space_metrics import HostSpaceMetrics
from .flasharray_metrics.host_performance_metrics import HostPerformanceMetrics
from .flasharray_metrics.host_connections_metrics import HostConnectionsMetrics
from .flasharray_metrics.pod_status_metrics import PodStatusMetrics
from .flasharray_metrics.pod_space_metrics import PodSpaceMetrics
from .flasharray_metrics.pod_performance_metrics import PodPerformanceMetrics
from .flasharray_metrics.network_interface_metrics import NetworkInterfacePerformanceMetrics
from .flasharray_metrics.directory_space_metrics import DirectorySpaceMetrics
from .flasharray_metrics.directory_performance_metrics import DirectoryPerformanceMetrics


class FlasharrayCollector():
    """
    Instantiates the collector's methods and properties to retrieve status,
    space occupancy and performance metrics from Puretorage FlashArray.
    Provides also a 'collect' method to allow Prometheus client registry
    to work properly. This collector is capable of routing both to local
    internal REST API server and to a remote external FlashArray REST API.
    """
    def __init__(self, fa_client, request = 'all'):
        self.fa = fa_client
        self.request = request

    def collect(self):
        """Global collector method for all the collected array metrics."""
        if self.request in ['all', 'array']:
            yield from ArrayInfoMetrics(self.fa).get_metrics()
            yield from ArrayHardwareMetrics(self.fa).get_metrics()
            yield from ArrayEventsMetrics(self.fa).get_metrics()
            yield from ArraySpaceMetrics(self.fa).get_metrics()
            yield from ArrayPerformanceMetrics(self.fa).get_metrics()
            yield from NetworkInterfacePerformanceMetrics(self.fa).get_metrics()
        if self.request in ['all', 'volumes']:
            yield from VolumeSpaceMetrics(self.fa).get_metrics()
            yield from VolumePerformanceMetrics(self.fa).get_metrics()
        if self.request in ['all', 'hosts']:
            yield from HostSpaceMetrics(self.fa).get_metrics()
            yield from HostPerformanceMetrics(self.fa).get_metrics()
        if self.request in ['all', 'pods']:
            yield from PodStatusMetrics(self.fa).get_metrics()
            yield from PodSpaceMetrics(self.fa).get_metrics()
            yield from PodPerformanceMetrics(self.fa).get_metrics()
        if self.request in ['all', 'directories']:
            yield from DirectorySpaceMetrics(self.fa).get_metrics()
            yield from DirectoryPerformanceMetrics(self.fa).get_metrics()
        if self.request in ['all']:
            yield from HostConnectionsMetrics(self.fa).get_metrics()
