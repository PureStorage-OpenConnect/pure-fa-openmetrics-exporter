package client

type Space struct {
	DataReduction        float64 `json:"data_reduction"`
	Shared               float64 `json:"shared"`
	Snapshots            float64 `json:"snapshots"`
	System               float64 `json:"system"`
	ThinProvisioning     float64 `json:"thin_provisioning"`
	TotalPhysical        float64 `json:"total_physical"`
	TotalProvisioned     float64 `json:"total_provisioned"`
	TotalReduction       float64 `json:"total_reduction"`
	Unique               float64 `json:"unique"`
	Virtual              float64 `json:"virtual"`
	Replication          float64 `json:"replication"`
        SharedEffective      float64 `json:"shared_effective"`
        SnapshotsEffective   float64 `json:"snapshots_effective"`
        UniqueEffective      float64  `json:"unique_effective"`
        TotalEffective       float64  `json:"total_effective"`
}
