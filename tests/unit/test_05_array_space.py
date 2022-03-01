from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import array_space_metrics

def test_array_space_name(fa_client):
    array_space = array_space_metrics.ArraySpaceMetrics(fa_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_array_space_data_reduction_ratio',
                              'purefa_array_space_capacity_bytes',
                              'purefa_array_space_provisioned_bytes',
                              'purefa_array_space_used_bytes']

def test_array_space_labels(fa_client):
    array_space = array_space_metrics.ArraySpaceMetrics(fa_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            if s.name in ['purefa_array_space_data_reduction_ratio',
                          'purefa_array_space_capacity_bytes',
                          'purefa_array_space_provisioned_bytes'] :
                assert s.labels == {}
            if s.name == 'purefa_array_space_used_bytes':
                assert len(s.labels['space']) > 0
                assert s.labels['space'] in ['shared',
                                             'snapshots',
                                             'system',
                                             'total_physical',
                                             'unique',
                                             'virtual']

def test_array_space_val(fa_client):
    array_space = array_space_metrics.ArraySpaceMetrics(fa_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
