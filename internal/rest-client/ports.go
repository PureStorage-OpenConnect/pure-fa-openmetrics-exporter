package client

type Port struct {
	Name     string `json:"name"`
	Iqn      string `json:"iqn"`
	Nqn      string `json:"nqn"`
	Portal   string `json:"portal"`
	Wwn      string `json:"wwn"`
	Failover string `json:"failover"`
}

type PortsList struct {
	ContinuationToken  string `json:"continuation_token"`
	TotalItemCount     int32  `json:"total_item_count"`
	MoreItemsRemaining bool   `json:"more_items_remaining"`
	Items              []Port `json:"items"`
}

func (fa *FAClient) GetPorts() *PortsList {
	uri := "/ports"
	result := new(PortsList)
	res, err := fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if err != nil {
		fa.Error = err
	}

	if res.StatusCode() == 401 {
		fa.RefreshSession()
		_, err = fa.RestClient.R().
			SetResult(&result).
			Get(uri)
		if err != nil {
			fa.Error = err
		}
	}

	return result
}
