import pytest
from flask import request

def test_array(client):
    rv = client.get('/metrics/array')
    assert b'purefa_info{array_name=' in rv.data
