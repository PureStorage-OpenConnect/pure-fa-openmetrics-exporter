package client

type PodReplicaLinksPerformance struct {
	Id                    string      `json:"id"`
	Time                  int         `json:"time"`
	Remotes               []ArrayTiny `json:"remotes"`
	BytesPerSecFromRemote float64     `json:"bytes_per_sec_from_remote"`
	BytesPerSecToRemote   float64     `json:"bytes_per_sec_to_remote"`
	BytesPerSecTotal      float64     `json:"bytes_per_sec_total"`
	LocalPod              PodShort    `json:"local_pod"`
	RemotePod             PodShort    `json:"remote_pod"`
	Direction             string      `json:"direction"`
}

type PodReplicaLinksPerformanceList struct {
	ContinuationToken  string                       `json:"continuation_token"`
	TotalItemCount     int                          `json:"total_item_count"`
	MoreItemsRemaining bool                         `json:"more_items_remaining"`
	Items              []PodReplicaLinksPerformance `json:"items"`
	Total              []PodReplicaLinksPerformance `json:"total"`
}

func (fa *FAClient) GetPodReplicaLinksPerformance() *PodReplicaLinksPerformanceList {
	uri := "/pod-replica-links/performance/replication"
	result := new(PodReplicaLinksPerformanceList)
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
