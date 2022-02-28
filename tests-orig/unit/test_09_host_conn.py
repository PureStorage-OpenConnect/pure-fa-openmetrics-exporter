import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import host_connections_metrics

def test_host_conn_name(mock_fa_client):
    h_conn = host_connections_metrics.HostConnectionsMetrics(mock_fa_client)
    for m in h_conn.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_host_connections_info']

def test_host_conn_labels(mock_fa_client):
    h_conn = host_connections_metrics.HostConnectionsMetrics(mock_fa_client)
    for m in h_conn.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['naaid']) > 0

def test_host_conn_val(mock_fa_client):
    h_conn = host_connections_metrics.HostConnectionsMetrics(mock_fa_client)
    for m in h_conn.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
