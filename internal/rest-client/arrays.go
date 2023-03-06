package client

type Array struct {
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Banner            string            `json:"banner"`
	Capacity          float64           `json:"capacity"`
	ConsoleLock       bool              `json:"console_lock_enabled"`
	Encryption        Encryption        `json:"encryption"`
	EradicationConfig EradicationConfig `json:"eradication_config"`
	IdleTimeout       int               `json:"idle_timeout"`
	NtpServers        []string          `json:"ntp_servers"`
	Os                string            `json:"os"`
	Parity            float64           `json:"parity"`
	SCSITimeo         int               `json:"scsi_timeout"`
	Space             Space             `json:"space"`
	Version           string            `json:"version"`
}

type ArrayShort struct {
        Id                string     `json:"id"`
        Name              string     `json:"name"`
        FrozenAt          int        `json:"frozen_at"`
        MediatorStatus    string     `json:"mediator_status"`
        PreElected        bool       `json:"pre_elected"`
        Progress          float64    `json:"progress"`
        Status            string     `json:"status"`
}

type ArrayTiny struct {
        Id      string     `json:"id"`
        Name    string     `json:"name"`
}


type Encryption struct {
	DataAtRest    DataAtRest `json:"data_at_rest"`
	ModuleVersion string     `json:"module_version"`
}

type DataAtRest struct {
	Algorithm string `json:"algorithm"`
	Enabled   bool   `json:"enabled"`
}

type EradicationConfig struct {
	EradicationDelay  int    `json:"eradication_delay"`
	ManualEradication string `json:"manual_eradication"`
}

type ArraysList struct {
	ContinuationToken  string  `json:"continuation_token"`
	TotalItemCount     int     `json:"total_item_count"`
	MoreItemsRemaining bool    `json:"more_items_remaining"`
	Items              []Array `json:"items"`
}

func (fa *FAClient) GetArrays() *ArraysList {
	uri := "/arrays"
	result := new(ArraysList)
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
