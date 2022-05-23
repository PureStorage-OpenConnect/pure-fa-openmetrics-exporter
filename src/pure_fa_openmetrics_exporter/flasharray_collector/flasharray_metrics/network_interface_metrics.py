from prometheus_client.core import GaugeMetricFamily

performance_eth_kpis = ['other_errors_per_sec',
                        'received_bytes_per_sec',
                        'received_crc_errors_per_sec',
                        'received_frame_errors_per_sec',
                        'received_packets_per_sec',
                        'total_errors_per_sec',
                        'transmitted_bytes_per_sec',
                        'transmitted_dropped_errors_per_sec',
                        'transmitted_packets_per_sec']

performance_fc_kpis = ['received_bytes_per_sec',
                       'received_crc_errors_per_sec',
                       'received_frames_per_sec',
                       'received_link_failures_per_sec',
                       'received_loss_of_signal_per_sec',
                       'received_loss_of_sync_per_sec',
                       'total_errors_per_sec',
                       'transmitted_bytes_per_sec',
                       'transmitted_frames_per_sec',
                       'transmitted_invalid_words_per_sec']

class NetworkInterfacePerformanceMetrics():
    """
    Base class for FlashArray Prometheus network interface performance metrics
    """

    def __init__(self, fa_client):
        self.interfaces = fa_client.network_interfaces_performance()
        self.performance = GaugeMetricFamily(
                               'purefa_network_interface_performance',
                               'FlashArray network interface performance',
                               labels = ['name', 'type', 'dimension'])

    def _build_metrics(self):
        cnt_i = 0
        for i in self.interfaces:
            i_type = i.interface_type
            if i_type == 'eth':
                for k in performance_eth_kpis:
                    v = getattr(i.eth, k)
                    if v is None:
                        continue
                    cnt_i += 1
                    self.performance.add_metric([i.name, i_type, k], v)
            if i_type == 'fc':
                for k in performance_fc_kpis:
                    v = getattr(i.fc, k)
                    if v is None:
                        continue
                    cnt_i += 1
                    self.performance.add_metric([i.name, i_type, k], v)
        if cnt_i == 0:
            self.performance = None

    def get_metrics(self):
        self._build_metrics()
        yield self.performance
