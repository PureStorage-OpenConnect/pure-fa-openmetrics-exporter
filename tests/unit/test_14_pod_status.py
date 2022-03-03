from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import pod_status_metrics

def test_pod_status_name(fa_client):
    pod_status = pod_status_metrics.PodStatusMetrics(fa_client)
    for m in pod_status.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_pod_status',
                              'purefa_pod_mediator_status']

def test_pod_status_labels(fa_client):
    pod_status = pod_status_metrics.PodStatusMetrics(fa_client)
    for m in pod_status.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['array_name']) > 0

def test_pod_status_val(fa_client):
    pod_status = pod_status_metrics.PodStatusMetrics(fa_client)
    for m in pod_status.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
