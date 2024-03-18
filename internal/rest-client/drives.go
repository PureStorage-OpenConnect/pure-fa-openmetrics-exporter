package client

type Drive struct {
	Name     string  `json:"name"`
	Details  string  `json:"details"`
	Capacity float64 `json:"capacity"`
	Protocol string  `json:"protocol"`
	Status   string  `json:"status"`
	Type     string  `json:"type"`
}

type DriveList struct {
	ContinuationToken  string  `json:"continuation_token"`
	TotalItemCount     int     `json:"total_item_count"`
	MoreItemsRemaining bool    `json:"more_items_remaining"`
	Items              []Drive `json:"items"`
}

func (fa *FAClient) GetDrives() *DriveList {
	uri := "/drives"
	result := new(DriveList)
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
