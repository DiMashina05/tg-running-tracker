package service

import (
	"fmt"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
)

func OpenMe(state *storage.State, fromId int64) string {
	state.ClearWaitingDistance(fromId)

	return fmt.Sprintf("Тебя зовут: %s\n", state.GetName(fromId)) + "Невероятно полезная информация, да?"
}

func OpenStats(state *storage.State, fromId int64) string {
	state.ClearWaitingDistance(fromId)

	stats, err := GetStats(state, fromId)

	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("🏃 Пробежек: %d\n"+"📏 Суммарная дистанция: %.2f\n"+
		"📈 Средняя дистанция: %.2f\n"+"⬆️ Самая длинная дистанция: %.2f\n"+"⬇️ Самая короткая дистанция: %.2f",
		stats.CountRuns, stats.SumDistans, stats.Average, stats.MaxDist, stats.MinDist)
}

func OpenAddTraining(state *storage.State, fromId int64) string {
	state.SetWaitingDistance(fromId)

	return "Сколько км пробежал?\nВведи число в километрах"
}

func OpenBack(state *storage.State, fromId int64) {
	state.ClearWaitingDistance(fromId)
}
