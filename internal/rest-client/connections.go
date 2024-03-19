package client

type HostShort struct {
	Name string `json:"name"`
}

type HostGroupShort struct {
	Name string `json:"name"`
}

type ProtocolEndpoint struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type VolumeShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Connection struct {
	Host             HostShort        `json:"host"`
	HostGroup        HostGroupShort   `json:"host_group"`
	Lun              int              `json:"lun"`
	ProtocolEndpoint ProtocolEndpoint `json:"protocol_endpoint"`
	Volume           VolumeShort      `json:"volume"`
}

type ConnectionsList struct {
	ContinuationToken  string       `json:"continuation_token"`
	TotalItemCount     int          `json:"total_item_count"`
	MoreItemsRemaining bool         `json:"more_items_remaining"`
	Items              []Connection `json:"items"`
}

func (fa *FAClient) GetConnections() *ConnectionsList {
	uri := "/connections"
	result := new(ConnectionsList)
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
