package service

import (
	"errors"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
)

func NameInput(store storage.Store, text string, fromID int64) (string, error) {
	name, err := ValidateName(text)

	if err != nil {
		return "", err
	}

	store.AddName(fromID, name)

	store.ClearWaitingName(fromID)

	return name, nil
}

func DistInput(store storage.Store, text string, fromID int64) (float64, error) {

	dist, err := ValidateDist(text)

	if err != nil {
		return 0, err
	}

	store.AddRun(fromID, dist)

	store.ClearWaitingDistance(fromID)

	return dist, nil
}

func CommandStart(store storage.Store, fromID int64) bool {
	if store.IsRegistered(fromID) {
		return true
	}

	store.SetWaitingName(fromID)

	return false
}

func GetStats(store storage.Store, fromID int64) (*storage.Stats, error) {

	userRuns := store.GetRuns(fromID)

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
