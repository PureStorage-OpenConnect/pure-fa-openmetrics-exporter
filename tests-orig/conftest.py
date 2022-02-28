import pytest


@pytest.fixture()
def mock_fa_client(monkeypatch):
    monkeypatch.syspath_prepend('tests/lib-fake')
    from purity_fa_prometheus_exporter.flasharray_rest_wrapper import flasharray_local
    fa_client = flasharray_local.FlashArrayLocal()
    yield fa_client

@pytest.fixture
def client(monkeypatch):
    monkeypatch.syspath_prepend('tests/lib-fake')
    from purity_fa_prometheus_exporter import fa_exporter
    app = fa_exporter.create_app()
    app.testing = True
    client = app.test_client()
    yield client
