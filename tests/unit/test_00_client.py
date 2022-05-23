import time

'''
These tests are meant to verify the fa_client actually caches the results of each query. 
The 2 seconds sleep between two calls to the same method should be enough to guarantee
that if those were two different calls to the FA endpoint the sampled metrics are different.
'''

def test_alerts(fa_client):
    alerts1 = fa_client.alerts()
    time.sleep(2)
    alerts2 = fa_client.alerts()
    assert alerts1 == alerts2

def test_array(fa_client):
    array1 = fa_client.arrays()
    time.sleep(2)
    array2 = fa_client.arrays()
    assert array1 == array2

def test_volumes(fa_client):
    volumes1 = fa_client.volumes()
    time.sleep(2)
    volumes2 = fa_client.volumes()
    assert volumes1 == volumes2

def test_pods(fa_client):
    pods1 = fa_client.pods()
    time.sleep(2)
    pods2 = fa_client.pods()
    assert pods1 == pods2

def test_hosts(fa_client):
    hosts1 = fa_client.hosts()
    time.sleep(2)
    hosts2 = fa_client.hosts()
    assert hosts1 == hosts2

def test_directories(fa_client):
    directories1 = fa_client.directories()
    time.sleep(2)
    directories2 = fa_client.directories()
    assert directories1 == directories2

def test_connections(fa_client):
    connections1 = fa_client.connections()
    time.sleep(2)
    connections2 = fa_client.connections()
    assert connections1 == connections2

def test_nics(fa_client):
    nics1 = fa_client.network_interfaces_performance()
    time.sleep(2)
    nics2 = fa_client.network_interfaces_performance()
    assert nics1 == nics2

def test_hw(fa_client):
    hw1 = fa_client.hardware()
    time.sleep(2)
    hw2 = fa_client.hardware()
    assert hw1 == hw2
