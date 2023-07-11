package client


type Ethernet struct {
	OtherErrorsPerSec		float64    `json:"other_errors_per_sec"`
	ReceivedBytesPerSec		float64    `json:"received_bytes_per_sec"`
	ReceivedCrcErrorsPerSec		float64    `json:"received_crc_errors_per_sec"`
	ReceivedFrameErrorsPerSec	float64    `json:"received_frame_errors_per_sec"`
	ReceivedPacketsPerSec		float64    `json:"received_packets_per_sec"`
	TotalErrorsPerSec		float64    `json:"total_errors_per_sec"`
	TransmittedBytesPerSec		float64    `json:"transmitted_bytes_per_sec"`
	TransmittedCarrierErrorsPerSec	float64    `json:"transmitted_carrier_errors_per_sec"`
	TransmittedDroppedErrorsPerSec	float64    `json:"transmitted_dropped_errors_per_sec"`
	TransmittedPacketsPerSec	float64    `json:"transmitted_packets_per_sec"`
}

type FibreChannel struct {
	ReceivedBytesPerSec		float64    `json:"received_bytes_per_sec"`
	ReceivedCrcErrorsPerSec		float64    `json:"received_crc_errors_per_sec"`
	ReceivedFramesPerSec		float64    `json:"received_frames_per_sec"`
	ReceivedLinkFailuresPerSec	float64    `json:"received_link_failures_per_sec"`
	ReceivedLossOfSignalPerSec	float64    `json:"received_loss_of_signal_per_sec"`
	ReceivedLossOfSyncPerSec	float64    `json:"received_loss_of_sync_per_sec"`
	TotalErrorsPerSec		float64    `json:"total_errors_per_sec"`
	TransmittedBytesPerSec		float64    `json:"transmitted_bytes_per_sec"`
	TransmittedFramesPerSec		float64    `json:"transmitted_frames_per_sec"`
	TransmittedInvalidWordsPerSec	float64    `json:"transmitted_invalid_words_per_sec"`
}

type NetworkInterface struct {
	Name            string        `json:"name"`
	Time            int           `json:"time"`
	InterfaceType   string        `json:"interface_type"`
	Eth             Ethernet      `json:"eth"`
	Fc              FibreChannel  `json:"fc"`
}

type NetworkInterfacesList struct {
        ContinuationToken    string               `json:"continuation_token"`
        TotalItemCount       int                  `json:"total_item_count"`
        MoreItemsRemaining   bool                 `json:"more_items_remaining"`
	Items                []NetworkInterface   `json:"items"`
}

func (fa *FAClient) GetNetworkInterfacesPerformance() *NetworkInterfacesList {
	uri := "/network-interfaces/performance"
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
