from prometheus_client.core import GaugeMetricFamily

class ArrayHardwareMetrics:
    """
    Base class for FlashArray OpenMetrics hardware metrics
    """
    def __init__(self, fa_client):
        self.hardware_status = None
        self.temperature = None
        self.power = None
        self.hardware = fa_client.hardware()

    def _hardware(self):
        self.hardware_status = GaugeMetricFamily(
                                  'purefa_hardware_health',
                                  'FlashArray hardware component health status',
                                   labels=['name', 'type'])

        self.temperature = GaugeMetricFamily(
                               'purefa_hardware_temperature_celsius',
                               'FlashArray hardware temperature sensors',
                               labels=['name',
                                       'type'])

        self.power = GaugeMetricFamily(
                         'purefa_hardware_power_volts',
                         'FlashArray hardware power supply voltage',
                         labels=['name', 'type'])


        for comp in self.hardware:
            if (comp.status == 'not_installed'):
                continue
            status = 1 if comp.status == 'ok' else 0
            self.hardware_status.add_metric([comp.name, comp.type], status)

            if comp.type == 'temp_sensor':
                self.power.add_metric([comp.name, comp.type], 
                                      float(comp.temperature))
            elif comp.type == 'power_supply':
                if comp.voltage is not None:
                     self.power.add_metric([comp.name, comp.type], 
                                           float(comp.voltage))

    def get_metrics(self):
        self._hardware()
        yield self.hardware_status
        yield self.temperature
        yield self.power
