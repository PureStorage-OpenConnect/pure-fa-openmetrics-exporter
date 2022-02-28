import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import array_space_metrics

def test_array_space_name(mock_fa_client):
    array_space = array_space_metrics.ArraySpaceMetrics(mock_fa_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_array_space_datareduction_ratio',
                              'purefa_array_space_capacity_bytes',
                              'purefa_array_space_provisioned_bytes',
                              'purefa_array_space_used_bytes']

def test_array_space_labels(mock_fa_client):
    array_space = array_space_metrics.ArraySpaceMetrics(mock_fa_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            if s.name in ['purefa_array_space_datareduction_ratio',
                          'purefa_array_space_capacity_bytes',
                          'purefa_array_space_provisioned_bytes'] :
                assert s.labels == {}
            if s.name == 'purefa_array_space_used_bytes':
                assert len(s.labels['space']) > 0
                assert s.labels['space'] in ['shared',
                                             'snapshots',
                                             'system',
                                             'thin_provisioning',
                                             'total_physical',
                                             'unique',
                                             'virtual',
                                             'unique_effective',
                                             'snapshots_effective',
                                             'total_effective',
                                             'replication',
                                             'shared_effective']

def test_array_space_val(mock_fa_client):
    array_space = array_space_metrics.ArraySpaceMetrics(mock_fa_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
