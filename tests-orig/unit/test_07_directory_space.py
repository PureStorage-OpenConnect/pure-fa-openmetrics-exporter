import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import directory_space_metrics

def test_directory_space_name(mock_fa_client):
    dir_space = directory_space_metrics.DirectorySpaceMetrics(mock_fa_client)
    for m in dir_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_directory_space_datareduction_ratio',
                              'purefa_directory_space_size_bytes',
                              'purefa_directory_space_used_bytes']

def test_array_space_labels(mock_fa_client):
    dir_space = directory_space_metrics.DirectorySpaceMetrics(mock_fa_client)
    for m in dir_space.get_metrics():
        for s in m.samples:
            if s.name in ['purefa_directory_space_datareduction_ratio'
                          'purefa_directory_space_size_bytes',
                          'purefa_directory_space_used_bytes']:
                assert len(s.labels['name']) > 0
                assert len(s.labels['filesystem']) > 0
                assert len(s.labels['path']) > 0
            if s.name == 'purefa_directory_space_used_bytes':
                assert len(s.labels['space']) > 0
                assert s.labels['space'] in ['shared',
                                             'snapshots',
                                             'system',
                                             'thin_provisioning',
                                             'total_physical',
                                             'total_provisioned',
                                             'total_reduction',
                                             'unique']

def test_directory_space_val(mock_fa_client):
    dir_space = directory_space_metrics.DirectorySpaceMetrics(mock_fa_client)
    for m in dir_space.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
