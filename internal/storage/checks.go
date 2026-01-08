package storage

func (state *State) IsRegistered(fromId int64) bool {
	return state.users[fromId]
}

func (state *State) IsWaitingDistance(fromId int64) bool {
	return state.waitingDistance[fromId]
}

func (state *State) IsWaitingName(fromId int64) bool {
	return state.waitingName[fromId]
}
