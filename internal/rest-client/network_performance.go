package client

type EthernetPerformance struct {
	OtherErrorsPerSec              float64 `json:"other_errors_per_sec"`
	ReceivedBytesPerSec            float64 `json:"received_bytes_per_sec"`
	ReceivedCrcErrorsPerSec        float64 `json:"received_crc_errors_per_sec"`
	ReceivedFrameErrorsPerSec      float64 `json:"received_frame_errors_per_sec"`
	ReceivedPacketsPerSec          float64 `json:"received_packets_per_sec"`
	TotalErrorsPerSec              float64 `json:"total_errors_per_sec"`
	TransmittedBytesPerSec         float64 `json:"transmitted_bytes_per_sec"`
	TransmittedCarrierErrorsPerSec float64 `json:"transmitted_carrier_errors_per_sec"`
	TransmittedDroppedErrorsPerSec float64 `json:"transmitted_dropped_errors_per_sec"`
	TransmittedPacketsPerSec       float64 `json:"transmitted_packets_per_sec"`
}

type FibreChannelPerformance struct {
	ReceivedBytesPerSec           float64 `json:"received_bytes_per_sec"`
	ReceivedCrcErrorsPerSec       float64 `json:"received_crc_errors_per_sec"`
	ReceivedFramesPerSec          float64 `json:"received_frames_per_sec"`
	ReceivedLinkFailuresPerSec    float64 `json:"received_link_failures_per_sec"`
	ReceivedLossOfSignalPerSec    float64 `json:"received_loss_of_signal_per_sec"`
	ReceivedLossOfSyncPerSec      float64 `json:"received_loss_of_sync_per_sec"`
	TotalErrorsPerSec             float64 `json:"total_errors_per_sec"`
	TransmittedBytesPerSec        float64 `json:"transmitted_bytes_per_sec"`
	TransmittedFramesPerSec       float64 `json:"transmitted_frames_per_sec"`
	TransmittedInvalidWordsPerSec float64 `json:"transmitted_invalid_words_per_sec"`
}

type NetworkInterfacePerformance struct {
	Name          string                  `json:"name"`
	Time          int                     `json:"time"`
	InterfaceType string                  `json:"interface_type"`
	Eth           EthernetPerformance     `json:"eth"`
	Fc            FibreChannelPerformance `json:"fc"`
}

type NetworkInterfacesPerformanceList struct {
	ContinuationToken  string                        `json:"continuation_token"`
	TotalItemCount     int                           `json:"total_item_count"`
	MoreItemsRemaining bool                          `json:"more_items_remaining"`
	Items              []NetworkInterfacePerformance `json:"items"`
}

func (fa *FAClient) GetNetworkInterfacesPerformance() *NetworkInterfacesPerformanceList {
	uri := "/network-interfaces/performance"
	result := new(NetworkInterfacesPerformanceList)
	res, err := fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if err != nil {
		fa.Error = err
	}

	if res.StatusCode() == 401 {
		fa.RefreshSession()
		_, err = fa.RestClient.R().
			SetResult(&result).
			Get(uri)
		if err != nil {
			fa.Error = err
		}
	}

	return result
}
