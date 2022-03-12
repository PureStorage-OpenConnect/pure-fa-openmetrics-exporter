import pytest
from httpx import Headers

def test_get_with_param(app_client, endpoint, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/metrics?endpoint=' + endpoint + '&param=true', headers = h)
    assert res.status_code == 400
    _, res = app_client.get('/metrics/volumes?endpoint=' + endpoint + '&param=true', headers = h)
    assert res.status_code == 400
