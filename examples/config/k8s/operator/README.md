# Deploying with the Prometheus Operator

Install Prometheus using the [Prometheus Operator](https://github.com/prometheus-operator/prometheus-operator), then deploy 
the Pure FlashArray Prometheus exporter using the deployment file [pure-fa-exporter-deployment.yaml](./pure-fa-exporter-deployment.yaml). Configure Prometheus using the Probe custom resouce defined in the [pure-fa-probe.yaml](./pure-fa-probe.yaml).
