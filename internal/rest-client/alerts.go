package client

type Alert struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Actual           string `json:"actual"`
	Category         string `json:"category"`
	Closed           int64  `json:"closed"`
	Code             int64  `json:"code"`
	ComponentName    string `json:"component_name"`
	ComponentType    string `json:"component_type"`
	Created          int64  `json:"created"`
	Description      string `json:"description"`
	Expected         string `json:"expected"`
	Flagged          bool   `json:"flagged"`
	Issue            string `json:"issue"`
	KnowledgeBaseUrl string `json:"knowledge_base_url"`
	Notified         int64  `json:"notified"`
	Severity         string `json:"severity"`
	State            string `json:"state"`
	Summary          string `json:"summary"`
	Updated          int64  `json:"updated"`
}

type AlertsList struct {
	ContinuationToken  string  `json:"continuation_token"`
	TotalItemCount     int32   `json:"total_item_count"`
	MoreItemsRemaining bool    `json:"more_items_remaining"`
	Items              []Alert `json:"items"`
}

func (fa *FAClient) GetAlerts(filter string) *AlertsList {
	uri := "/alerts"
	result := new(AlertsList)
	req := fa.RestClient.R().SetResult(&result)
	if filter != "" {
		req = req.SetQueryParam("filter", filter)
	}
	res, _ := req.Get(uri)
	if res.StatusCode() == 401 {
		fa.RefreshSession()
		req.Get(uri)
	}

	return result
}
