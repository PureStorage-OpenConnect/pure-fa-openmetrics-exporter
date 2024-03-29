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
