import pytest
from purity_fa_prometheus_exporter.flasharray_collector.flasharray_metrics import array_hardware_metrics

def test_array_hw_name(mock_fa_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(mock_fa_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert (s.name in ['purefa_hardware_chassis_health',
                               'purefa_hardware_controller_health',
                               'purefa_hardware_component_health',
                               'purefa_hardware_power_volts',
                               'purefa_hardware_temperature_celsius'])

def test_array_hw_labels(mock_fa_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(mock_fa_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            if s.name == 'purefa_hardware_chassis_health':
                assert len(s.labels['chassis']) > 0
            if s.name == 'purefa_hardware_controller_health':
                assert len(s.labels['controller']) > 0
            if s.name == 'purefa_hardware_component_health':
                assert (len(s.labels['controller']) > 0 or
                        len(s.labels['chassis']) > 0)
                assert len(s.labels['component']) > 0
                assert len(s.labels['index']) > 0
            if s.name == 'purefa_hardware_power_volts':
                assert len(s.labels['chassis']) > 0
                assert len(s.labels['power_supply']) > 0

def test_array_hw_val(mock_fa_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(mock_fa_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
