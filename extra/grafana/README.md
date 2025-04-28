# Using Pure Storage FlashArray OpenMetrics Exporter with Prometheus and Grafana
How to setup Prometheus to scrape metrics and display dashboards in Grafana using [Pure Storage FlashArray OpenMetrics Exporter][1].

## Support Statement
This exporter is provided under Best Efforts support by the Pure Portfolio Solutions Group, Open Source Integrations team. For feature requests and bugs please use GitHub Issues. We will address these as soon as we can, but there are no specific SLAs.

## TL;DR
1. Configure Pure Storage OpenMetrics Exporter ([pure-fa-openmetrics-exporter][1]).
2. Deploy and configure Prometheus ([prometheus-docs][2]). [Example prometheus.yml here](../prometheus/prometheus.yml).
3. Deploy and configure Grafana ([grafana-docs][3]).
4. Import [grafana-purefa-flasharray-overview.json](grafana-purefa-flasharray-overview.json) into Grafana.

# Overview
Take a holistic overview of your Pure Storage FlashArray estate on-premise with Prometheus and Grafana to summarize statistics such as:
  * Array Utilization
  * Purity OS version
  * Data Reduction Rate
  * Number and type of open alerts

Drill down into specific arrays and identify top busy hosts while correlating read and write operations and throughput to quickly highlight or eliminate investigation enquiries.
<br>
<img src="./images/grafana_purefa_overview_dash_1.png" width="66%" height="66%">
<img src="./images/grafana_purefa_overview_dash_2.png" width="33%" height="33%">
<br>

These dashboards provide an overview of your fleet to give you early indications of potential issues and a look back in time to recent history. Once you are pulling the metrics, you can create your own dashboards bespoke to your environment, even correlating metrics from other technologies.

Would you like to collect metrics from your Pure Storage FlashArray fleet and monitor your environment on-premise? Let's walk through it.

Included in this guide
* Overview
* How does it work?
* Setup
* Troubleshooting No Data Issues
* Troubleshooting Specific Errors
* Pure Storage FlashArray Overview Grafana Dashboard

# How does it work?
Time-series database Prometheus polls the Pure Storage OpenMetrics exporter (OME) with details of the device endpoint (Pure Storage FlashArray).

Pure Storage OME polls the device and relays the metrics back to Prometheus for data retention.
Now we can query the Prometheus database either in Prometheus UI or using Grafana.

Grafana can be configured to query all of the metrics available in the Prometheus database and display them in meaningful dashboards. These dashboards can be imported or developed further, not just to correlate data of the FlashArray, but exporters for other elements of your infrastructure can be queried and correlated bespoke to your environment.
<br>
<img src="./images/pure_om_exporters_prometheus_grafana.png">
<br>

# Setup
## Prerequisites and Dependencies
This deployment assumes the [Pure Storage FlashArray OpenMetrics Exporter][1] is previously setup and configured.
Prometheus and Grafana are open source platforms which are under constant development, therefore updates to our overview dashboard cannot be tested with every version. When new versions of the dashboard is released we endeavor to test between current and previous tested versions.
We have tested our dashboard with the following versions:

The Grafana dashboards have been developed and tested using the following software versions:
| OME            | Prometheus  | Grafana        |
| -------------- | ----------- |--------------- |
| v1.0.14        | 2.45.1 LTS  | 9.3.2 & 10.1.15 & 10.3.3 |
| v1.0.9         | 2.45.1 LTS  | 9.3.2 & 10.1.15|
| v1.0.9         | 2.45.0      | 9.3.2 & 10.0.1 |
| v1.0.5.hotfix1 | 2.41.1 LTS  | 9.3.2 & 9.4.1  |

Purity 6.1.0 and above (REST API 2.x continues to be developed in later Purity versions)
Dashboards may have limited functionality with earlier versions and some modifications may be required.

## Prometheus
1. Install Prometheus on your chosen OS platform ([prometheus-docs][2]).

2. Generate an API token from your chosen user account or create a new readonly user.

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

    ![Alt text](../images/purefa_create_api_token.png)
    </details>

