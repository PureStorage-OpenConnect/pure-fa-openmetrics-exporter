
def test_array(fa_client):
    arr1 = fa_client.arrays()
    arr2 = fa_client.arrays()
    assert arr1 == arr2

def test_arrays_performance(fa_client):
    array_perf1 = fa_client.arrays_performance()
    array_perf2 = fa_client.arrays_performance()
    assert array_perf1 == array_perf2

def test_arrays_space(fa_client):
    array_space1 = fa_client.arrays_space()
    array_space2 = fa_client.arrays_space()
    assert array_space1 == array_space2

def test_volumes_performance(fa_client):
    volumes_perf1 = fa_client.volumes_performance()
    volumes_perf2 = fa_client.volumes_performance()
    assert volumes_perf1 == volumes_perf2

def test_volumes_space(fa_client):
    volumes_space1 = fa_client.volumes_space()
    volumes_space2 = fa_client.volumes_space()
    assert volumes_space1 == volumes_space2

def test_pods_performance(fa_client):
    pods_perf1 = fa_client.pods_performance()
    pods_perf2 = fa_client.pods_performance()
    assert pods_perf1 == pods_perf2

def test_pods_space(fa_client):
    pods_space1 = fa_client.pods_space()
    pods_space2 = fa_client.pods_space()
    assert pods_space1 == pods_space2

def test_hosts_performance(fa_client):
    hosts_perf1 = fa_client.hosts_performance()
    hosts_perf2 = fa_client.hosts_performance()
    assert hosts_perf1 == hosts_perf2

def test_hosts_space(fa_client):
    hosts_space1 = fa_client.hosts_space()
    hosts_space2 = fa_client.hosts_space()
    assert hosts_space1 == hosts_space2

def test_directories_performance(fa_client):
    directories_perf1 = fa_client.directories_performance()
    directories_perf2 = fa_client.directories_performance()
    assert directories_perf1 == directories_perf2

def test_directories_space(fa_client):
    directories_space1 = fa_client.directories_space()
    directories_space2 = fa_client.directories_space()
    print(directories_space1)
    assert directories_space1 == directories_space2
