package storage

type Stats struct {
	CountRuns  int     `json:"count_runs"`
	SumDistans float64 `json:"sum_distance"`
	Average    float64 `json:"average"`
	MaxDist    float64 `json:"max_dist"`
	MinDist    float64 `json:"min_dist"`
}
