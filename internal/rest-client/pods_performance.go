package client


type PodPerformance struct {
	Id      string  `json:"id"`
	Name      string  `json:"name"`
	BytesPerMirroredWrite      float64  `json:"bytes_per_mirrored_write"`
	BytesPerOp      float64  `json:"bytes_per_op"`
	BytesPerRead      float64  `json:"bytes_per_read"`
	BytesPerWrite      float64  `json:"bytes_per_write"`
	MirroredWriteBytesPerSec      float64  `json:"mirrored_write_bytes_per_sec"`
	MirroredWritesPerSec      float64  `json:"mirrored_writes_per_sec"`
	QosRateLimitUsecPerMirroredWriteOp      float64  `json:"qos_rate_limit_usec_per_mirrored_write_op"`
	QosRateLimitUsecPerReadOp      float64  `json:"qos_rate_limit_usec_per_read_op"`
	QosRateLimitUsecPerWriteOp      float64  `json:"qos_rate_limit_usec_per_write_op"`
	QueueUsecPerMirroredWriteOp      float64  `json:"queue_usec_per_mirrored_write_op"`
	QueueUsecPerReadOp      float64  `json:"queue_usec_per_read_op"`
	QueueUsecPerWriteOp      float64  `json:"queue_usec_per_write_op"`
	ReadBytesPerSec      float64  `json:"read_bytes_per_sec"`
	ReadsPerSec      float64  `json:"reads_per_sec"`
	SanUsecPerMirroredWriteOp      float64  `json:"san_usec_per_mirrored_write_op"`
	SanUsecPerReadOp      float64  `json:"san_usec_per_read_op"`
	SanUsecPerWriteOp      float64  `json:"san_usec_per_write_op"`
	ServiceUsecPerMirroredWriteOp      float64  `json:"service_usec_per_mirrored_write_op"`
	ServiceUsecPerReadOp      float64  `json:"service_usec_per_read_op"`
	ServiceUsecPerWriteOp      float64  `json:"service_usec_per_write_op"`
	Time      int  `json:"time"`
	UsecPerMirroredWriteOp      float64  `json:"usec_per_mirrored_write_op"`
	UsecPerReadOp      float64  `json:"usec_per_read_op"`
	UsecPerWriteOp      float64  `json:"usec_per_write_op"`
	WriteBytesPerSec      float64  `json:"write_bytes_per_sec"`
	WritesPerSec      float64  `json:"writes_per_sec"`
	ServiceUsecPerReadOpCacheReduction      float64  `json:"service_usec_per_read_op_cache_reduction"`
	OthersPerSec      float64  `json:"others_per_sec"`
	UsecPerOtherOp      float64  `json:"usec_per_other_op"`
}

type PodsPerformanceList struct {
        ContinuationToken    string   `json:"continuation_token"`
        TotalItemCount       int      `json:"total_item_count"`
        MoreItemsRemaining   bool     `json:"more_items_remaining"`
        Items                []PodPerformance `json:"items"`
        Total                []PodPerformance `json:"total"`
}

func (fa *FAClient) GetPodsPerformance() *PodsPerformanceList {
	uri := "/pods/performance"
        result := new(PodsPerformanceList)
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
