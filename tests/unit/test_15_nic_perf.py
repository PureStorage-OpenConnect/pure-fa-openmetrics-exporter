from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import network_interface_metrics

def test_nic_perf_name(fa_client):
    nic_perf = network_interface_metrics.NetworkInterfacePerformanceMetrics(fa_client)
    for m in nic_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_network_interface_performance']

def test_nic_perf_labels(fa_client):
    nic_perf = network_interface_metrics.NetworkInterfacePerformanceMetrics(fa_client)
    for m in nic_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['type']) > 0
            assert len(s.labels['dimension']) > 0
            assert s.labels['type'] in ['eth', 'fc']
            if s.labels['type'] == 'eth':
                assert s.labels['dimension'] in ['other_errors_per_sec',
                                                 'received_bytes_per_sec',
                                                 'received_crc_errors_per_sec',
                                                 'received_frame_errors_per_sec',
                                                 'received_packets_per_sec',
                                                 'total_errors_per_sec',
                                                 'transmitted_bytes_per_sec',
                                                 'transmitted_dropped_errors_per_sec',
                                                 'transmitted_packets_per_sec']
            if s.labels['type'] == 'fc':
                assert s.labels['dimension'] in ['received_bytes_per_sec',
                                                 'received_crc_errors_per_sec',
                                                 'received_frames_per_sec',
                                                 'received_link_failures_per_sec',
                                                 'received_loss_of_signal_per_sec',
                                                 'received_loss_of_sync_per_sec',
                                                 'total_errors_per_sec',
                                                 'transmitted_bytes_per_sec',
                                                 'transmitted_frames_per_sec',
                                                 'transmitted_invalid_words_per_sec']

def test_nic_perf_val(fa_client):
    nic_perf = network_interface_metrics.NetworkInterfacePerformanceMetrics(fa_client)
    for m in nic_perf.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
