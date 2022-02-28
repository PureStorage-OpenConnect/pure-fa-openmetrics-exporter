import pytest

def test_home(client):
    rv = client.get("/")
    assert b'/metrics' in rv.data
