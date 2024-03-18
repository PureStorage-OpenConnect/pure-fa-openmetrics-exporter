package client

type DirectoryPerformance struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	BytesPerOp       float64 `json:"bytes_per_op"`
	BytesPerRead     float64 `json:"bytes_per_read"`
	BytesPerWrite    float64 `json:"bytes_per_write"`
	OthersPerSec     float64 `json:"others_per_sec"`
	ReadBytesPerSec  float64 `json:"read_bytes_per_sec"`
	ReadsPerSec      float64 `json:"reads_per_sec"`
	Time             int     `json:"time"`
	UsecPerOtherOp   float64 `json:"usec_per_other_op"`
	UsecPerReadOp    float64 `json:"usec_per_read_op"`
	UsecPerWriteOp   float64 `json:"usec_per_write_op"`
	WriteBytesPerSec float64 `json:"write_bytes_per_sec"`
	WritesPerSec     float64 `json:"writes_per_sec"`
}

type DirectoriesPerformanceList struct {
	ContinuationToken  string                 `json:"continuation_token"`
	TotalItemCount     int                    `json:"total_item_count"`
	MoreItemsRemaining bool                   `json:"more_items_remaining"`
	Items              []DirectoryPerformance `json:"items"`
	Total              []DirectoryPerformance `json:"total"`
}

func (fa *FAClient) GetDirectoriesPerformance() *DirectoriesPerformanceList {
	uri := "/directories/performance"
	result := new(DirectoriesPerformanceList)
	res, err := fa.RestClient.R().
		SetResult(&result).
		Get(uri)
	if err != nil {
		fa.Error = err
	}
	if res.StatusCode() == 401 {
		fa.RefreshSession()
		fa.RestClient.R().
			SetResult(&result).
			Get(uri)
	}

	return result
}
