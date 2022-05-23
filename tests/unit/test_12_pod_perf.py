from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import pod_performance_metrics

def test_pod_perf_name(fa_client):
    pod_perf = pod_performance_metrics.PodPerformanceMetrics(fa_client)
    for m in pod_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_pod_performance_latency_usec',
                              'purefa_pod_performance_bandwidth_bytes',
                              'purefa_pod_performance_iops',
                              'purefa_pod_performance_avg_block_bytes']

def test_host_perf_labels(fa_client):
    pod_perf = pod_performance_metrics.PodPerformanceMetrics(fa_client)
    for m in pod_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['dimension']) > 0
            if s.name == 'purefa_pod_performance_latency_usec':
                assert s.labels['dimension'] in ['queue_usec_per_mirrored_write_op',
                                                 'queue_usec_per_read_op',
                                                 'queue_usec_per_write_op',
                                                 'san_usec_per_mirrored_write_op',
                                                 'san_usec_per_read_op',
                                                 'san_usec_per_write_op',
                                                 'service_usec_per_mirrored_write_op',
                                                 'service_usec_per_read_op',
                                                 'service_usec_per_write_op',
                                                 'usec_per_mirrored_write_op',
                                                 'usec_per_read_op',
                                                 'usec_per_write_op']
            if s.name == 'purefa_pod_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['read_bytes_per_sec',
                                                 'write_bytes_per_sec',
                                                 'mirrored_write_bytes_per_sec']
            if s.name == 'purefa_pod_performance_iops':
                assert s.labels['dimension'] in ['reads_per_sec',
                                                 'writes_per_sec',
                                                 'mirrored_writes_per_sec']
            if s.name == 'purefa_pod_performance_avg_block_bytes':
                assert s.labels['dimension'] in ['bytes_per_read',
                                                 'bytes_per_write',
                                                 'bytes_per_op',
                                                 'bytes_per_mirrored_write']

def test_pod_perf_val(fa_client):
    pod_perf = pod_performance_metrics.PodPerformanceMetrics(fa_client)
    for m in pod_perf.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
