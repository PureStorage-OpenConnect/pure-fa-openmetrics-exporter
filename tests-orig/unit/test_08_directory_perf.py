import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import directory_performance_metrics

def test_directory_perf_name(mock_fa_client):
    dir_perf = directory_performance_metrics.DirectoryPerformanceMetrics(mock_fa_client)
    for m in dir_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_directory_performance_latency_usec',
                              'purefa_directory_performance_bandwidth_bytes',
                              'purefa_directory_performance_iops',
                              'purefa_directory_performance_avg_block_bytes']

def test_directory_space_labels(mock_fa_client):
    dir_perf = directory_performance_metrics.DirectoryPerformanceMetrics(mock_fa_client)
    for m in dir_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['filesystem']) > 0
            assert len(s.labels['path']) > 0
            assert len(s.labels['dimension']) > 0
            if s.name == 'purefa_directory_performance_latency_usec':
                assert s.labels['dimension'] in ['usec_per_read_op',
                                                 'usec_per_write_op',
                                                 'usec_per_other_op']
            if s.name == 'purefa_directory_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['write_bytes_per_sec',
                                                 'read_bytes_per_sec']
            if s.name == 'purefa_directory_performance_iops':
                assert s.labels['dimension'] in ['reads_per_sec',
                                                 'writes_per_sec',
                                                 'others_per_sec']
            if s.name == 'purefa_directory_performance_avg_block_bytes':
                assert s.labels['dimension'] in ['bytes_per_op',
                                                 'bytes_per_read',
                                                 'bytes_per_write']

def test_directory_perf_val(mock_fa_client):
    dir_perf = directory_performance_metrics.DirectoryPerformanceMetrics(mock_fa_client)
    for m in dir_perf.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
