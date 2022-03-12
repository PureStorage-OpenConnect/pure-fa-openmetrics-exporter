import pytest
from httpx import Headers

def test_home(app_client, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/', headers = h)
    assert res.status_code == 200
