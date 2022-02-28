import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import host_space_metrics

def test_host_space_name(mock_fa_client):
    host_space = host_space_metrics.HostSpaceMetrics(mock_fa_client)
    for m in host_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_host_space_datareduction_ratio',
                              'purefa_host_space_size_bytes',
                              'purefa_host_space_used_bytes']

def test_host_space_labels(mock_fa_client):
    host_space = host_space_metrics.HostSpaceMetrics(mock_fa_client)
    for m in host_space.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert s.labels['hostgroup'] is not None
            if s.name == 'purefa_host_space_used_bytes':
                assert s.labels['space'] in ['shared',
                                             'snapshots',
                                             'system',
                                             'thin_provisioning',
                                             'total_physical',
                                             'total_provisioned',
                                             'total_reduction',
                                             'unique']

def test_host_space_val(mock_fa_client):
    host_space = host_space_metrics.HostSpaceMetrics(mock_fa_client)
    for m in host_space.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
