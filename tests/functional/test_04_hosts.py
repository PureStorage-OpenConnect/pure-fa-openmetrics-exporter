import pytest


def test_hosts(client):
    rv = client.get("/metrics/hosts")
    for e in ['purefa_host_space_datareduction_ratio',
              'purefa_host_space_size_bytes',
              'purefa_host_space_bytes'
              'shared',
              'snapshots',
              'system',
              'thin_provisioning',
              'total_physical',
              'total_provisioned',
              'total_reduction',
              'unique'
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
              'service_usec_per_read_op_cache_reduction',
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
        assert b"e" in rv.data
