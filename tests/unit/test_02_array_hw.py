import pytest
from pure_fa_openmetrics_exporter.flasharray_collector.flasharray_metrics import array_hardware_metrics

def test_array_hw_name(fa_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(fa_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert (s.name in ['purefa_hardware_health',
                               'purefa_hardware_component_health',
                               'purefa_hardware_power_volts',
                               'purefa_hardware_temperature_celsius'])

def test_array_hw_labels(fa_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(fa_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            assert len(s.labels['type']) > 0

def test_array_hw_val(fa_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(fa_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
