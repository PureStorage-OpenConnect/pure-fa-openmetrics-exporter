from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import array_events_metrics

def test_array_events_name(fa_client):
    array_ev = array_events_metrics.ArrayEventsMetrics(fa_client)
    for m in array_ev.get_metrics():
        for s in m.samples:
            assert s.name == 'purefa_alerts_open'

def test_array_events_labels(fa_client):
    array_ev = array_events_metrics.ArrayEventsMetrics(fa_client)
    for m in array_ev.get_metrics():
        for s in m.samples:
            assert len(s.labels['severity']) > 0
            assert len(s.labels['component_type']) > 0
            assert len(s.labels['component_name']) > 0
