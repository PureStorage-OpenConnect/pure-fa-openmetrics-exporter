package client

type Controllers struct {
	Name      string `json:"name"`
	Mode      string `json:"mode"`
	Model     string `json:"model"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Version   string `json:"version"`
	ModeSince int64  `json:"mode_since"`
}

type ControllersList struct {
	ContinuationToken  string        `json:"continuation_token"`
	TotalItemCount     int32         `json:"total_item_count"`
	MoreItemsRemaining bool          `json:"more_items_remaining"`
	Items              []Controllers `json:"items"`
}

func (fa *FAClient) GetControllers() *ControllersList {
	uri := "/controllers"
	result := new(ControllersList)
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
