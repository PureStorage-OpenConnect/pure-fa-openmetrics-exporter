import pytest
from httpx import Headers

def test_pods(app_client, endpoint, api_token):
    h = Headers({'Authorization': "Bearer " + api_token})
    _, res = app_client.get('/metrics/pods?endpoint=' + endpoint, headers = h)

    for e in ['purefa_pod_space_data_reduction_ratio',
              'purefa_pod_space_size_bytes',
              'purefa_pod_space_used_bytes',
              'replication',
              'shared',
              'snapshots',
              'total_physical',
              'total_provisioned',
              'unique',
              'purefa_pod_performance_latency_usec',
              'purefa_pod_performance_bandwidth_bytes',
              'purefa_pod_performance_iops',
              'purefa_pod_performance_avg_block_bytes',
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
              'purefa_pod_performance_bandwidth_bytes',
              'read_bytes_per_sec',
              'write_bytes_per_sec',
              'mirrored_write_bytes_per_sec']:
        assert e in res.text

