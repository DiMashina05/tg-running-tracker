package service

import "errors"

var(
	ErrNotRegistered = errors.New("Not registered")
	ErrAlreadyRegistered = errors.New("Already registered")
	ErrInvalidName = errors.New("Invalid name")
	ErrInvalidDistance = errors.New("Invalid distance")
	ErrNoRuns = errors.New("No runs")
)