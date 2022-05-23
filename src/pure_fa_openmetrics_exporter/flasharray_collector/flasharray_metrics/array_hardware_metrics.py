from prometheus_client.core import GaugeMetricFamily

class ArrayHardwareMetrics:
    """
    Base class for FlashArray OpenMetrics hardware metrics
    """
    def __init__(self, fa_client):
        self.hardware = fa_client.hardware()
        self.hardware_status = GaugeMetricFamily(
                                  'purefa_hardware_health',
                                  'FlashArray hardware component health status',
                                   labels=['name', 'type'])

        self.temperature = GaugeMetricFamily(
                               'purefa_hardware_temperature_celsius',
                               'FlashArray hardware temperature sensors',
                               labels=['name', 'type'])

        self.power = GaugeMetricFamily(
                         'purefa_hardware_power_volts',
                         'FlashArray hardware power supply voltage',
                         labels=['name', 'type'])

    def _build_metrics(self):
        cnt_s = 0
        cnt_t = 0
        cnt_p = 0
        for comp in self.hardware:
            if (comp.status == 'not_installed'):
                continue
            status = 1 if comp.status == 'ok' else 0
            self.hardware_status.add_metric([comp.name, comp.type], status)
            cnt_s += 1
            if comp.type == 'temp_sensor':
                self.temperature.add_metric([comp.name, comp.type], 
                                      float(comp.temperature))
                cnt_t += 1
            elif comp.type == 'power_supply':
                if comp.voltage is not None:
                    self.power.add_metric([comp.name, comp.type], 
                                          float(comp.voltage))
                    cnt_p += 1
        if cnt_s == 0:
            self.hardware_status = None
        if cnt_t == 0:
            self.temperature = None
        if cnt_p == 0:
            self.power = None

    def get_metrics(self):
        self._build_metrics()
        yield self.hardware_status
        yield self.temperature
        yield self.power
