!/usr/bin/env python

from functools import wraps

from sanic import Sanic
from sanic.log import logger
from sanic.response import text, html, raw
from sanic.exceptions import SanicException
from sanic.handlers import ErrorHandler
from prometheus_client import generate_latest, CollectorRegistry
from pure_fa_openmetrics_exporter.flasharray_collector.collector import FlasharrayCollector
from pure_fa_openmetrics_exporter.flasharray_client.client import FlasharrayClient

import re
import argparse

class CustomHandler(ErrorHandler):
    def default(self, request, exception):
        # Here, we have access to the exception object
        # and can do anything with it (log, send to external service, etc)

        # Some exceptions are trivial and built into Sanic (404s, etc)
        if not isinstance(exception, SanicException):
            print(exception)

        # Then, we must finish handling the exception by returning
        # our response to the client
        # For this we can just call the super class' default handler
        return super().default(request, exception)

excp_handler = CustomHandler()
app = Sanic('purefa_openmetrics_exporter', error_handler=excp_handler)


def check_request_for_authorization_status(request):
    pattern_str = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
    if (request.token is None):
        return False
    regx = re.compile(pattern_str)
    match = regx.search(request.token)
    return match is not None

def authorized(f):
    @wraps(f)
    async def decorated_function(request, *args, **kwargs):
        is_authorized = check_request_for_authorization_status(request)

        if is_authorized:
            # the user is authorized.
            # run the handler method and return the response
            response = await f(request, *args, **kwargs)
            return response
        else:
            # the user is not authorized.
            return text("not_authorized", 403)
    return decorated_function

@app.get('/')
async def index_handler(request):
    """Display an overview of the exporters capabilities."""
    msg = '''
<h1>Pure Storage Flasharray OpenMetrics Exporter</h1>
<table>
    <thead>
        <tr>
            <td>Type</td>
            <td>Endpoint</td>
        </tr>
    </thead>
        <tbody>
        <tr>
            <td>Full metrics</td>
            <td><a href="/metrics">/metrics</a></td>
        </tr>
        <tr>
            <td>Volumes metrics</td>
            <td><a href="/metrics/volumes">/metrics/volumes</a></td>
            <td>Retrieves only volume related metrics</td>
        </tr>
        <tr>
            <td>Hosts metrics</td>
            <td><a href="/metrics/hosts">/metrics/hosts</a></td>
            <td>Retrieves only host related metrics</td>
        </tr>
        <tr>
            <td>Pods metrics</td>
            <td><a href="/metrics/pods">/metrics/pods</a></td>
            <td>Retrieves only pod related metrics</td>
        </tr>
        <tr>
            <td>Directories metrics</td>
            <td><a href="/metrics/directories">/metrics/directories</a></td>
            <td>Retrieves only directories related metrics</td>
        </tr>
    </tbody>
</table>
'''
    return html(msg)

@app.get('/metrics/<tag:all|array|volumes|hosts|pods|directories>/")
@authorized
def flasharray_handler():
    """Produce FlashArray metrics."""
    registry = CollectorRegistry()
    collector = FlasharrayCollector
    endpoint = request.args.get('endpoint', None)
    fa_client = FlasharrayClient(endpoint, request.token, app.ctx.disable_cert_warn)
    registry.register(collector(fb_client, request=tag))
    resp = generate_latest(registry)
    del fa_client, collector, registry
    return raw(resp)


@app.get('/metrics', strict_slashes=True)
def flasharray_handler_full(request):
    return flasharray_handler(request, 'all')


if __name__ == "__main__":
    argparser = argparse.ArgumentParser()
    argparser.add_argument("-H", "--host", default="127.0.0.1", help="Address to host the server on")
    argparser.add_argument("-P", "--port", default="9491", help="Port to host the server on")
    argparser.add_argument("-D", "--debug", default=False, help="Run in debug mode")
    argparser.add_argument("-W", "--workers", type=int, default=1, help="Number of workers")
    argparser.add_argument("-L", "--log", default=False, action="store_true", help="Enable log")
    argparser.add_argument("-X", "--disable-cert-warning", action="store_true",
                           help = "Disable SSL certificate verification warning")
    args = argparser.parse_args()
    app.ctx.disable_cert_warn = args.disable_cert_warning
    app.run(host=args.host, port=args.port, workers=args.workers,
            access_log=args.log, debug=args.debug)

