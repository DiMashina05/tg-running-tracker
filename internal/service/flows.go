package service

import (
	"fmt"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
)

func OpenMe(store storage.Store, fromID int64) string {
	store.ClearWaitingDistance(fromID)

	if !store.IsRegistered(fromID) {
		return "Сначала зарегистрируйся: введи команду /start"
	}

	return fmt.Sprintf("Тебя зовут: %s, твой ID: %d\n", store.GetName(fromID), fromID) +
		"В будущем id будет использоваться для подачи заявок в друзья"
}

func OpenStats(store storage.Store, fromID int64) string {
	store.ClearWaitingDistance(fromID)

	if !store.IsRegistered(fromID) {
		return "Сначала зарегистрируйся: введи команду /start"
	}

	stats, err := GetStats(store, fromID)
	if err != nil {
		return "Произошла ошибка. Попробуй позже."
	}

	if stats.CountRuns == 0 {
		return "У тебя ещё не было тренировок"
	}

	return fmt.Sprintf("🏃 Пробежек: %d\n"+"📏 Суммарная дистанция: %.2f\n"+
		"📈 Средняя дистанция: %.2f\n"+"⬆️ Самая длинная дистанция: %.2f\n"+"⬇️ Самая короткая дистанция: %.2f",
		stats.CountRuns, stats.SumDistans, stats.Average, stats.MaxDist, stats.MinDist)
}

func OpenAddTraining(store storage.Store, fromID int64) string {
	if !store.IsRegistered(fromID) {
		return "Сначала зарегистрируйся: введи команду /start"
	}

	store.SetWaitingDistance(fromID)
	return "Сколько км пробежал?\nВведи число в километрах"
}

func OpenBack(store storage.Store, fromID int64) {
	store.ClearWaitingDistance(fromID)
}
