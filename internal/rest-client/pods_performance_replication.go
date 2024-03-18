package client

type PerformanceReplication struct {
	FromRemoteBytesPerSec float64 `json:"from_remote_bytes_per_sec"`
	ToRemoteBytesPerSec   float64 `json:"to_remote_bytes_per_sec"`
	TotalBytesPerSec      float64 `json:"total_bytes_per_sec"`
}

type PodPerformanceReplication struct {
	Pod                   PodShort               `json:"pod"`
	Time                  int                    `json:"time"`
	ContinuousBytesPerSec PerformanceReplication `json:"continuous_bytes_per_sec"`
	ResyncBytesPerSec     PerformanceReplication `json:"resync_bytes_per_sec"`
	SyncBytesPerSec       PerformanceReplication `json:"sync_bytes_per_sec"`
	PeriodicBytesPerSec   PerformanceReplication `json:"periodic_bytes_per_sec"`
	TotalBytesPerSec      float64                `json:"total_bytes_per_sec"`
}

type PodsPerformanceReplicationList struct {
	ContinuationToken  string                      `json:"continuation_token"`
	TotalItemCount     int                         `json:"total_item_count"`
	MoreItemsRemaining bool                        `json:"more_items_remaining"`
	Items              []PodPerformanceReplication `json:"items"`
	Total              []PodPerformanceReplication `json:"total"`
}

func (fa *FAClient) GetPodsPerformanceReplication() *PodsPerformanceReplicationList {
	uri := "/pods/performance/replication"
	result := new(PodsPerformanceReplicationList)
	res, err := fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if err != nil {
		fa.Error = err
	}
	if res.StatusCode() == 401 {
		fa.RefreshSession()
		fa.RestClient.R().
			SetResult(&result).
			Get(uri)
	}

	return result
}
