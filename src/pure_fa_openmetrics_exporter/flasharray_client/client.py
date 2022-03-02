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
        self._network_interfaces_performance = None
        if self._disable_ssl_warn:
            urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
        self.client = flasharray.Client(target=target, 
                                        api_token=api_token,
                                        user_agent='Pure_FA_OpenMetrics_exporter/0.8')


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

    def arrays(self):
        if self._arrays:
            return list(self._arrays.values())
        try:
            res = self.client.get_arrays()
            adict = {}
            if isinstance(res, flasharray.ValidResponse):
                for a in res.items:
                    arr = {}
                    arr['array'] = a
                    arr['performance'] = {}
                    adict[a.id] = arr
            for p in ['all', 'nfs', 'smb']:
                res = self.client.get_arrays_performance()
                if isinstance(res, flasharray.ValidResponse):
                    for a in res.items:
                        adict[a.id]['performance'][p] = a
            self._arrays = adict
        except:
            pass
        return list(self._arrays.values())

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
            return list(self._directories.values())
        try:
            res = self.client.get_directories()
            ddict = {}
            if isinstance(res, flasharray.ValidResponse):
                for d in res.items:
                    dir = {}
                    dir['directory'] = d
                    dir['performance'] = {}
                    ddict[d.id] = dir
            res = self.client.get_directories_performance()
            if isinstance(res, flasharray.ValidResponse):
                for d in res.items:
                    ddict[d.id]['performance'] = d
            self._directories = ddict
        except:
             pass
        return list(self._directories.values())
