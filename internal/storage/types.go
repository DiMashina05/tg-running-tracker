package storage

type State struct {
	users           map[int64]bool
	names           map[int64]string
	waitingName     map[int64]bool
	waitingDistance map[int64]bool
	runs            map[int64][]float64
}

func NewState() *State {
	return &State{
		users:           make(map[int64]bool),
		names:           make(map[int64]string),
		waitingName:     make(map[int64]bool),
		waitingDistance: make(map[int64]bool),
		runs:            make(map[int64][]float64),
	}
}

type Stats struct {
	CountRuns  int
	SumDistans float64
	Average    float64
	MaxDist    float64
	MinDist    float64
}

