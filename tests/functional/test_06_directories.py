import pytest


def test_directories(client):
    rv = client.get("/metrics/directories")
    for e in ['purefa_directory_space_datareduction_ratio',
              'purefa_directory_space_size_bytes',
              'purefa_directory_space_bytes',
              'name',
              'filesystem',
              'path',
              'space',
              'shared',
              'snapshots',
              'system',
              'thin_provisioning',
              'total_physical',
              'total_provisioned',
              'total_reduction',
              'unique',
              'purefa_directory_performance_latency_usec',
              'purefa_directory_performance_bandwidth_bytes',
              'purefa_directory_performance_iops',
              'purefa_directory_performance_avg_block_bytes'
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
        assert b"e" in rv.data

