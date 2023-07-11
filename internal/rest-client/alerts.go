package client


type Alert struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Actual        string    `json:"actual"`
	Closed        int       `json:"closed"`
	Code          int       `json:"code"`
	ComponentName string    `json:"component_name"`
	ComponentType string    `json:"component_type"`
	Created       int       `json:"created"`
	Description   string    `json:"description"`
	Expected      string    `json:"expected"`
	Flagged       bool      `json:"flagged"`
	Issue         string    `json:"issue"`
	Index         int       `json:"index"`
	KnowledgeBaseUrl         string    `json:"knowledge_base_url"`
	Notified      int       `json:"notified"`
	Severity      string    `json:"severity"`
	State         string    `json:"state"`
	Summary       string    `json:"summary"`
	Updated       int       `json:"updated"`
}

type AlertsList struct {
        ContinuationToken    string   `json:"continuation_token"`
        TotalItemCount       int      `json:"total_item_count"`
        MoreItemsRemaining   bool     `json:"more_items_remaining"`
	Items                []Alert  `json:"items"`
}

func (fa *FAClient) GetAlerts(filter string) *AlertsList {
	uri := "/alerts"
	result := new(AlertsList)
	req := fa.RestClient.R().SetResult(&result)
	if filter != "" {
		req = req.SetQueryParam("filter", filter)
	}
	res, err := req.Get(uri)
	if err != nil {
		fa.Error = err
	}
	if res.StatusCode() == 401 {
		fa.RefreshSession()
	}
	res, err = req.Get(uri)
	if err != nil {
		fa.Error = err
	}
	
	return result
}
