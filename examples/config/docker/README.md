### Start the exporter with default settings

The application executes with the default number of worker processes (1) and with the SSL warning enabled

```shell

docker run -d -p 9491:9491  --rm --name pure-exporter pure-fa-ome:<version>
```

### Start the exporter with specific Sanic parameters

Specific Sanic parmeters can be passed to the application in the docker command line

```shell
docker run -d -p 9491:9491  --rm --name pure-exporter pure-fa-ome:<version> --workers=<n_workers> --host 0.0.0.0 --port 9491
```

### Disable SSL warning

If you want to prevent the app from logging a warning when the FA endpoint is not provided with a trusted SSL certificate, you can disable that by invoking the app with the <kbd>--disable-cert-warning</kbd> flag
```shell
docker run -d -p 9491:9491  --rm --name pure-exporter pure-fa-ome:<version> --disable-cert-warning
