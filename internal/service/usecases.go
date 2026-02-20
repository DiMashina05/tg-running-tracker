package service

import (
	"fmt"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
)

func NameInput(store storage.Store, text string, fromID int64) (string, error) {

	name, err := SetName(store, text, fromID)
	if err == nil {
		store.ClearWaitingName(fromID)
	}

	return name, err
}

func SetName(store storage.Store, text string, fromID int64) (string, error) {
	if store.IsRegistered(fromID) {
		return "", ErrAlreadyRegistered
	}

	name, err := ValidateName(text)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrInvalidName, err)
	}

	store.AddName(fromID, name)
	return name, nil
}

func DistInput(store storage.Store, text string, fromID int64) (float64, error) {

	dist, err := AddRun(store, text, fromID)
	if err == nil {
		store.ClearWaitingDistance(fromID)
	}

	return dist, err
}

func AddRun(store storage.Store, text string, fromID int64) (float64, error) {
	if !store.IsRegistered(fromID) {
		return 0, ErrNotRegistered
	}

	dist, err := ValidateDist(text)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidDistance, err)
	}

	store.AddRun(fromID, dist)
	return dist, nil
}

func CommandStart(store storage.Store, fromID int64) error {
	if store.IsRegistered(fromID) {
		return ErrAlreadyRegistered
	}

	store.SetWaitingName(fromID)
	return nil
}

func GetStats(store storage.Store, fromID int64) (*storage.Stats, error) {

	if !store.IsRegistered(fromID) {
		return nil, ErrNotRegistered
	}
	userRuns := store.GetRuns(fromID)

	if len(userRuns) == 0 {
		return &storage.Stats{
			CountRuns:  0,
			SumDistans: 0,
			Average:    0,
			MaxDist:    0,
			MinDist:    0,
		}, nil
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

	return &storage.Stats{
		CountRuns:  countRuns,
		SumDistans: sumDistans,
		Average:    average,
		MaxDist:    maxDist,
		MinDist:    minDist,
	}, nil
}
