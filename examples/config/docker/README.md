### Start the exporter with default settings

The application executes with the default numner of worker processes (2) and with the SSL warning enabled

```shell

docker run -d -p 9491:9491  --rm --name pure-exporter pure-fa-ome:<version>
```

### Start the exporter with specific Gunicorn parameters

Specific Gunicorn parmeters can be passed to the application in the docker command line

```shell
docker run -d -p 9491:9491  --rm --name pure-exporter pure-fa-ome:<version> --workers=<n_workers> --access-logfile=- --error-logfile=-
```

### Disable SSL warning

If you want to prevent the app from logging a warning when the FB endpoint is not provided with a trusted SSL certificate, you can disable that by invoking the Flask app with the <kbd>disable_ssl_warn=True</kbd> parameter

```shell
docker run -d -p 9491:9491  --rm --name pure-exporter pure-fa-ome:<version> 'pure_fa_exporter:create_app(disable_ssl_warn=True)'
```
