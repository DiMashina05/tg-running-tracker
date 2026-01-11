package service

import (
	"fmt"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
)

func OpenMe(store storage.Store, fromId int64) string {
	store.ClearWaitingDistance(fromId)

	return fmt.Sprintf("Тебя зовут: %s\n", store.GetName(fromId)) + "Невероятно полезная информация, да?"
}

func OpenStats(store storage.Store, fromId int64) string {
	store.ClearWaitingDistance(fromId)

	stats, err := GetStats(store, fromId)

	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("🏃 Пробежек: %d\n"+"📏 Суммарная дистанция: %.2f\n"+
		"📈 Средняя дистанция: %.2f\n"+"⬆️ Самая длинная дистанция: %.2f\n"+"⬇️ Самая короткая дистанция: %.2f",
		stats.CountRuns, stats.SumDistans, stats.Average, stats.MaxDist, stats.MinDist)
}

func OpenAddTraining(store storage.Store, fromId int64) string {
	store.SetWaitingDistance(fromId)

	return "Сколько км пробежал?\nВведи число в километрах"
}

func OpenBack(store storage.Store, fromId int64) {
	store.ClearWaitingDistance(fromId)
}
