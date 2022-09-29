package client


type Hardware struct {
	Name          string    `json:"name"`
	Details       string    `json:"details"`
	IdentityEnabled bool    `json:"identity_enabled"`
	Index         int       `json:"index"`
	Model         string    `json:"model"`
	Serial        string    `json:"serial"`
	Slot          int       `json:"slot"`
	Speed         int       `json:"speed"`
	Status        string    `json:"status"`
	Temperature   int       `json:"temperature"`
	Type          string    `json:"type"`
	Voltage       int       `json:"voltage"`
}

type HardwareList struct {
        ContinuationToken    string   `json:"continuation_token"`
        TotalItemCount       int      `json:"total_item_count"`
        MoreItemsRemaining   bool     `json:"more_items_remaining"`
	Items                []Alert  `json:"items"`
}

func (fa *FAClient) GetHardware() *HardwareList {
	result := new(HardwareList)
	_, err := fa.RestClient.R().
		SetResult(&result).
		Get("/hardware")
	if err != nil {
		fa.Error = err
	}
	return result
}
