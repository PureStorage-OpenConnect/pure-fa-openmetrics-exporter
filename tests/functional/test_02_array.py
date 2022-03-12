import pytest
from httpx import Headers

def test_array(app_client, endpoint, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/metrics/array?endpoint=' + endpoint, headers = h)
    assert 'purefa_info{array_name=' in res.text
