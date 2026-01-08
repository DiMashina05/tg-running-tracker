package service

import (
	"errors"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
)

func NameInput(state *storage.State, text string, fromId int64) (string, error) {
	name, err := ValidateName(text)

	if err != nil {
		return "", err
	}

	state.SetUser(fromId)
	state.AddName(fromId, name)

	state.ClearWaitingName(fromId)

	return name, nil
}

func DistInput(state *storage.State, text string, fromId int64) (float64, error) {

	dist, err := ValidateDist(text)

	if err != nil {
		return 0, err
	}

	state.AddRun(fromId, dist)

	state.ClearWaitingDistance(fromId)

	return dist, nil
}

func CommandStart(state *storage.State, fromId int64) bool {
	if state.IsRegistered(fromId) {
		return true
	}

	state.SetWaitingName(fromId)

	return false
}

func GetStats(state *storage.State, fromId int64) (*storage.Stats, error) {

	userRuns := state.GetRuns(fromId)

	if len(userRuns) == 0 {
		return nil, errors.New("У тебя ещё не было тренировок")
	}

	countRuns := len(userRuns)

	sumDistans := userRuns[0]
	minDist := userRuns[0]
	maxDist := userRuns[0]

	for _, v := range userRuns[1:] {
		maxDist = max(maxDist, v)
		minDist = min(minDist, v)

		sumDistans += v
	}

	average := sumDistans / float64(countRuns)

	stats := &storage.Stats{CountRuns: countRuns, SumDistans: sumDistans, Average: average, MaxDist: maxDist, MinDist: minDist}

	return stats, nil
}
