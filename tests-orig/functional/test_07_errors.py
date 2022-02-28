import pytest


def test_not_existing_endpoint(client):
    rv = client.get("/not-existing")
    assert rv.status_code == 404

@pytest.mark.parametrize("test_input,expected",
                         [('/metrics/array/', 404),
                          ('/metrics/array/not-an-array', 404),
                          ('/metrics/volumes/', 404),
                          ('/metrics/volumes/not-a-volume', 404),
                          ('/metrics/hosts/', 404),
                          ('/metrics/hosts/not-a-host', 404),
                          ('/metrics/directories/', 404),
                          ('/metrics/directories/not-a-dir', 404),
                          ('/metrics/pods/', 404),
                          ('/metrics/pods/not-a-pod', 404)])
def test_wrong_endpoint(client, test_input, expected):
    rv = client.get(test_input)
    assert rv.status_code == expected

@pytest.mark.parametrize("test_input,expected",
                         [('/', 405),
                          ('/metrics/array', 405),
                          ('/metrics/array/', 404),
                          ('/metrics/array/not-an-array', 404),
                          ('/metrics/volumes', 405),
                          ('/metrics/volumes/', 404),
                          ('/metrics/volumes/not-a-volume', 404),
                          ('/metrics/hosts', 405),
                          ('/metrics/hosts/', 404),
                          ('/metrics/hosts/not-a-host', 404),
                          ('/metrics/directories', 405),
                          ('/metrics/directories/', 404),
                          ('/metrics/directories/not-a-dir', 404),
                          ('/metrics/pods', 405),
                          ('/metrics/pods/', 404),
                          ('/metrics/pods/not-a-pod', 404)])
def test_post_put_del(client, test_input, expected):
    rv = client.post(test_input)
    assert rv.status_code == expected
    rv = client.put(test_input)
    assert rv.status_code == expected
    rv = client.delete(test_input)
    assert rv.status_code == expected
