import pytest


def test_pods(client):
    rv = client.get("/metrics/pods")
    for e in ['purefa_pod_space_datareduction_ratio',
              'purefa_pod_space_size_bytes',
              'purefa_pod_space_bytes',
              'replication',
              'shared',
              'snapshots',
              'system',
              'thin_provisioning',
              'total_physical',
              'total_provisioned',
              'total_reduction',
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
              'service_usec_per_read_op_cache_reduction',
              'service_usec_per_write_op',
              'usec_per_mirrored_write_op',
              'usec_per_read_op',
              'usec_per_write_op',
              'purefa_pod_performance_bandwidth_bytes',
              'read_bytes_per_sec',
              'write_bytes_per_sec',
              'mirrored_write_bytes_per_sec']:
        assert b"e" in rv.data

