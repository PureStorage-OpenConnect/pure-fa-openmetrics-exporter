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

The exporter is full Python application based on [Sanic](https://sanic.dev/) framework. It is preferably built and launched via Docker. You can also scale the exporter deployment to multiple containers on Kubernetes thanks to the stateless nature of the application.

---

#### The official docker images are available at Quay.io

```shell
docker pull quay.io/purestorage/pure-fa-ome:<release>
```

where the release tag follows the semantic versioning.

---

### Local development
If you want to contribute to the development or simply build the package locally you should use python virtualenv

The following commands describe how to run a typical build using the basic virtualenv package:
```shell

python -m venv pure-fa-ome-build
source ./pure-fa-ome-build/bin/activate

# install dependencies
python -m pip install --upgrade pip
pip install build

# optionally install pytest and sanic_testing if you want to run tests
pip install pytest
pip install sanic_testing

# clone the repository
git clone git@github.com:PureStorage-OpenConnect/pure-fa-openmetrics-exporter.git

# modify the code and build the package
cd pure-fa-openmetrics-exporter
...
python -m build

```

The newly built package can be found in the <kbd>./dist</kbd> directory.

Tests can be run using tox. You need a running FlashArray you can use for executing the tests. Modify the tests/conftest.py file accordingly with your FlashArray management endpoint and API token.

### Docker image

The provided dockerfile can be used to generate a docker image of the exporter. It accepts the version of the python package as the build parameter, therefore you can build the image using docker as follows

```shell

VERSION=<version>
docker build --build-arg exporter_version=$VERSION -t pure-fa-ome:$VERSION .
```

### Scraping endpoints

The exporter uses a RESTful schema to provide scraping endpoints to OpenMetrics clients.

**Authentication**

Authentication is used by the exporter as the mechanism to cross authenticate to the scraped appliance, therefore for each array it is required to provide the REST API token for an account that has a 'readonly' role. The api-token must be provided in the http request using the HTTP Authorization header of type 'Bearer'. In case Prometheus is used, this is achieved by specifying the api-token value as the authorization parameter of the specific job in the Prometheus configuration file.


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

The [examples](examples) directory provides an example of deployment of a Prometheus Grafana stack on k8s that can be used as the starting point to build your own solution.
