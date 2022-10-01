package collectors


import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)

type HardwareCollector struct {
	StatusDesc       *prometheus.Desc
	TemperatureDesc  *prometheus.Desc
        VoltageDesc      *prometheus.Desc
	Client           *client.FAClient
}

func (c *HardwareCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *HardwareCollector) Collect(ch chan<- prometheus.Metric) {
	hwl := c.Client.GetHardware()
	if len(hwl.Items) == 0 {
		return
	}
	var s float64
	for _, h := range hwl.Items {
		if h.Status != "ok" {
			s = 1
		} else {
			s = 0
		}
		ch <- prometheus.MustNewConstMetric(
			c.StatusDesc,
			prometheus.GaugeValue,
			s,
			h.Name, h.Type,
		)
		if h.Temperature > 0 {
			ch <- prometheus.MustNewConstMetric(
				c.TemperatureDesc,
				prometheus.GaugeValue,
				float64(h.Temperature),
				h.Name, h.Type,
			)
		}
		if h.Voltage > 0 {
			ch <- prometheus.MustNewConstMetric(
				c.VoltageDesc,
				prometheus.GaugeValue,
				float64(h.Voltage),
				h.Name, h.Type,
			)
		}
	}
}

func NewHardwareCollector(fa *client.FAClient) *HardwareCollector {
	return &HardwareCollector{
		StatusDesc: prometheus.NewDesc(
			"purefa_hw_component_status",
			"FlashArray hardware component status",
			[]string{"component_name", "component_type"},
			prometheus.Labels{},
		),
		TemperatureDesc: prometheus.NewDesc(
			"purefa_hw_component_temperature_celsius",
			"FlashArray hardware component temperature",
			[]string{"component_name", "component_type"},
			prometheus.Labels{},
		),
		VoltageDesc: prometheus.NewDesc(
			"purefa_hw_component_voltage_volt",
			"FlashArray hardware component voltage",
			[]string{"component_name", "component_type"},
			prometheus.Labels{},
		),
		Client: fa,
	}
}
