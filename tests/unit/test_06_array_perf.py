from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import array_performance_metrics

def test_array_performance_name(fa_client):
    array_perf = array_performance_metrics.ArrayPerformanceMetrics(fa_client)
    for m in array_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_array_performance_latency_usec',
                              'purefa_array_performance_iops',
                              'purefa_array_performance_bandwidth_bytes',
                              'purefa_array_performance_average_block_bytes']


def test_array_performance_labels(fa_client):
    array_perf = array_performance_metrics.ArrayPerformanceMetrics(fa_client)
    for m in array_perf.get_metrics():
        for s in m.samples:
            assert s.labels['protocol'] in ['all', 'nfs', 'smb']
            if s.name in ['purefa_array_performance_latency_usec',
                          'purefa_array_performance_iops',
                          'purefa_array_performance_bandwidth_bytes',
                          'purefa_array_performance_average_block_bytes']:
                assert len(s.labels['dimension']) > 0
            if s.name == 'purefa_array_performance_latency_usec':
                assert s.labels['dimension'] in ['local_queue_usec_per_op',
                                                 'queue_usec_per_read_op',
                                                 'queue_usec_per_write_op',
                                                 'queue_usec_per_mirrored_write_op',
                                                 'san_usec_per_read_op',
                                                 'san_usec_per_write_op',
                                                 'san_usec_per_mirrored_write_op',
                                                 'service_usec_per_mirrored_write_op',
                                                 'service_usec_per_read_op',
                                                 'service_usec_per_write_op',
                                                 'usec_per_read_op',
                                                 'usec_per_write_op',
                                                 'usec_per_mirrored_write_op',
                                                 'usec_per_other_op']
            if s.name == 'purefa_array_performance_iops':
                assert s.labels['dimension'] in ['reads_per_sec',
                                                 'writes_per_sec',
                                                 'mirrored_writes_per_sec',
                                                 'others_per_sec']
            if s.name == 'purefa_array_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['read_bytes_per_sec',
                                                 'write_bytes_per_sec',
                                                 'mirrored_write_bytes_per_sec']
            if s.name == 'purefa_array_performance_avg_block_bytes':
                assert s.labels['dimension'] in ['bytes_per_read',
                                                 'bytes_per_write',
                                                 'bytes_per_op']

def test_array_performance_val(fa_client):
    array_perf = array_performance_metrics.ArrayPerformanceMetrics(fa_client)
    for m in array_perf.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
