package client


type Qos struct  {
	BandwidthLimit  int   `json:"bandwidth_limit"`
	IopsLimit       int   `json:"iops_limit"`
}

type PriorityAdjustment struct {
	PriorityAdjustmentOperator   int  `json:string","priority_adjustment_operator"`
	PriorityAdjustmentValue      int  `json:"priority_adjustment_value"`
}

type PodShort struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
}

type Source struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
}

type VolumeGroupShort struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
}

type Volume struct {
	Id                       string              `json:"id"`
	Name                     string              `json:"name"`
	ConnectionCount          int                 `json:"connection_count"`
	Created                  int                 `json:"created"`
	Destroyed                bool                `json:"destroyed"`
	HostEncryptionKeyStatus  string              `json:"host_encryption_key_status"`
	PriorityAdjustment       PriorityAdjustment  `json:"priority_adjustment"`
	Provisioned              int                 `json:"provisioned"`
	Serial                   string              `json:"serial"`
	Space                    Space               `json:"space"`
	TimeRemaining            int                 `json:"time_remaining"`
	Pod                      PodShort            `json:"pod"`
	Source                   Source              `json:"source"`
	Subtype                  string              `json:"subtype"`
	VolumeGroup              VolumeGroupShort    `json:"volume_group"`
	RequestedPromotionState  string              `json:"requested_promotion_state"`
	PromotionStatus          string              `json:"promotion_status"`
	Priority                 int                 `json:"priority"`
}

type VolumesList struct {
        ContinuationToken    string   `json:"continuation_token"`
        TotalItemCount       int      `json:"total_item_count"`
        MoreItemsRemaining   bool     `json:"more_items_remaining"`
        Items                []Volume `json:"items"`
        Total                []Volume `json:"total"`
}

func (fa *FAClient) GetVolumes() *VolumesList {
        result := new(VolumesList)
        _, err := fa.RestClient.R().
                SetResult(&result).
                Get("/volumes")

        if err != nil {
                fa.Error = err
        }
        return result
}
