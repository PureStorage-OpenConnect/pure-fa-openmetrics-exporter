from prometheus_client.core import GaugeMetricFamily

class ArrayHardwareMetrics:
    """
    Base class for FlashArray OpenMetrics hardware metrics
    """
    def __init__(self, fa_client):
        self.array_hardware_status = None
        self.array_hardware_temperature = None
        self.array_hardware_power = None
        self.hardware = fa_client.hardware()

        self.array_hardware_status = GaugeMetricFamily(
                                  'purefa_hardware_health',
                                  'FlashArray hardware component health status',
                                   labels=['name', 'chassis'])

        self.temperature = GaugeMetricFamily(
                               'purefa_hardware_temperature_celsius',
                               'FlashArray hardware temperature sensors',
                               labels=['name',
                                       'sensor',
                                       'index'])

        self.power = GaugeMetricFamily(
                         'purefa_hardware_power_volts',
                         'FlashArray hardware power supply voltage',
                         labels=['name', 'power_supply'])


        for comp in self.hardware_status():
            if (comp['status'] == 'not_installed'):
                continue
            c_name = comp['name']
            c_status = 1 if (comp['status'] == 'ok') else 0
            c_type = comp['type']
            c_index = str(comp['index'])


                if c_type == 'temp_sensor':
                            float(comp['temperature']))
                elif c_type == 'power_supply':
                    if comp['voltage'] is not None:
                        self.power.add_metric([c_name, c_base_index, c_index],
                                         float(comp['voltage']))

    def get_metrics(self):
        self._array_hardware_status()
        yield self.array_hardware_status()
        yield self.array_temperature
        yield self.array_power
