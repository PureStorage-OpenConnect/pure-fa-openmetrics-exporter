import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import volume_performance_metrics

def test_vol_perf_name(mock_fa_client):
    vol_perf = volume_performance_metrics.VolumePerformanceMetrics(mock_fa_client)
    for m in vol_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_volume_performance_latency_usec',
                              'purefa_volume_performance_bandwidth_bytes',
                              'purefa_volume_performance_iops',
                              'purefa_volume_performance_avg_block_bytes']

def test_host_perf_labels(mock_fa_client):
    vol_perf = volume_performance_metrics.VolumePerformanceMetrics(mock_fa_client)
    for m in vol_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['naaid']) > 0
            assert s.labels['pod'] is not None
            assert s.labels['vgroup'] is not None
            assert len(s.labels['dimension']) > 0
            if s.name == 'purefa_volume_performance_latency_usec':
                assert s.labels['dimension'] in ['queue_usec_per_mirrored_write_op',
                                                 'queue_usec_per_read_op',
                                                 'queue_usec_per_write_op',
                                                 'san_usec_per_mirrored_write_op',
                                                 'san_usec_per_read_op',
                                                 'san_usec_per_write_op',
                                                 'service_usec_per_mirrored_write_op',
                                                 'service_usec_per_read_op',
                                                 'service_usec_per_read_op_cache_reduction',
                                                 'service_usec_per_write_op',
                                                 'usec_per_mirrored_write_op',
                                                 'usec_per_read_op',
                                                 'usec_per_write_op']
            if s.name == 'purefa_volume_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['read_bytes_per_sec',
                                                 'write_bytes_per_sec',
                                                 'mirrored_write_bytes_per_sec']
            if s.name == 'purefa_volume_performance_iops':
                assert s.labels['dimension'] in ['reads_per_sec',
                                                 'writes_per_sec',
                                                 'mirrored_writes_per_sec']
            if s.name == 'purefa_volume_performance_avg_block_bytes':
                assert s.labels['dimension'] in ['bytes_per_read',
                                                 'bytes_per_write',
                                                 'bytes_per_op',
                                                 'bytes_per_mirrored_write']

def test_vol_perf_val(mock_fa_client):
    vol_perf = volume_performance_metrics.VolumePerformanceMetrics(mock_fa_client)
    for m in vol_perf.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
