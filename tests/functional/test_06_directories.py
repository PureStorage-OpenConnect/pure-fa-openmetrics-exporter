import pytest
from httpx import Headers

def test_directories(app_client, endpoint, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/metrics/directories?endpoint=' + endpoint, headers = h)
    for e in ['purefa_directory_space_data_reduction_ratio',
              'purefa_directory_space_size_bytes',
              'purefa_directory_space_used_bytes',
              'name',
              'filesystem',
              'path',
              'space',
              'snapshots',
              'total_physical',
              'unique',
              'purefa_directory_performance_latency_usec',
              'purefa_directory_performance_bandwidth_bytes',
              'purefa_directory_performance_iops',
              'purefa_directory_performance_avg_block_bytes',
              'usec_per_read_op',
              'usec_per_write_op',
              'usec_per_other_op',
              'write_bytes_per_sec',
              'read_bytes_per_sec',
              'reads_per_sec',
              'writes_per_sec',
              'others_per_sec',
              'bytes_per_op',
              'bytes_per_read',
              'bytes_per_write']:
        assert e in res.text

