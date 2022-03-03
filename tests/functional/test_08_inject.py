import pytest


def test_get_with_param(client):
    rv = client.get('/metrics?param=true')
    rv = client.get('/metrics/volumes?param=true')
    assert rv.status_code == 400
