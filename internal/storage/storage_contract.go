package storage

type Store interface {
	// регистрация
	IsRegistered(fromID int64) bool
	SetUser(fromID int64)

	// Профиль
	GetName(fromID int64) string
	AddName(fromID int64, name string)

	// ожидание
	IsWaitingName(fromID int64) bool
	SetWaitingName(fromID int64)
	ClearWaitingName(fromID int64)

	IsWaitingDistance(fromID int64) bool
	SetWaitingDistance(fromID int64)
	ClearWaitingDistance(fromID int64)

	// пробежки
	AddRun(fromID int64, dist float64)
	GetRuns(fromID int64) []float64
}