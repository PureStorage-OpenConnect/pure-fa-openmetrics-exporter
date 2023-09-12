### Start the exporter with default settings

Launch the exporter using the docker command according to the following: 

```shell

docker run -d -p 9490:9490  --rm --name pure-fa-om-exporter quay.io/purestorage/pure-fa-om-exporter:<version>
```

## Docker Compose Implementation

This Docker Compose manifest starts up three Services, prometheus, grafana and the pure-fa-openmetrics exporter. Each Service has a basic configuration to help you get off the ground as fast as possible with this monitoring stack. 

1. Edit `./prometheus/prometheus.yml` for the scrape configuration for your enviroment. A basic implementation is provided for a single Array, its Hosts and Volumes. You will need to update `authorization.credential` and `params.endpoint` in each scrape config.
2. Change your working directory to the same as the `docker-compose.yaml` file
3. Then use `docker compose up` to launch the monitoring stack. Add `--detach` to run in background.
4. Open a browser to [http://localhost:3000](http://localhost:3000). The username is `admin` and the password is `admin!`.