import re
from prometheus_client.core import GaugeMetricFamily

class NetworkInterfacePerformanceMetrics():
    """
    Base class for FlashArray Prometheus network interface performance metrics
    """

    def __init__(self, fa):
        self.fa = fa

        self.performance = GaugeMetricFamily(
                               'purefa_network_interface_performance',
                               'FlashArray network interface performance',
                               labels = ['name',
                                         'controller',
                                         'interface',
                                         'index',
                                         'dimension'])

    def _performance(self) -> None:
        """
        Create array network interface metrics of gauge type.
        """
        p = re.compile("^CT([0-9]+)\\.(ETH|FC)([0-9]+)$")
        for n in self.fa.get_network_interfaces():
            i_type = n['interface_type']
            if i_type == 'eth':
                for k in ['other_errors_per_sec',
                          'received_bytes_per_sec',
                          'received_crc_errors_per_sec',
                          'received_frame_errors_per_sec',
                          'received_packets_per_sec',
                          'total_errors_per_sec',
                          'transmitted_bytes_per_sec',
                          'transmitted_carrier_errors_per_sec',
                          'transmitted_dropped_errors_per_sec',
                          'transmitted_packets_per_sec']: 
                    val = n[i_type][k] 
                    val = val if val is not None else 0
                    name = n['name'].upper()
                    m = re.match(p, name)
                    self.performance.add_metric([name, m.group(1),
                                                i_type, m.group(3), k], val)
            if i_type == 'fc':
                for k in ['received_bytes_per_sec',
                          'received_crc_errors_per_sec',
                          'received_frames_per_sec',
                          'received_link_failures_per_sec',
                          'received_loss_of_signal_per_sec',
                          'received_loss_of_sync_per_sec',
                          'total_errors_per_sec',
                          'transmitted_bytes_per_sec',
                          'transmitted_frames_per_sec',
                          'transmitted_invalid_words_per_sec']:
                    val = n[i_type][k] 
                    val = val if val is not None else 0
                    name = n['name'].upper()
                    m = re.match(p, name)
                    self.performance.add_metric([name, m.group(1),
                                                i_type, m.group(3), k], val)

    def get_metrics(self) -> None:
        self._performance()
        yield self.performance
