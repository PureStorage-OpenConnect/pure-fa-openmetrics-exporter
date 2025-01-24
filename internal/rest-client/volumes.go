package client

type Qos struct {
	BandwidthLimit *int64 `json:"bandwidth_limit"`
	IopsLimit      *int64 `json:"iops_limit"`
}

type PriorityAdjustment struct {
	PriorityAdjustmentOperator int `json:"priority_adjustment_operator"`
	PriorityAdjustmentValue    int `json:"priority_adjustment_value"`
}

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type VolumeGroupShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Volume struct {
	Id                      string             `json:"id"`
	Name                    string             `json:"name"`
	ConnectionCount         int                `json:"connection_count"`
	Created                 int                `json:"created"`
	Destroyed               bool               `json:"destroyed"`
	HostEncryptionKeyStatus string             `json:"host_encryption_key_status"`
	PriorityAdjustment      PriorityAdjustment `json:"priority_adjustment"`
	Provisioned             int                `json:"provisioned"`
	QoS                     Qos                `json:"qos"`
	Serial                  string             `json:"serial"`
	Space                   Space              `json:"space"`
	TimeRemaining           int                `json:"time_remaining"`
	Pod                     PodShort           `json:"pod"`
	Source                  Source             `json:"source"`
	Subtype                 string             `json:"subtype"`
	VolumeGroup             VolumeGroupShort   `json:"volume_group"`
	RequestedPromotionState string             `json:"requested_promotion_state"`
	PromotionStatus         string             `json:"promotion_status"`
	Priority                int                `json:"priority"`
}

type VolumesList struct {
	ContinuationToken  string   `json:"continuation_token"`
	TotalItemCount     int      `json:"total_item_count"`
	MoreItemsRemaining bool     `json:"more_items_remaining"`
	Items              []Volume `json:"items"`
	Total              []Volume `json:"total"`
}

func (fa *FAClient) GetVolumes() *VolumesList {
	uri := "/volumes"
	result := new(VolumesList)
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
