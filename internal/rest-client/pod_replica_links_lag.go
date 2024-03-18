package client

type Lag struct {
	Avg float64 `json:"avg"`
	Max float64 `json:"max"`
}

type PodReplicaLinksLag struct {
	Id            string      `json:"id"`
	Lag           Lag         `json:"lag"`
	Time          int         `json:"time"`
	Remotes       []ArrayTiny `json:"remotes"`
	LocalPod      PodShort    `json:"local_pod"`
	RemotePod     PodShort    `json:"remote_pod"`
	Direction     string      `json:"direction"`
	Status        string      `json:"status"`
	RecoveryPoint int         `json:"recovery_point"`
}

type PodReplicaLinksLagList struct {
	ContinuationToken  string               `json:"continuation_token"`
	TotalItemCount     int                  `json:"total_item_count"`
	MoreItemsRemaining bool                 `json:"more_items_remaining"`
	Items              []PodReplicaLinksLag `json:"items"`
}

func (fa *FAClient) GetPodReplicaLinksLag() *PodReplicaLinksLagList {
	uri := "/pod-replica-links/lag"
	result := new(PodReplicaLinksLagList)
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
