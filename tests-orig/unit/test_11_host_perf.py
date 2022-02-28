import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import host_performance_metrics

def test_host_perf_name(mock_fa_client):
    host_perf = host_performance_metrics.HostPerformanceMetrics(mock_fa_client)
    for m in host_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_host_performance_latency_usec',
                              'purefa_host_performance_bandwidth_bytes',
                              'purefa_host_performance_iops']

def test_host_perf_labels(mock_fa_client):
    host_perf = host_performance_metrics.HostPerformanceMetrics(mock_fa_client)
    for m in host_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['dimension']) > 0
            if s.name == 'purefa_host_performance_latency_usec':
                assert s.labels['dimension'] in ['queue_usec_per_read_op',
                                                 'queue_usec_per_write_op',
                                                 'queue_usec_per_mirrored_write_op',
                                                 'san_usec_per_read_op',
                                                 'san_usec_per_write_op',
                                                 'san_usec_per_mirrored_write_op',
                                                 'service_usec_per_mirrored_write_op',
                                                 'service_usec_per_read_op',
                                                 'service_usec_per_read_op_cache_reduction',
                                                 'service_usec_per_write_op',
                                                 'usec_per_read_op',
                                                 'usec_per_write_op',
                                                 'usec_per_mirrored_write_op']
            if s.name == 'purefa_host_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['read_bytes_per_sec',
                                                 'write_bytes_per_sec',
                                                 'mirrored_write_bytes_per_sec']
            if s.name == 'purefa_host_performance_iops':
                assert s.labels['dimension'] in ['reads_per_sec',
                                                 'writes_per_sec',
                                                 'mirrored_writes_per_sec']

def test_host_perf_val(mock_fa_client):
    host_perf = host_performance_metrics.HostPerformanceMetrics(mock_fa_client)
    for m in host_perf.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
