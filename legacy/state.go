package legacy

//хранилище в памяти компьютера

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


// переписанные методы
func (state *State) SetUser(fromId int64) {
	state.users[fromId] = true
}

func (state *State) AddName(fromId int64, name string) {
	state.names[fromId] = name
}

func (state *State) GetName(fromId int64) string {
	return state.names[fromId]
}

func (state *State) SetWaitingName(fromId int64) {
	state.waitingName[fromId] = true
}

func (state *State) ClearWaitingName(fromId int64) {
	delete(state.waitingName, fromId)
}

func (state *State) AddRun(fromId int64, dist float64) {
	state.runs[fromId] = append(state.runs[fromId], dist)
}

func (state *State) ClearWaitingDistance(fromId int64) {
	delete(state.waitingDistance, fromId)
}

func (state *State) SetWaitingDistance(fromId int64) {
	state.waitingDistance[fromId] = true
}

func (state *State) GetRuns(fromId int64) []float64 {
	out := make([]float64, len(state.runs[fromId]))
	copy(out, state.runs[fromId])
	return out
}


//переписанные методы
func (state *State) IsRegistered(fromId int64) bool {
	return state.users[fromId]
}

func (state *State) IsWaitingDistance(fromId int64) bool {
	return state.waitingDistance[fromId]
}

func (state *State) IsWaitingName(fromId int64) bool {
	return state.waitingName[fromId]
}