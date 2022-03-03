from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import host_connections_metrics

def test_host_conn_name(fa_client):
    h_conn = host_connections_metrics.HostConnectionsMetrics(fa_client)
    for m in h_conn.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_host_connections_info']

def test_host_conn_labels(fa_client):
    h_conn = host_connections_metrics.HostConnectionsMetrics(fa_client)
    for m in h_conn.get_metrics():
        for s in m.samples:
            assert len(s.labels['hostname']) > 0
            assert len(s.labels['naaid']) > 0

def test_host_conn_val(fa_client):
    h_conn = host_connections_metrics.HostConnectionsMetrics(fa_client)
    for m in h_conn.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