3. Configure `/etc/prometheus/prometheus.yml` to point use the OpenMetrics exporter to query the device endpoint.

    This is an example [prometheus.yml](../prometheus/prometheus.yml) file.

    Let's take a walkthrough an example of scraping the `/metrics/array` endpoint.

    ```yaml
    # Scrape job for one Pure Storage FlashArray scraping /metrics/array
    # Each Prometheus scrape requires a job name. In this example we have structures the name `exporter_endpoint_arrayname`
      - job_name: 'purefa_array_arrayname01'
        # Specify the array endpoint from /metrics/array
        metrics_path: /metrics/array
        # Provide FlashArray authorization API token
        authorization:
          credentials: 11111111-1111-1111-1111-111111111111
        # Provide parameters to pass the exporter the device to connect to. Provide FQDN or IP address
        params:
          endpoint: ['arrayname01.fqdn.com']

        static_configs:
        # Tell Prometheus which exporter to make the request
        - targets:
          - 10.0.2.10:9490
          # Finally provide labels to the device.
          labels:
            # Instance should be the device name and is used to correlate metrics between different endpoints in Prometheus and Grafana. Ensure this is the same for each endpoint for the same device.
            instance: arrayname01
            # location, site and env are specific to your environment. Feel free to add more labels but maintain these three to minimize changes to Grafana which is expecting to use location, site and env as filter variables. 
            location: uk
            site: London
            env: production

      # Repeat for the above for end points:
      # /metrics/volumes
      # /metrics/hosts
      # /metrics/pods
      # /metrics/directories
      # It is recommended to collect expensive queries less frequently such as /metrics/directories.
      scrape_interval:     15m # Set the scrape interval to every 15min. Default is every 1 minute. This overrides the global setting.
      scrape_timeout:      15m # Set the scrape timeout to shorter than or equal to scrape_interval. Default is every 1 minute.

      # Repeat again for more Pure Storage FlashArrays
      ```

4. Test the prometheus.yml file is valid

    ```console
    > promtool check config /etc/prometheus/prometheus.yml
    Checking prometheus.yml
    SUCCESS: prometheus.yml is valid prometheus config file syntax
    ```

5. Restart Prometheus to ensure changes take effect

6. Navigate to your Prometheus instance via web browser http://prometheus:9090
  -  Type `purefa_info` in the query box and hit return
  
7. All going well, you will see your device listed:
    ```
    purefa_info{array_name="ARRAYNAME01", env="production", instance="arrayname01", job="purefa_array_arrayname01", location="uk", os="Purity//FA", site="London", system_id="11111111-1111-1111-1111-111111111111", version="6.5.0"}
    ```

## Grafana

1. Install Grafana on your chosen OS platform ([grafana-docs][3]).
2. Point Grafana as your Prometheus data source ([grafana-datasource][4]]).
<br>
<img src="./images/grafana_add_datasource.png" width="40%" height="40%">
<br>

3. Download dashboard(s) from this folder postfixed with .json.
    Either copy the contents of the .json file or download the file and import it into Grafana.
    [Pure Storage FlashArray Overview Grafana Dashboard](grafana-purefa-flasharray-overview.json)
4. Import the .json file into Grafana and specify the Prometheus datasource.
5. Now open the dashboard and check that data from your devices are visible.
<br>
<img src="./images/grafana_purefa_overview_dash_1.png" width="40%" height="40%">
<br>


# Troubleshooting No Data Issues
Whilst this guide does not replace Prometheus and Grafana support and community portals, it should give you an insight on where we might look first and where to perform checks.

## No data is visible in dashboard
Check the data is accessible to each component in the stack. If at any on these points do not work, resolve them before moving on to the next component.
  * Check Pure OpenMetrics Exporter
  * Check Prometheus
  * Check Grafana

### Check Pure OpenMetrics Exporter
1. Run cURL against the exporter and pass is the bearer token and endpoint. 
    ```
    curl -H 'Authorization: Bearer 11111111-1111-1111-1111-111111111111' -X GET 'http://<exporter_ip>:9490/metrics/array?endpoint=arrayname01.fqdn.com'
    ```

### Check Prometheus
2. Using the Prometheus UI, run a simple query to see if any results are returned.
<br>
<img src="./images/prometheus_purefa_simple_query.png" width="40%" height="40%">
<br>

3. Make sure that in the query results from step 1 include the required labels `instance`, `location`, `site`, `env` for all targets.

   If the labels are not being added, in [prometheus.yml](../prometheus/prometheus.yml), check that the required labels are being applied in accordance with the [sample configuration](../prometheus/prometheus.yml).
    ```yaml
          labels:
            # Instance should be the device name and is used to correlate metrics between different endpoints in Prometheus and Grafana. Ensure this is the same for each endpoint for the same device.
            instance: arrayname01
            # location, site and env are specific to your environment. Feel free to add more labels but maintain these three to minimize changes to Grafana which is expecting to use location, site and env as filter variables. 
            location: uk
            site: London
            env: production
    ```

