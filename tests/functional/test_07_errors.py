import pytest
from httpx import Headers

def test_not_existing_endpoint(app_client, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/not-existing', headers = h)
    assert res.status_code == 404

@pytest.mark.parametrize("test_input,expected",
                         [('/metrics/array/', 400),
                          ('/metrics/array/not-an-array', 404),
                          ('/metrics/volumes/', 400),
                          ('/metrics/volumes/not-a-volume', 404),
                          ('/metrics/hosts/', 400),
                          ('/metrics/hosts/not-a-host', 404),
                          ('/metrics/directories/', 400),
                          ('/metrics/directories/not-a-dir', 404),
                          ('/metrics/pods/', 400),
                          ('/metrics/pods/not-a-pod', 404)])
def test_wrong_endpoint(app_client, api_token, test_input, expected):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get(test_input, headers = h)
    assert res.status_code == expected

@pytest.mark.parametrize("test_input,expected",
                         [('/', 405),
                          ('/metrics/array', 405),
                          ('/metrics/array/', 405),
                          ('/metrics/array/not-an-array', 404),
                          ('/metrics/volumes', 405),
                          ('/metrics/volumes/', 405),
                          ('/metrics/volumes/not-a-volume', 404),
                          ('/metrics/hosts', 405),
                          ('/metrics/hosts/', 405),
                          ('/metrics/hosts/not-a-host', 404),
                          ('/metrics/directories', 405),
                          ('/metrics/directories/', 405),
                          ('/metrics/directories/not-a-dir', 404),
                          ('/metrics/pods', 405),
                          ('/metrics/pods/', 405),
                          ('/metrics/pods/not-a-pod', 404)])
def test_post_put_del(app_client, api_token, test_input, expected):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.post(test_input, headers = h)
    assert res.status_code == expected
    _, res = app_client.put(test_input)
    assert res.status_code == expected
    _, res = app_client.delete(test_input)
    assert res.status_code == expected
