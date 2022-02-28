import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import array_events_metrics

def test_array_events_name(mock_fa_client):
    array_ev = array_events_metrics.ArrayEventsMetrics(mock_fa_client)
    for m in array_ev.get_metrics():
        for s in m.samples:
            assert s.name == 'purefa_alerts_open'

def test_array_events_labels(mock_fa_client):
    array_ev = array_events_metrics.ArrayEventsMetrics(mock_fa_client)
    for m in array_ev.get_metrics():
        for s in m.samples:
            assert len(s.labels['severity']) > 0
            assert len(s.labels['component']) > 0
            assert len(s.labels['name']) > 0
