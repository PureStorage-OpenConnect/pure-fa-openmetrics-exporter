package client


type PodPerformance struct {
	Id      string  `json:"id"`
	Name      string  `json:"name"`
	BytesPerMirroredWrite      int  `json:"bytes_per_mirrored_write"`
	BytesPerOp      int  `json:"bytes_per_op"`
	BytesPerRead      int  `json:"bytes_per_read"`
	BytesPerWrite      int  `json:"bytes_per_write"`
	MirroredWriteBytesPerSec      int  `json:"mirrored_write_bytes_per_sec"`
	MirroredWritesPerSec      int  `json:"mirrored_writes_per_sec"`
	QosRateLimitUsecPerMirroredWriteOp      int  `json:"qos_rate_limit_usec_per_mirrored_write_op"`
	QosRateLimitUsecPerReadOp      int  `json:"qos_rate_limit_usec_per_read_op"`
	QosRateLimitUsecPerWriteOp      int  `json:"qos_rate_limit_usec_per_write_op"`
	QueueUsecPerMirroredWriteOp      int  `json:"queue_usec_per_mirrored_write_op"`
	QueueUsecPerReadOp      int  `json:"queue_usec_per_read_op"`
	QueueUsecPerWriteOp      int  `json:"queue_usec_per_write_op"`
	ReadBytesPerSec      int  `json:"read_bytes_per_sec"`
	ReadsPerSec      int  `json:"reads_per_sec"`
	SanUsecPerMirroredWriteOp      int  `json:"san_usec_per_mirrored_write_op"`
	SanUsecPerReadOp      int  `json:"san_usec_per_read_op"`
	SanUsecPerWriteOp      int  `json:"san_usec_per_write_op"`
	ServiceUsecPerMirroredWriteOp      int  `json:"service_usec_per_mirrored_write_op"`
	ServiceUsecPerReadOp      int  `json:"service_usec_per_read_op"`
	ServiceUsecPerWriteOp      int  `json:"service_usec_per_write_op"`
	Time      int  `json:"time"`
	UsecPerMirroredWriteOp      int  `json:"usec_per_mirrored_write_op"`
	UsecPerReadOp      int  `json:"usec_per_read_op"`
	UsecPerWriteOp      int  `json:"usec_per_write_op"`
	WriteBytesPerSec      int  `json:"write_bytes_per_sec"`
	WritesPerSec      int  `json:"writes_per_sec"`
	ServiceUsecPerReadOpCacheReduction      int  `json:"service_usec_per_read_op_cache_reduction"`
	OthersPerSec      int  `json:"others_per_sec"`
	UsecPerOtherOp      int  `json:"usec_per_other_op"`
}

type PodsPerformanceList struct {
        ContinuationToken    string   `json:"continuation_token"`
        TotalItemCount       int      `json:"total_item_count"`
        MoreItemsRemaining   bool     `json:"more_items_remaining"`
        Items                []PodPerformance `json:"items"`
        Total                []PodPerformance `json:"total"`
}

func (fa *FAClient) GetPodsPerformance() *PodsPerformanceList {
        result := new(PodsPerformanceList)
        _, err := fa.RestClient.R().
                SetResult(&result).
                Get("/pods/performance")

        if err != nil {
                fa.Error = err
        }
        return result
}
