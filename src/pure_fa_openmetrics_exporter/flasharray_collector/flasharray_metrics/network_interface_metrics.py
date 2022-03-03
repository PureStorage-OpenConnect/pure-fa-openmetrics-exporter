from prometheus_client.core import GaugeMetricFamily

class NetworkInterfacePerformanceMetrics():
    """
    Base class for FlashArray Prometheus network interface performance metrics
    """

    def __init__(self, fa_client):
        self.performance = None
        self.interfaces = fa_client.network_interfaces_performance()


    def _performance(self):
        self.performance = GaugeMetricFamily(
                               'purefa_network_interface_performance',
                               'FlashArray network interface performance',
                               labels = ['name', 'type', 'dimension'])
        for i in self.interfaces:
            i_type = i.interface_type
            if i_type == 'eth':
                self.performance.add_metric([i.name, i_type, 
                                            'other_errors_per_sec'],
                                            i.eth.other_errors_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_bytes_per_sec'],
                                            i.eth.received_bytes_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_crc_errors_per_sec'],
                                            i.eth.received_crc_errors_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_frame_errors_per_sec'],
                                            i.eth.received_frame_errors_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_packets_per_sec'],
                                            i.eth.received_packets_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'total_errors_per_sec'],
                                            i.eth.total_errors_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'transmitted_bytes_per_sec'],
                                            i.eth.transmitted_bytes_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'transmitted_dropped_errors_per_sec'],
                                            i.eth.transmitted_dropped_errors_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'transmitted_packets_per_sec'],
                                            i.eth.transmitted_packets_per_sec or 0)
            if i_type == 'fc':
                self.performance.add_metric([i.name, i_type, 
                                            'received_bytes_per_sec'],
                                            i.fc.received_bytes_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_crc_errors_per_sec'],
                                            i.fc.received_crc_errors_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_frames_per_sec'],
                                            i.fc.received_frames_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_link_failures_per_sec'],
                                            i.fc.received_link_failures_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_loss_of_signal_per_sec'],
                                            i.fc.received_loss_of_signal_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'received_loss_of_sync_per_sec'],
                                            i.fc.received_loss_of_sync_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'total_errors_per_sec'],
                                            i.fc.total_errors_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'transmitted_bytes_per_sec'],
                                            i.fc.transmitted_bytes_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'transmitted_frames_per_sec'],
                                            i.fc.transmitted_frames_per_sec or 0)
                self.performance.add_metric([i.name, i_type, 
                                            'transmitted_invalid_words_per_sec'],
                                            i.fc.transmitted_invalid_words_per_sec or 0)

    def get_metrics(self):
        self._performance()
        yield self.performance
