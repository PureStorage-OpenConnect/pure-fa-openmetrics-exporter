package client

type Subnet struct {
	Name string `json:"name"`
}

type EthernetSubtype struct {
	Name string `json:"name"`
}

type Ethernet struct {
	Address       string            `json:"address"`
	Gateway       string            `json:"gateway"`
	MacAddress    string            `json:"mac_address"`
	Mtu           int32             `json:"mtu"`
	Netmask       string            `json:"netmask"`
	Subtype       string            `json:"subtype"`
	Subinterfaces []EthernetSubtype `json:"transmitted_bytes_per_sec"`
	Subnet        Subnet            `json:"transmitted_carrier_errors_per_sec"`
	Vlan          int32             `json:"transmitted_dropped_errors_per_sec"`
}

type FibreChannel struct {
	Wwn string `json:"wwn"`
}

type NetworkInterface struct {
	Name          string       `json:"name"`
	Enabled       bool         `json:"enabled"`
	InterfaceType string       `json:"interface_type"`
	Services      []string     `json:"services"`
	Speed         int64        `json:"speed"`
	Eth           Ethernet     `json:"eth"`
	Fc            FibreChannel `json:"fc"`
}

type NetworkInterfacesList struct {
	ContinuationToken  string             `json:"continuation_token"`
	TotalItemCount     int                `json:"total_item_count"`
	MoreItemsRemaining bool               `json:"more_items_remaining"`
	Items              []NetworkInterface `json:"items"`
}

func (fa *FAClient) GetNetworkInterfaces() *NetworkInterfacesList {
	uri := "/network-interfaces"
	result := new(NetworkInterfacesList)
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
