import pytest
from pypureclient import PureError
from pure_fa_openmetrics_exporter.flasharray_client import client


@pytest.fixture()
def fa_client(scope="session"):
    try:
        c = client.FlasharrayClient(
                             target = '10.225.112.90',
                             api_token = 'b5cb29e7-a93c-b40a-b02f-da2b90c8c65e', 
                             disable_ssl_warn = True)
    except PureError as pe:
        pytest.fail("Could not connect to flasharray {0}".format(pe))
    yield c
