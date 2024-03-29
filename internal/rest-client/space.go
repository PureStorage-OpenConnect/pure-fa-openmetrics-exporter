package client

type Space struct {
	DataReduction      *float64 `json:"data_reduction"`
	Shared             *int64   `json:"shared"`
	Snapshots          *int64   `json:"snapshots"`
	System             *int64   `json:"system"`
	ThinProvisioning   *float64 `json:"thin_provisioning"`
	TotalPhysical      *int64   `json:"total_physical"`
	TotalProvisioned   *int64   `json:"total_provisioned"`
	TotalReduction     *float64 `json:"total_reduction"`
	Unique             *int64   `json:"unique"`
	Virtual            *int64   `json:"virtual"`
	UsedProvisioned    *int64   `json:"used_provisioned"`
	TotalUsed          *int64   `json:"total_used"`
	SharedEffective    *int64   `json:"shared_effective"`
	SnapshotsEffective *int64   `json:"snapshots_effective"`
	UniqueEffective    *int64   `json:"unique_effective"`
	TotalEffective     *int64   `json:"total_effective"`
	Replication        *int64   `json:"replication"`
}
