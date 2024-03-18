package client

type Chap struct {
	HostPassword   string `json:"host_password"`
	HostUser       string `json:"host_user"`
	TargetPassword string `json:"target_password"`
	TargetUser     string `json:"target_user"`
}

type PortConnectivity struct {
	Details string `json:"details"`
	Status  string `json:"status"`
}

type Host struct {
	Name             string           `json:"name"`
	Chap             Chap             `json:"chap"`
	ConnectionCount  int              `json:"connection_count"`
	HostGroup        HostGroupShort   `json:"host_group"`
	IQNs             []string         `json:"iqns"`
	NQNs             []string         `json:"nqns"`
	Personality      string           `json:"personality"`
	PortConnectivity PortConnectivity `json:"port_connectivity"`
	Space            Space            `json:"space"`
	PreferredArrays  []ArrayShort     `json:"preferred_arrays"`
	WWNs             []string         `json:"wwns"`
	IsLocal          bool             `json:"is_local"`
}

type HostsList struct {
	ContinuationToken  string `json:"continuation_token"`
	TotalItemCount     int    `json:"total_item_count"`
	MoreItemsRemaining bool   `json:"more_items_remaining"`
	Items              []Host `json:"items"`
}

func (fa *FAClient) GetHosts() *HostsList {
	uri := "/hosts"
	result := new(HostsList)
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
