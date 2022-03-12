import pytest
from httpx import Headers

def test_hosts(app_client, endpoint, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/metrics/hosts?endpoint=' + endpoint, headers = h)

    for e in ['purefa_host_space_data_reduction_ratio',
              'purefa_host_space_size_bytes',
              'purefa_host_space_used_bytes',
              'snapshots',
              'total_physical',
              'total_provisioned',
              'unique',
              'purefa_host_performance_latency_usec',
              'purefa_host_performance_bandwidth_bytes',
              'purefa_host_performance_iops',
              'queue_usec_per_read_op',
              'queue_usec_per_write_op',
              'queue_usec_per_mirrored_write_op',
              'san_usec_per_read_op',
              'san_usec_per_write_op',
              'san_usec_per_mirrored_write_op',
              'service_usec_per_mirrored_write_op',
              'service_usec_per_read_op',
              'service_usec_per_write_op',
              'usec_per_read_op',
              'usec_per_write_op',
              'usec_per_mirrored_write_op',
              'read_bytes_per_sec',
              'write_bytes_per_sec',
              'mirrored_write_bytes_per_sec',
              'read_bytes_per_sec',
              'write_bytes_per_sec',
              'mirrored_write_bytes_per_sec']:
        assert e in res.text
