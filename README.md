![Current version](https://img.shields.io/github/v/tag/PureStorage-OpenConnect/pure-fa-openmetrics-exporter?label=current%20version)

# Pure Storage FlashArray OpenMetrics exporter
OpenMetrics exporter for Pure Storage FlashArray.

## Support Statement
This exporter is provided under Best Efforts support by the Pure Portfolio Solutions Group, Open Source Integrations team.
For feature requests and bugs please use GitHub Issues.
We will address these as soon as we can, but there are no specific SLAs.
##

### Overview

This application aims to help monitor Pure Storage FlashArrays by providing an "exporter", which means it extracts data from the Purity API and converts it to the OpenMetrics format, which is for instance consumable by Prometheus.

The stateless design of the exporter allows for easy configuration management as well as scalability for a whole fleet of Pure Storage systems. Each time the OpenMetrics client scrapes metrics for a specific system, it should provide the hostname via GET parameter and the API token as Authorization token to this exporter.

To monitor your Pure Storage appliances, you will need to create a new dedicated user on your array, and assign read-only permissions to it. Afterwards, you also have to create a new API key.


### Building and Deploying

The exporter is a Go application based on the Prometheus Go client library and [Resty](https://github.com/go-resty/resty), a simple but reliable HTTP and REST client library for Go . It is preferably built and launched via Docker. You can also scale the exporter deployment to multiple containers on Kubernetes thanks to the stateless nature of the application.

---

#### The official docker images are available at Quay.io

```shell
docker pull quay.io/purestorage/pure-fa-om-exporter:<release>
```

where the release tag follows the semantic versioning.

---

#### Binaries

Binary downloads of the exporter can be found on the [Releases] page (https://github.com/PureStorage-OpenConnect/pure-fa-openmetrics-exporter/releases/latest).

---
### Local development

The following commands describe how to run a typical build :
```shell

# clone the repository
git clone git@github.com:PureStorage-OpenConnect/pure-fa-openmetrics-exporter.git

# modify the code and build the package
cd pure-fa-openmetrics-exporter
...
make build

```

The newly built exporter executable can be found in the <kbd>./out/bin</kbd> directory.

### Docker image

The provided dockerfile can be used to generate a docker image of the exporter. The image can be built using docker as follows

```shell

VERSION=<version>
docker build -t pure-fa-ome:$VERSION .
```

### Scraping endpoints

The exporter uses a RESTful API schema to provide Prometheus scraping endpoints.

**Authentication**

Authentication is used by the exporter as the mechanism to cross authenticate to the scraped appliance, therefore for each array it is required to provide the REST API token for an account that has a 'readonly' role. The api-token must be provided in the http request using the HTTP Authorization header of type 'Bearer'. This is achieved by specifying the api-token value as the authorization parameter of the specific job in the Prometheus configuration file.

### Scraping endpoints

The exporter uses a RESTful API schema to provide Prometheus scraping endpoints.


URL | GET parameters | Description
---|---|---
http://\<exporter-host\>:\<port\>/metrics | endpoint | Full array metrics
http://\<exporter-host\>:\<port\>/metrics/array | endpoint| Array only metrics
http://\<exporter-host\>:\<port\>/metrics/volumes | endpoint | Volumes only metrics
http://\<exporter-host\>:\<port\>/metrics/hosts | endpoint | Hosts only metrics
http://\<exporter-host\>:\<port\>/metrics/pods | endpoint | Pods only metrics
http://\<exporter-host\>:\<port\>/metrics/directories| endpoint | Directories only metrics

Depending on the target array, scraping for the whole set of metrics could result into timeout issues, in which case it is suggested either to increase the scraping timeout or to scrape each single endpoint instead.


### Usage examples

In a typical production scenario, it is recommended to use a visual frontend for your metrics, such as [Grafana](https://github.com/grafana/grafana). Grafana allows you to use your Prometheus instance as a datasource, and create Graphs and other visualizations from PromQL queries. Grafana, Prometheus, are all easy to run as docker containers.

To spin up a very basic set of those containers, use the following commands:
```bash
# Pure exporter
docker run -d -p 9490:9490 --name pure-fa-om-exporter quay.io/purestorage/pure-fa-om-exporter:<version>

# Prometheus with config via bind-volume (create config first!)
docker run -d -p 9090:9090 --name=prometheus -v /tmp/prometheus-pure.yml:/etc/prometheus/prometheus.yml -v /tmp/prometheus-data:/prometheus prom/prometheus:latest

# Grafana
docker run -d -p 3000:3000 --name=grafana -v /tmp/grafana-data:/var/lib/grafana grafana/grafana
```
Please have a look at the documentation of each image/application for adequate configuration examples.

A simple but complete example to deploy a full monitoring stack on kubernetes can be found in the [examples](examples/config/k8s) directory
