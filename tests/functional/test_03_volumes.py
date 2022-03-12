import pytest
from httpx import Headers

def test_volumes(app_client, endpoint, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/metrics/volumes?endpoint=' + endpoint, headers = h)
    for e in ['purefa_volume_space_data_reduction_ratio',
              'purefa_volume_space_size_bytes',
              'purefa_volume_space_used_bytes',
              'snapshots',
              'snapshots_effective',
              'total_effective',
              'total_physical',
              'total_provisioned',
              'unique',
              'virtual',
              'purefa_volume_performance_latency_usec',
              'purefa_volume_performance_bandwidth_bytes',
              'purefa_volume_performance_iops',
              'purefa_volume_performance_avg_block_bytes',
              'queue_usec_per_mirrored_write_op',
              'queue_usec_per_read_op',
              'queue_usec_per_write_op',
              'san_usec_per_mirrored_write_op',
              'san_usec_per_read_op',
              'san_usec_per_write_op',
              'service_usec_per_mirrored_write_op',
              'service_usec_per_read_op',
              'service_usec_per_write_op',
              'usec_per_mirrored_write_op',
              'usec_per_read_op',
              'usec_per_write_op',
              'read_bytes_per_sec',
              'write_bytes_per_sec',
              'mirrored_write_bytes_per_sec',
              'reads_per_sec',
              'writes_per_sec',
              'mirrored_writes_per_sec',
              'bytes_per_read',
              'bytes_per_write',
              'bytes_per_op',
              'bytes_per_mirrored_write']:
        assert e in res.text
