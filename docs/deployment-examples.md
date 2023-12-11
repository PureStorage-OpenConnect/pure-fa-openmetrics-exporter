# Example Deployment Methods
There are a number of methods to deploy the OpenMetrics exporter. In this document we explore some of the common features and deployment methods by providing walk through examples.

## Contents
 - [Prerequisites](#prerequisites)
 - [Container Deployment](#container-deployment)
    - [Container - default - http passing API token with query](#container---default---http-passing-api-token-with-query)
    - [Container - Tokens File - http with API tokens file embedded in exporter](#container---tokens-file---http-with-api-tokens-file-embedded-in-exporter)
    - [Container - TLS - https passing API token with query](#container---tls---https-passing-api-token-with-query)
 - [Executable Binary Deployment](#executable-binary-deployment)
    - [Binary Specific Prerequisites](#binary-specific-prerequisites)
    - [Binary - default - http passing API token with query](#binary---default---http-passing-api-token-with-query)
    - [Binary - Tokens File - http with API tokens file embedded in exporter](#binary---tokens-file---http-with-api-tokens-file-embedded-in-exporter)
    - [Binary - TLS - https passing API token with query](#binary---tls---https-passing-api-token-with-query)

# Prerequisites
All deployments will require an API token to authenticate with the array. Read-only only user access is recommended.

Generate an API token from your chosen user account or create a new readonly user.
API token can be retrieved from either Purity GUI ot CLI.

<details>
<summary>Expand for CLI example</summary>

```console
pureuser@arrayname01> pureadmin create --role readonly o11y-readonly
Name           Type   Role    
o11y-readonly  local  readonly

pureuser@arrayname01> pureadmin create --api-token o11y-readonly
Name           Type   API Token                             Created                  Expires
o11y-readonly  local  11111111-1111-1111-1111-111111111111  2022-11-30 08:58:40 EST  -      
```

</details>

<details>
<summary>Expand for GUI example</summary>

![Alt text](../extra/images/purefa_create_api_token.png)
</details>

---

# Container Deployment
## Container - default - http passing API token with query
We build container images and publish them to RedHat quay.io. Here they can be continually scanned for vulnerabilities and updated as soon as a new version are available.
You can navigate to [https://quay.io/repository/purestorage/pure-fa-om-exporter](https://quay.io/repository/purestorage/pure-fa-om-exporter) to view the available versions. In this example we will pull a container down using Docker.

In this example we will use the default port 9490, set the name as pure-fa-om-exporter and pull the latest version of the image from quay.io/purestorage/pure-fa-om-exporter. The `--detach` switch sends the process to the background.

1. **Run and pull the image**
    ```console
    $ docker run --detach --publish 9490:9490 --name pure-fa-om-exporter --restart unless-stopped quay.io/purestorage/pure-fa-om-exporter:latest
    ```
    In this guide we will not be covering Docker troubleshooting.

2. **Check the container is running**
    ```console
    $ docker ps
    CONTAINER ID   IMAGE                                            COMMAND                  CREATED         STATUS         PORTS                                       NAMES
    0076ded8a073   quay.io/purestorage/pure-fa-om-exporter:latest   "/pure-fa-om-exporte…"   10 seconds ago   Up 10 seconds   0.0.0.0:9490->9490/tcp, :::9490->9490/tcp   pure-fa-om-exporter
    ```

3. **Test the exporter**
    Use `curl` to test the exporter returns results.
    ```console
    $ curl -H 'Authorization: Bearer 11111111-1111-1111-1111-111111111111' -X GET 'http://localhost:9490/metrics/array?endpoint=array01' -silent | grep ^purefa_info
    purefa_info{array_name="ARRAY01",os="Purity//FA",system_id="11111111-1111-1111-1111-111111111111",version="6.5.0"} 1
    ```

    We expect to return a single line displaying `purefa_info` returning the name, Purity OS, system ID and the Purity version.
    You can remove the `-silent | grep ^purefa_info` from the command to see a list of all results.

## Container - Tokens File - http with API tokens file embedded in exporter

1. Create a tokens.yaml file to pass to the exporter.
    ```console
    $ cat /directorypath/tokens.yaml
    # alias: of the array to be used in the exporter query
    array01:
      # FQDN or IP address of the array
      address: 'array01.fqdn.com'
      # API token of a user on the array
      api_token: '11111111-1111-1111-1111-111111111111'

    # Example of a second array
    array02:
      address: 'array02.fqdn.com'
      api_token: '22222222-2222-2222-2222-222222222222'
    ```

2. **Run the container**
    Run the container this time specifying a volume (in our case a single file) to pass the token.yaml file from the host server to the guest container.
    The container expects to see the file here: `/etc/pure-fa-om-exporter/tokens.yaml`

    ```console
    $ docker run -d -p 9490:9490 --name pure-fa-om-exporter --volume /directorypath/tokens.yaml:/etc/pure-fa-om-exporter/tokens.yaml quay.io/purestorage/pure-fa-om-exporter:latest
    ```
3. **Check the container is running**
    ```console
    $ docker ps
    CONTAINER ID   IMAGE                                            COMMAND                  CREATED         STATUS         PORTS                                       NAMES
    0076ded8a073   quay.io/purestorage/pure-fa-om-exporter:latest   "/pure-fa-om-exporte…"   10 seconds ago   Up 10 seconds   0.0.0.0:9490->9490/tcp, :::9490->9490/tcp   pure-fa-om-exporter
    ```

4. **Test the exporter**
    Use `curl` to test the exporter returns results. We don't need to pass the bearer (API) token for authorization as the exporter has a record of these which makes queries simpler.
    ```console
    $ curl -X GET 'http://localhost:9490/metrics/array?endpoint=array01 ' -silent | grep ^purefa_info
    purefa_info{array_name="ARRAY01",os="Purity//FA",system_id="11111111-1111-1111-1111-111111111111",version="6.5.0"} 1
    ```

## Container - TLS - https passing API token with query
TLS with Docker image is currently not supported. Please use executable binary only for this feature.

# Executable Binary Deployment
## Binary Specific Prerequisites

Deploying the binary requires [go](https://go.dev) to compile the code and running the binary in a Linux environment.
1. **Install go**: Go is required: to compile the build into an executable binary

    Clear instructions: can be found here: [https://go.dev/doc/install](https://go.dev/doc/install)

2. **Install git**
    Install git for your specific operating system [https://github.com/git-guides/install-git](https://github.com/git-guides/install-git).

3. **Clone git repo**
 
    ```console
    $ git clone https://github.com/PureStorage-OpenConnect/pure-fa-openmetrics-exporter.git
    ```

4. **Build the package**

    ```console
    $ cd pure-fa-openmetrics-exporter
    $ make build .
    ```

## Binary - Default - http passing API token with query

1. **Binary - Default - http passing API token with query**
    ```console
    $ ls out/bin
    $ .out/bin/pure-fa-openmetrics-exporter
    Start Pure FlashArray exporter v1.0.12 on 0.0.0.0:9490
    ```

2. **Test the exporter**
    Use `curl` to test the exporter returns results.
    ```console
    $ curl -H 'Authorization: Bearer 11111111-1111-1111-1111-111111111111' -X GET 'http://localhost:9490/metrics/array?endpoint=array01' -silent | grep ^purefa_info
    purefa_info{array_name="ARRAY01",os="Purity//FA",system_id="11111111-1111-1111-1111-111111111111",version="6.5.0"} 1
    ```

    We expect to return a single line displaying `purefa_info` returning the name, Purity OS, system ID and the Purity version.
    You can remove the `' -silent | grep ^purefa_info` from the command to see a list of all results.

3. **Build it as a service**

    Copy binary to `/usr/bin`
    ```console
    $ cp .out/bin/pure-fa-openmetrics-exporter /usr/bin
    ```

    Create a `.servicefile` in `/etc/systemd/system/`
    ```console
    $ cat /etc/systemd/system/purefa-ome.service
    ```

    Example:
    ```console
    [Unit]
    Description=Pure Storage FlashArray OpenMetrics Exporter
    After=network.target
    StartLimitIntervalSec=0
    [Service]
    Type=simple
    Restart=always
    RestartSec=1

    # Start OME
    ExecStart=/usr/bin/pure-fa-om-exporter --port 9490

    # Start OME with TLS
    # ExecStart=/usr/bin/pure-fa-om-exporter --port 9490 -c /etc/pki/tls/certs/purefa-ome/pure-ome.crt -k /etc/pki/tls/private/pure-ome.key

    [Install]
    WantedBy=multi-user.target 
    ```

4. **Start, enable and test the service**
    Reload the system daemon to read in changes to services

    ```console
    $ systemctl daemon-reload
    ```

    Start the service and check the status is successful
    
    ```console
    $ systemctl start purefa-ome
    $ systemctl status purefa-ome
    ```
    
    Enable the service to start on OS startup
    ```console
    $ systemctl enable purefa-ome
    ```

## Binary - Tokens File - http with API tokens file embedded in exporter
Similar steps as basic but we just need to cover a couple of minor changes to running and testing the deployment
Follow steps 1-4 and 7-8 of the default binary deployment, but substitute the following steps for executing and testing.

1. **Create a tokens.yaml file to pass to the exporter**

    ```console
    $ cat /directorypath/tokens.yaml
    # alias: of the array to be used in the exporter query
    array01:
      # FQDN or IP address of the array
      address: 'array01.fqdn.com'
      # API token of a user on the array
      api_token: '11111111-1111-1111-1111-111111111111'

    # Example of a second array
    array02:
      address: 'array02.fqdn.com'
      api_token: '22222222-2222-2222-2222-222222222222'
    ```

2. **Run the binary**
    ```console
    $ ls out/bin
    $ .out/bin/pure-fa-openmetrics-exporter --tokens /directorypath/tokens.yaml
    Start Pure FlashArray exporter v1.0.12 on 0.0.0.0:9490
    ```

3. **Test the exporter**
    Use `curl` to test the exporter returns results.
    ```console
    $ curl -X GET 'http://localhost:9490/metrics/array?endpoint=gse-array01' -silent | grep ^purefa_info
    purefa_info{array_name="ARRAY01",os="Purity//FA",system_id="11111111-1111-1111-1111-111111111111",version="6.5.0"} 1
    ```

## Binary - TLS - https passing API token with query
Similar steps as basic but we just need to cover a couple of minor changes to running and testing the deployment
Follow steps 1-4 and 7-8 of the default binary deployment, but substitute the following steps for executing and testing.

Create the certificate and key and pass the exporter the files. There are many different methods of generating certificates which we won't discuss here as each organizations has different standards and requirements.

1. **Pass the certificate and private key to the exporter**

    ```console
    $ pure-fa-om-exporter --port 9490 -c /etc/pki/tls/certs/purefa-ome/pure-ome.crt -k /etc/pki/tls/private/pure-ome.key
    ```

2. **Test the exporter**

    Use `curl` to test the exporter returns results.

    Please note to use `https` in the queries.

    **TLS https - skipping SSL verification**

    cURL with -k skips SSL verification.
    ```console
    $ curl -k -H 'Authorization: Bearer 11111111-1111-1111-1111-111111111111' -X GET 'https://localhost:9490/metrics/array?endpoint=array01'  --silent | grep ^purefa_info
    ```

    **TLS https - with SSL verification**

    Run the following commands from the server with the certificates installed.

    A basic check of the certificate is installed and the exporter is responding correctly.
    ```console
    $ curl https://pure-ome.fqdn.com:9490
    <html>
    <body>
    <h1>Pure Storage FlashArray OpenMetrics Exporter</h1>
    <table>
        <thead>
            <tr>
            <td>Type</td>
            <td>Endpoint</td>
            <td>GET parameters</td>
            <td>Description</td>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Full metrics</td>
                <td><a href="/metrics?endpoint=host">/metrics</a></td>
                <td>endpoint</td>
                <td>All array metrics. Expect slow response time.</td>
            </tr>
            <tr>
                <td>Array metrics</td>
                <td><a href="/metrics/array?endpoint=host">/metrics/array</a></td>
                <td>endpoint</td>
                <td>Provides only array related metrics.</td>
            </tr>
            <tr>
                <td>Volumes metrics</td>
                <td><a href="/metrics/volumes?endpoint=host">/metrics/volumes</a></td>
                <td>endpoint</td>
                <td>Provides only volumes related metrics.</td>
            </tr>
            <tr>
                <td>Hosts metrics</td>
                <td><a href="/metrics/hosts?endpoint=host">/metrics/hosts</a></td>
                <td>endpoint</td>
                <td>Provides only hosts related metrics.</td>
            </tr>
            <tr>
                <td>Pods metrics</td>
                <td><a href="/metrics/pods?endpoint=host">/metrics/pods</a></td>
                <td>endpoint</td>
                <td>Provides only pods related metrics.</td>
            </tr>
            <tr>
                <td>Directories metrics</td>
                <td><a href="/metrics/directories?endpoint=host">/metrics/directories</a></td>
                <td>endpoint</td>
                <td>Provides only directories related metrics.</td>
            </tr>
        </tbody>
    </table>
    </body>
    ```

    Full check using certificate.
    ```console
    $ curl --cacert pure-ome.crt -H 'Authorization: Bearer 11111111-1111-1111-1111-111111111111' -X GET 'http://pure-ome.fqdn.com:9490/metrics/array?endpoint=array01'
    ```
