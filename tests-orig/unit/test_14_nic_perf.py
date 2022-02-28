import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import network_interface_metrics

def test_nic_perf_name(mock_fa_client):
    nic_perf = network_interface_metrics.NetworkInterfacePerformanceMetrics(mock_fa_client)
    for m in nic_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefa_network_interface_performance']

def test_nic_perf_labels(mock_fa_client):
    nic_perf = network_interface_metrics.NetworkInterfacePerformanceMetrics(mock_fa_client)
    for m in nic_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['controller']) > 0
            assert len(s.labels['interface']) > 0
            assert len(s.labels['index']) > 0
            assert len(s.labels['dimension']) > 0

def test_nic_perf_val(mock_fa_client):
    nic_perf = network_interface_metrics.NetworkInterfacePerformanceMetrics(mock_fa_client)
    for m in nic_perf.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
            assert s.value >= 0
