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
        self._hardware = None
        self._volumes = None
        self._vgroups = None
        self._pods = None
        self._hosts = None
        self._connections = None
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
            adict = {}
            res = self.client.get_arrays()
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

    def volumes(self):
        if self._volumes:
            return list(self._volumes.values())
        try:
            vdict = {}
            res = self.client.get_volumes()
            if isinstance(res, flasharray.ValidResponse):
                for v in res.items:
                    vol = {}
                    vol['volume'] = v
                    vol['performance'] = {}
                    vdict[v.id] = vol
            res = self.client.get_volumes_performance()
            if isinstance(res, flasharray.ValidResponse):
                for v in res.items:
                    vdict[v.id]['performance'] = v
            self._volumes = vdict
        except:
            pass
        return list(self._volumes.values())

    def pods(self):
        if self._pods:
            return list(self._pods.values())
        try:
            pdict = {}
            res = self.client.get_pods()
            if isinstance(res, flasharray.ValidResponse):
                for p in res.items:
                    pod = {}
                    pod['pod'] = p
                    pod['performance'] = {}
                    pdict[p.id] = pod
            res = self.client.get_pods_performance()
            if isinstance(res, flasharray.ValidResponse):
                for p in res.items:
                    pdict[p.id]['performance'] = p
            self._pods = pdict
        except:
            pass
        return list(self._pods.values())

    def hosts(self):
        if self._hosts:
            return list(self._hosts.values())
        try:
            hdict = {}
            res = self.client.get_hosts()
            if isinstance(res, flasharray.ValidResponse):
                for h in res.items:
                    host = {}
                    host['host'] = h
                    host['performance'] = {}
                    hdict[h.name] = host
            res = self.client.get_hosts_performance()
            if isinstance(res, flasharray.ValidResponse):
                for h in res.items:
                    hdict[h.name]['performance'] = h
            self._hosts = hdict
        except:
            pass
        return list(self._hosts.values())

    def directories(self):
        if self._directories:
            return list(self._directories.values())
        try:
            ddict = {}
            res = self.client.get_directories()
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

    def connections(self):
        if self._connections is not None:
            return list(self._connections)
        try:
            self._connections = []
            res = self.client.get_connections()
            if isinstance(res, flasharray.ValidResponse):
                self.volumes()
                for hc in res.items:
                    hcdict = {}
                    hcdict['host'] = hc.host.name
                    hcdict['volume_serial'] = self._volumes[hc.volume.id]['volume'].serial
                    self._connections.append(hcdict)
        except Exception:
            pass
        return list(self._connections)

    def network_interfaces_performance(self):
        if self._network_interfaces_performance is not None:
            return self._network_interfaces_performance
        try:
            res = self.client.get_network_interfaces_performance()
            if isinstance(res, flasharray.ValidResponse):
                self._network_interfaces_performance = list(res.items)
        except Exception as e:
            self._network_interfaces = []
            pass
        return self._network_interfaces_performance

    def hardware(self):
        if self._hardware is not None:
            return self._hardware
        try:
            res = self.client.get_hardware()
            if isinstance(res, flasharray.ValidResponse):
                self._hardware = list(res.items)
        except Exception as e:
            self._hardware = []
            pass
        return self._hardware
