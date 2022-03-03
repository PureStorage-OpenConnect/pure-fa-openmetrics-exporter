from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import pod_space_metrics

def test_pod_space_name(fa_client):
    pod_space = pod_space_metrics.PodSpaceMetrics(fa_client)
    for m in pod_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_pod_space_data_reduction_ratio',
                              'purefa_pod_space_size_bytes',
                              'purefa_pod_space_used_bytes']

def test_pod_space_labels(fa_client):
    pod_space = pod_space_metrics.PodSpaceMetrics(fa_client)
    for m in pod_space.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            if s.name == 'purefa_pod_space_used_bytes':
                assert s.labels['space'] in ['replication',
                                             'shared',
                                             'snapshots',
                                             'total_physical',
                                             'total_provisioned',
                                             'unique']

def test_pod_space_val(fa_client):
    pod_space = pod_space_metrics.PodSpaceMetrics(fa_client)
    for m in pod_space.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
