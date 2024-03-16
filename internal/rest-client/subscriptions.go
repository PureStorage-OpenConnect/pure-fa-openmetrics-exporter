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
	res, err := fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if err != nil {
		fa.Error = err
	}
	if res.StatusCode() == 401 {
		fa.RefreshSession()
	}
	res, err = fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if err != nil {
		fa.Error = err
	}

	return result
}
