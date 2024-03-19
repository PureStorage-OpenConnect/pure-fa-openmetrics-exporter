package client

type Pod struct {
	Id                      string            `json:"id"`
	Name                    string            `json:"name"`
	Arrays                  []ArrayShort      `json:"arrays"`
	Destroyed               bool              `json:"destroyed"`
	FailoverPreferences     []Array           `json:"failover_preferences"`
	Footprint               int               `json:"footprint"`
	Mediator                string            `json:"mediator"`
	MediatorVersion         string            `json:"mediator_version"`
	Source                  Source            `json:"source"`
	Space                   Space             `json:"space"`
	TimeRemaining           int               `json:"time_remaining"`
	RequestedPromotionState string            `json:"requested_promotion_state"`
	PromotionStatus         string            `json:"promotion_status"`
	LinkSourceCount         int               `json:"link_source_count"`
	LinkTargetCount         int               `json:"link_target_count"`
	ArrayCount              int               `json:"array_count"`
	EradicationConfig       EradicationConfig `json:"eradication_config"`
}

type PodsList struct {
	ContinuationToken  string `json:"continuation_token"`
	TotalItemCount     int    `json:"total_item_count"`
	MoreItemsRemaining bool   `json:"more_items_remaining"`
	Items              []Pod  `json:"items"`
	Total              []Pod  `json:"total"`
}

type PodShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (fa *FAClient) GetPods() *PodsList {
	uri := "/pods"
	result := new(PodsList)
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
