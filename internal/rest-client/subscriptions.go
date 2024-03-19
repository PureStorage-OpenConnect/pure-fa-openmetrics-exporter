package client

type Subscription struct {
	Id      string `json:"id"`
	Service string `json:"service"`
}

type SubscriptionsList struct {
	ContinuationToken  string         `json:"continuation_token"`
	TotalItemCount     int32          `json:"total_item_count"`
	MoreItemsRemaining bool           `json:"more_items_remaining"`
	Items              []Subscription `json:"items"`
}

func (fa *FAClient) GetSubscriptions() *SubscriptionsList {
	uri := "/subscriptions"
	result := new(SubscriptionsList)
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
