![Current version](https://img.shields.io/badge/current%20version-0.0.1-blue) ![Dependencies](https://img.shields.io/badge/dependecies-python3--prometheus--client%20-orange)

# Pure Storage OpenMetrics exporter for FlashArray
OpenMetrics exporter for Pure Storage FlashArray.

## Overview

This applications aims to help monitor Pure Storage FlashArrays whith Prometheus, providing an "exporter" that is fully embeddable into Purity. The exporter is a Flask + Gunicorn application (100% Python 3 code) that extracts the performance, size and health status indicators from the FlashArray it runs onto and converts it the format Prometeus metrics) which is easily digested by Prometheus. This package is tested on Purity 6.2.


## Build

This package is mainly meant to generate the .deb package to be installed on Purity, although it is possible to merely generate the Python package via tox or pybuild.
The best and simples option to build the installable package is to use a docker image specifically meant to build Ubuntu packages. The source code to generate  that image is provided locally to this Pure Storage OpenConnect project in this repository [docker-deb-builder](https://github.com/PureStorage-OpenConnect/docker-deb-builder), which is a sighlty modified fork of a public GitHub repository.
In this way, all what you need to have installed on your build machine is docker.

Start by cloning the <kbd>deb-builder</kbd> repo and then generate the docker build image
```shell
git clone https://github.com/PureStorage-OpenConnect/docker-deb-builder.git
cd docker-deb-builder
docker build -t docker-deb-builder:20.04 -f Dockerfile-ubuntu-20.04 .
```
You have now a <kbd>docker-deb-builder:20.04</kbd> that can be used for building the final deb package.

Download or clone this repository and build the package. In the following example I am assuming your initial working directory is your homedir.

```shell
git clone git@github.com:PureStorage-OpenConnect/purity-fa-prometheus-exporter.git
mkdir output
cd docker-deb-builder
./build -i docker-deb-builder:20.04 -o ~/output ~/purity-fa-prometheus-exporter
```
The resulting <kbd>purity-fa-prometheus-exporter_\<version\>-\<deb-version\>_all.deb</kbd> package can be retrived in the <kbd>~/output</kbd> directory.

## Install
Once you have the package built, copy it over your target FlashArray.
Installation requires a running Purity 6.2 environment or a test environment that provides the same Python 3 deb packages. In addition to that, there is another Python packages the exporter depend on

- python3-prometheus-client, which is not installed in Purity but is a standard Ubuntu package available also in the Purity repository


Enable the pure repository by uncommenting the first four entries in the <kbd>/etc/apt/sources.list</kbd> configuration file. 
The sources.list file on Purity 6.2.x should look like this

```shell
## Uncomment the following lines to enable
## Purity apt repository
deb [arch=amd64] http://c7-apt-svc.dev.purestorage.com/focal/focal-20210407 focal main universe multiverse
deb [arch=amd64] http://c7-apt-svc.dev.purestorage.com/focal/focal-20210407 focal-security main universe multiverse
deb [arch=amd64] http://c7-apt-svc.dev.purestorage.com/focal/focal-20210407 focal-updates main universe multiverse
deb [arch=amd64] http://c7-apt-svc.dev.purestorage.com/focal/focal-20210407 focal main/debian-installer
# deb [arch=amd64] http://c7-apt-svc.dev.purestorage.com/hm-8251a77713ba220a4d0e7a64748540554acf30f5 focal main                   # openssh
# deb [arch=amd64] http://c7-apt-svc.dev.purestorage.com/hm-8251a77713ba220a4d0e7a64748540554acf30f5 focal main/debian-installer  # openssh

```
**Note**
You need to be connected to the Pure internal network in order to be able to install the missing Python Prometheus client package this way. As an alternative you can download the equivalent Ubuntu 20.04 package from the Canonical public repository, since the two packages are actually the same.
  
Install the Python Prometheus client package.

```shell
apt cache update
apt-get install python3-prometheus-client
```

Build the packages, then copy those to the target FlashArray and install with <kbd>dpkg</kbd>.

```shell
dpkg -i purity-prometheus-exporter_<version>-<deb-version>_all.deb
```
  
## Run
  
At this stage the installer intentionally does not enable neither starts the exporter, therefore it is necessary to manually enable and start it after the installation completes.
  
```shell
systemctl start purity-fa-prometheus-exporter
systemctl enable purity-fa-prometheus-exporter
```
In addition to that, the Purity nginx reverse proxy configuration must be reloaded.
```shell
systemctl reload nginx
```

and the Purity firewall needs to be restarted

```shell
firewall -d; firewall -e
```

### To do

At this stage the packaged replaces the original firewall.py and adds the necessary nginx configuration file to enable the access to the exporte, but this is a sub-optimal behavior. The missing step to fully integrate the Prometheus exporter into Purity would merely consist on changing the upstream <kbd>firewall.py</kbd> to include the Prometheus IP port. The nginx configuration file should also be converted into the template form defined for the similar Purity nginx services.

### Scraping endpoints

The exporter uses a RESTful API schema to provide Prometheus scraping endpoints.


URL | Description
---|---
http://\<fa-mgnt-endpoint\>:\<port\>/metrics | Full array metrics
http://\<fa-mgnt-endpoint\>:\<port\>/metrics/array | Array only metrics
http://\<fa-mgnt-endpoint\>:\<port\>/metrics/volumes | Volumes only metrics
http://\<fa-mgnt-endpoint\>:\<port\>/metrics/hosts | Hosts only metrics
http://\<fa-mgnt-endpoint\>:\<port\>/metrics/pods | Pods only metrics
http://\<fa-mgnt-endpoint\>:\<port\>/metrics/directories| Directories only metrics

Depending on the target array, scraping for the whole set of metrics could result into timeout issues, in which case it is suggested either to increase the scraping timeout or to scrape each single endpoint instead.

  
### Prometheus configuration examples

The [exemples](examples) directory provides an example of deployment of a Prometheus Grafana stack on k8s that can be used as the starting point to build your own solution.


## Copyright


