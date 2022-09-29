package client


type DirectoryPerformance struct {
	Id      string  `json:"id"`
	Name      string  `json:"name"`
	BytesPerOp      int  `json:"bytes_per_op"`
	BytesPerRead      int  `json:"bytes_per_read"`
	BytesPerWrite      int  `json:"bytes_per_write"`
	OthersPerSec      int  `json:"others_per_sec"`
	ReadBytesPerSec      int  `json:"read_bytes_per_sec"`
	ReadsPerSec      int  `json:"reads_per_sec"`
	Time      int  `json:"time"`
	UsecPerOtherOp      int  `json:"usec_per_other_op"`
	UsecPerReadOp      int  `json:"usec_per_read_op"`
	UsecPerWriteOp      int  `json:"usec_per_write_op"`
	WriteBytesPerSec      int  `json:"write_bytes_per_sec"`
	WritesPerSec      int  `json:"writes_per_sec"`
}

type DirectoriesPerformanceList struct {
        ContinuationToken    string   `json:"continuation_token"`
        TotalItemCount       int      `json:"total_item_count"`
        MoreItemsRemaining   bool     `json:"more_items_remaining"`
        Items                []DirectoryPerformance `json:"items"`
        Total                []DirectoryPerformance `json:"total"`
}

func (fa *FAClient) GetDirectoriesPerformance() *DirectoriesPerformanceList {
        result := new(DirectoriesPerformanceList)
        _, err := fa.RestClient.R().
                SetResult(&result).
                Get("/directories/performance")

        if err != nil {
                fa.Error = err
        }
        return result
}
