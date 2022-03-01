import urllib3
from pypureclient import flasharray, PureError

PURE_NAA = 'naa.624a9370'

class FlasharrayClient():
    """
    This is a simple wrapper to the Pure REST API 2.x specifically meant
    to optimize the "scraping" of space and performance metrics by Prometheus.
    Each endpoint is scraped only once and the result cached internally, so
    that any subsequent call does not actually query the endpoint and uses
    instead the internal result.
    """
    def __init__(self, target, api_token, disable_ssl_warn=False):
        self._disable_ssl_warn = disable_ssl_warn
        self._arrays = None
        self._arrays_performance = None
        self._arrays_space = None
        self._alerts = None
        self._volumes = None
        self._volumes_performance = None
        self._volumes_space = None
        self._vgroups = None
        self._pods = None
        self._pods_performance = None
        self._pods_space = None
        self._hosts = None
        self._hosts_performance = None
        self._hosts_space = None
        self._host_connections = None
        self._directories = None
        self._directories_performance = None
        self._directories_space = None
        self._network_interfaces_performance = None
        if self._disable_ssl_warn:
            urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
        self.client = flasharray.Client(target=target, 
                                        api_token=api_token,
                                        user_agent='Pure_FA_OpenMetrics_exporter/0.8')

    def arrays(self):
        if self._arrays:
            return self._arrays
        self._arrays = []
        try:
            res = self.client.get_arrays()
            if isinstance(res, flasharray.ValidResponse):
                self._arrays = list(res.items)
        except:
            pass
        return self._arrays

    def arrays_performance(self):
        if self._arrays_performance:
            return self._arrays_performance
        array_perf = {}
        try:
            for p in ['all', 'nfs', 'smb']:
                array_perf[p] = None
                res = self.client.get_arrays_performance(protocol=p)
                if not isinstance(res, flasharray.ValidResponse):
                    continue
                array_perf[p] = next(res.items)
        except:
            pass
        self._arrays_performance = array_perf
        return self._arrays_performance

    def arrays_space(self):
        if self._arrays_space:
            return self._arrays_space
        array_space = {}
        try:
            res = self.client.get_arrays_space()
            if isinstance(res, flasharray.ValidResponse):
                array_space = next(res.items)
        except:
            pass
        self._arrays_space = array_space
        return self._arrays_space

    def alerts(self):
        if self._alerts:
            return self._alerts
        alerts = {}
        try:
            res = self.client.get_alerts()
            if isinstance(res, flasharray.ValidResponse):
                alerts = list(res.items)
        except:
            pass
        self._alerts = alerts
        return self._alerts

    def volumes_performance(self):
        if self._volumes_performance:
            return self._volumes_performance
        vols_perf = []
        try:
            res = self.client.get_volumes_performance()
            if isinstance(res, flasharray.ValidResponse):
                vols_perf = list(res.items)
        except:
            pass
        self._volumes_performance = vols_perf
        return self._volumes_performance

    def volumes(self):
        if self._volumes:
            return self._volumes
        vols = []
        try:
            res = self.client.get_volumes()
            if isinstance(res, flasharray.ValidResponse):
                vols = list(res.items)
        except:
            pass
        self._volumes = vols
        return self._volumes

    def volumes_space(self):
        if self._volumes_space:
            return self._volumes_space
        vols_space = []
        try:
            res = self.client.get_volumes_space()
            if isinstance(res, flasharray.ValidResponse):
                vols_space = list(res.items)
        except:
            pass
        self._volumes_space = vols_space
        return self._volumes_space

    def pods(self):
        if self._pods:
            return self._pods
        pods = []
        try:
            res = self.client.get_pods()
            if isinstance(res, flasharray.ValidResponse):
                pods = list(res.items)
        except:
            pass
        self._pods = pods
        return self._pods

    def pods_performance(self):
        if self._pods_performance:
            return self._pods_performance
        pods_perf = []
        try:
            res = self.client.get_pods_performance()
            if isinstance(res, flasharray.ValidResponse):
                pods_perf = list(res.items)
        except:
            pass
        self._pods_performance = pods_perf
        return self._pods_performance

    def pods_space(self):
        if self._pods_space:
            return self._pods_space
        pods_space = []
        try:
            res = self.client.get_pods_space()
            if isinstance(res, flasharray.ValidResponse):
                pods_space = list(res.items)
        except:
            pass
        self._pods_space = pods_space
        return self._pods_space

    def hosts(self):
        if self._hosts:
            return self._hosts
        hosts = []
        try:
            res = self.client.get_hosts()
            if isinstance(res, flasharray.ValidResponse):
                hosts = list(res.items)
        except:
            pass
        self._hosts = hosts
        return self._hosts

    def hosts_performance(self):
        if self._hosts_performance:
            return self._hosts_performance
        hosts_perf = []
        try:
            res = self.client.get_hosts_performance()
            if isinstance(res, flasharray.ValidResponse):
                hosts_perf = list(res.items)
        except:
            pass
        self._hosts_performance = hosts_perf
        return self._hosts_performance

    def hosts_space(self):
        if self._hosts_space:
            return self._hosts_space
        hosts_space = []
        try:
            res = self.client.get_hosts_space()
            if isinstance(res, flasharray.ValidResponse):
                hosts_space = list(res.items)
        except:
            pass
        self._hosts_space = hosts_space
        return self._hosts_space

    def directories(self):
        if self._directories:
            return self._directories
        dirs = []
        try:
            res = self.client.get_directories()
            if isinstance(res, flasharray.ValidResponse):
                dirs = list(res.items)
        except:
            pass
        self._directories = dirs
        return self._directories

    def directories_performance(self):
        if self._directories_performance:
            return self._directories_performance
        dir_perf = []
        try:
            res = self.client.get_directories_performance()
            if isinstance(res, flasharray.ValidResponse):
                dir_perf = list(res.items)
        except:
            pass
        self._directories_performance = dir_perf
        return self._directories_performance

    def directories_space(self):
        if self._directories_space:
            return self._directories_space
        dir_space = []
        try:
            res = self.client.get_directories_space()
            if isinstance(res, flasharray.ValidResponse):
                dir_space = list(res.items)
        except:
            pass
        self._directories_space = dir_space
        return self._directories_space