4. If the query does not return results, check the status of the targets for status errors.
<br>
<img src="./images/prometheus_purefa_target_status.png" width="40%" height="40%">
<br>

5. Run prometheus.yml through the yaml checker. Check the configuration is correct and restart Prometheus.
    ```console
    > promtool check config /etc/prometheus/prometheus.yml
    Checking prometheus.yml
    SUCCESS: prometheus.yml is valid prometheus config file syntax
    ```

6. Check messages log for Prometheus errors.

### Check Grafana
7. Perform a simple test for Grafana by navigating to 'Explore' and entering a simple query.
<br>
<img src="./images/grafana_purefa_simple_query.png" width="40%" height="40%">
<br>

### Check Docker container is running
8. Ensure that the docker image is running, if for any reason the container is unexpectedly halting while running you may add `--restart unless-stopped` to the `docker run` command to ensure the container is restarted automatically on failure.
   

# Troubleshooting Specific Errors
## Some panels have errors
If the panels have errors then the query could be unsupported. Check the versions of Prometheus and Grafana are above the prerequisite versions tested above.

## Selected graphs are not loading or taking a long time to load
Volumes panels may take longer to load when a greater number of more FlashArray's are selected combined with longer time ranges are selected.

The more samples Prometheus is running queries against, the longer the query will take to complete and may even time out. If we are scraping data for 10x FlashArrays every 30 seconds, each with a modest 1,000 volumes and we set a time range to 15 days, the sum is:

`2x (samples per min) X 60 min X 24 hours X 15 days X 10 FlashArrays X 1,000 volumes = 432,000,000 sampled metrics.`

Now Prometheus is trying to calculate the top(k) of these results.

Try reducing the scope of the query and/or increase processing and memory resources to Prometheus.

# Pure Storage FlashArray Overview Grafana Dashboard

## Dashboard Template Variables and Filters
The dashboards are fully templated which means they will work with your FlashArray configuration and allow you to filter by environment, array and how many top(k) metrics you wish to display on your dashboard.
<br>
<img src="./images/grafana_purefa_dashboard_template.png">
<br>

## Average values
Many of the dashboard panels average metrics to help smooth out graphs and identify trends rather than isolated max and min values. Adjust the time range to drill down in to specific values.
## Accessibility - Dark/Light Mode
Colors have been selected to work equally in both dark and light mode. 
<br>
<img src="./images/grafana_purefa_dark_light_modes.png">
<br>

## Threshold Defaults
### Utilization
Grafana dashboards are configured with utilization thresholds to pro-actively highlight potential capacity issues with plenty of time to address. These values can be adjusted to suit your own threshold policies.

The utilization bar and value text will change color according to the array utilization.

* <70% Base is Pure Ocean
* $\ge$ 70% is Pure Yellow
* $\ge$ 80% is Pure Orange
* $\ge$ 90% is Pure Magenta
<br>
<img src="./images/grafana_utilization_thresholds.png" width="40%" height="40%">
<br>

## Alerts
Alerts are set to highlight how many and how serious the alerts across your FlashArray fleet.
Alert panel is ordered by Critical alerts to ensure any arrays with critical alerts are moved to the top of the table for your immediate attention. 

**Total number of alerts**
* 0  background is transparent zero & replaced with value '-'
* $\gt$ 0 is above threshold

**Alert Criticality**
* $\gt$ 0 Info Alert background is Pure Ocean
* $\gt$ 0 Warning Alert background is Pure Magenta
* $\gt$ 0 Critical Alert background is Pure Magenta

## Data Reduction Rate
Data Reduction Rate also has a threshold set for informational purposes to highlight when an array might be storing incompressible or unique datasets. This can be adjusted to suit your environment policies.
* Default Data Reduction Rate <1.01:1 = Pure Yellow
<br>
<img src="./images/grafana_purefa_dashboard_drr_threshold.png" width="50%" height="50%">
<br>

[1]: https://github.com/PureStorage-OpenConnect/pure-fa-openmetrics-exporter "pure-fa-openmetrics-exporter"
[2]: https://prometheus.io/docs/introduction/overview/ "prometheus-docs"
[3]: https://grafana.com/docs/grafana/latest/ "grafana-docs"
[4]: https://grafana.com/docs/grafana/latest/administration/data-source-management/ "grafana-datasource"

## Known Issues

### Average Top Latent Volumes Array & IO Type Takes a Long Time to Load or Timesout
This query is very expensive on Prometheus and the browser. For large environments it may be advisable to use a tabular view or create a recording rule. https://prometheus.io/docs/prometheus/latest/querying/basics/#avoiding-slow-queries-and-overloads