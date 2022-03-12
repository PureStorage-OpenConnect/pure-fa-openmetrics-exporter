import pytest
from pypureclient import PureError
from pure_fa_openmetrics_exporter.flasharray_client import client
from pure_fa_openmetrics_exporter import pure_fa_exporter
from sanic_testing.testing import SanicTestClient

@pytest.fixture(scope="session")
def fa_client():
    try:
        c = client.FlasharrayClient(
                             target = '10.225.112.90',
                             api_token = 'b5cb29e7-a93c-b40a-b02f-da2b90c8c65e', 
                             disable_ssl_warn = True)
    except PureError as pe:
        pytest.fail("Could not connect to flasharray {0}".format(pe))
    yield c

@pytest.fixture()
def api_token():
    return 'b5cb29e7-a93c-b40a-b02f-da2b90c8c65e' 

@pytest.fixture()
def endpoint():
    return '10.225.112.90' 

@pytest.fixture(scope="session")
def app_client():
    app = pure_fa_exporter.app
    app.ctx.disable_cert_warn = True
    client = SanicTestClient(app)
    return client
