package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func getKeyboardMenu() tgbotapi.InlineKeyboardMarkup {
	btnme := tgbotapi.NewInlineKeyboardButtonData("👤 Профиль", "me")
	btnStats := tgbotapi.NewInlineKeyboardButtonData("📊 Статистика", "stats")

	row1 := tgbotapi.NewInlineKeyboardRow(btnme, btnStats)

	btnAdd := tgbotapi.NewInlineKeyboardButtonData("🏃 Добавить тренировку", "add_training")

	row2 := tgbotapi.NewInlineKeyboardRow(btnAdd)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)
	return keyboard
}
