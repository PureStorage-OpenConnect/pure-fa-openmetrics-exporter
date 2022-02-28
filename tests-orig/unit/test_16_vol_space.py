import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import volume_space_metrics

def test_vol_space_name(mock_fa_client):
    vol_space = volume_space_metrics.VolumeSpaceMetrics(mock_fa_client)
    for m in vol_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_volume_space_datareduction_ratio',
                              'purefa_volume_space_size_bytes',
                              'purefa_volume_space_used_bytes']

def test_vol_space_labels(mock_fa_client):
    vol_space = volume_space_metrics.VolumeSpaceMetrics(mock_fa_client)
    for m in vol_space.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['naaid']) > 0
            assert s.labels['pod'] is not None
            assert s.labels['vgroup'] is not None
            if s.name == 'purefa_volume_space_used_bytes':
                assert s.labels['space'] is not None
                assert s.labels['space'] in ['shared',
                                             'snapshots',
                                             'snapshots_effective',
                                             'system',
                                             'thin_provisioning',
                                             'total_effective',
                                             'total_physical',
                                             'total_provisioned',
                                             'total_reduction',
                                             'unique',
                                             'virtual']

def test_vol_space_val(mock_fa_client):
    vol_space = volume_space_metrics.VolumeSpaceMetrics(mock_fa_client)
    for m in vol_space.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
