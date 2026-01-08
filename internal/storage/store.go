package storage

func (state *State) AddRun(fromId int64, dist float64) {
	state.runs[fromId] = append(state.runs[fromId], dist)
}

func (state *State) GetRuns(fromId int64) []float64 {
	out := make([]float64, len(state.runs[fromId]))
	copy(out, state.runs[fromId])
	return out
}

func (state *State) GetName(fromId int64) string {
	return state.names[fromId]
}

func (state *State) AddName(fromId int64, name string) {
	state.names[fromId] = name
}

func (state *State) SetWaitingName(fromId int64) {
	state.waitingName[fromId] = true
}

func (state *State) ClearWaitingName(fromId int64) {
	delete(state.waitingName, fromId)
}

func (state *State) ClearWaitingDistance(fromId int64) {
	delete(state.waitingDistance, fromId)
}

func (state *State) SetWaitingDistance(fromId int64) {
	state.waitingDistance[fromId] = true
}

func (state *State) SetUser(fromId int64) {
	state.users[fromId] = true
}
