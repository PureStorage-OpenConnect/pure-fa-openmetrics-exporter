package client

type VolumeGroup struct {
	Id                      string             `json:"id"`
	Name                    string             `json:"name"`
	QoS                     Qos                `json:"qos"`
	Pod                     PodShort           `json:"pod"`
}

type VolumeGroupsList struct {
	ContinuationToken  string        `json:"continuation_token"`
	TotalItemCount     int           `json:"total_item_count"`
	MoreItemsRemaining bool          `json:"more_items_remaining"`
	Items              []VolumeGroup `json:"items"`
	Total              []VolumeGroup `json:"total"`
}

func (fa *FAClient) GetVolumeGroups() *VolumeGroupsList {
	uri := "/volume-groups"
	result := new(VolumeGroupsList)
	res, _ := fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if res.StatusCode() == 401 {
		fa.RefreshSession()
		fa.RestClient.R().
			SetResult(&result).
			Get(uri)
	}
	return result
}
