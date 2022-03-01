import time

'''
These tests are meant to verify the fa_client actually caches the results of each query. 
The 2 seconds sleep between two calls to the same method should be enough to guarantee
that if those were two differen calls to the FA endpoint the sampled metrics are different.
'''

def test_array(fa_client):
    array1 = fa_client.arrays()
    time.sleep(2)
    array2 = fa_client.arrays()
    print(array1)
    assert array1 == array2

def test_arrays_performance(fa_client):
    array_perf1 = fa_client.arrays_performance()
    time.sleep(2)
    array_perf2 = fa_client.arrays_performance()
    print(array_perf1)
    assert array_perf1 == array_perf2

def test_arrays_space(fa_client):
    array_space1 = fa_client.arrays_space()
    time.sleep(2)
    array_space2 = fa_client.arrays_space()
    print(array_space1)
    assert array_space1 == array_space2

def test_volumes(fa_client):
    volumes1 = fa_client.volumes()
    time.sleep(2)
    volumes2 = fa_client.volumes()
    print(volumes1)
    assert volumes1 == volumes2

def test_volumes_performance(fa_client):
    volumes_perf1 = fa_client.volumes_performance()
    time.sleep(2)
    volumes_perf2 = fa_client.volumes_performance()
    assert volumes_perf1 == volumes_perf2

def test_volumes_space(fa_client):
    volumes_space1 = fa_client.volumes_space()
    time.sleep(2)
    volumes_space2 = fa_client.volumes_space()
    print(volumes_space1)
    assert volumes_space1 == volumes_space2

def test_pods(fa_client):
    pods1 = fa_client.pods()
    time.sleep(2)
    pods2 = fa_client.pods()
    print(pods1)
    assert pods1 == pods2

def test_pods_performance(fa_client):
    pods_perf1 = fa_client.pods_performance()
    time.sleep(2)
    pods_perf2 = fa_client.pods_performance()
    assert pods_perf1 == pods_perf2

def test_pods_space(fa_client):
    pods_space1 = fa_client.pods_space()
    time.sleep(2)
    pods_space2 = fa_client.pods_space()
    print(pods_space1)
    assert pods_space1 == pods_space2

def test_hosts(fa_client):
    hosts1 = fa_client.hosts()
    time.sleep(2)
    hosts2 = fa_client.hosts()
    print(hosts1)
    assert hosts1 == hosts2

def test_hosts_performance(fa_client):
    hosts_perf1 = fa_client.hosts_performance()
    time.sleep(2)
    hosts_perf2 = fa_client.hosts_performance()
    assert hosts_perf1 == hosts_perf2

def test_hosts_space(fa_client):
    hosts_space1 = fa_client.hosts_space()
    time.sleep(2)
    hosts_space2 = fa_client.hosts_space()
    print(hosts_space1)
    assert hosts_space1 == hosts_space2

def test_directories(fa_client):
    directories1 = fa_client.directories()
    time.sleep(2)
    directories2 = fa_client.directories()
    print(directories1)
    assert directories1 == directories2

def test_directories_performance(fa_client):
    directories_perf1 = fa_client.directories_performance()
    time.sleep(2)
    directories_perf2 = fa_client.directories_performance()
    assert directories_perf1 == directories_perf2

def test_directories_space(fa_client):
    directories_space1 = fa_client.directories_space()
    time.sleep(2)
    directories_space2 = fa_client.directories_space()
    print(directories_space1)
    assert directories_space1 == directories_space2
