import pytest

def test_array(mock_fa_client):
    array = mock_fa_client.get_array()
    assert array['performance']['id'] == array['id']
    assert array['space']['capacity'] is not None
