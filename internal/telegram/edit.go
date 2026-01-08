package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func EditMenu(bot *tgbotapi.BotAPI, chatId int64, messageID int) {
	edit := tgbotapi.NewEditMessageText(chatId, messageID, "Меню: ⬇")

	keyboard := getKeyboardMenu()

	edit.ReplyMarkup = &keyboard

	_, err := bot.Send(edit)
	if err != nil {
		log.Print(err)
	}
}

func EditBack(bot *tgbotapi.BotAPI, chatId int64, messageID int, text string) {
	edit := tgbotapi.NewEditMessageText(chatId, messageID, text)

	btnback := tgbotapi.NewInlineKeyboardButtonData("⬅ Назад", "back")
	row := tgbotapi.NewInlineKeyboardRow(btnback)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	edit.ReplyMarkup = &keyboard
	_, err := bot.Send(edit)
	if err != nil {
		log.Print(err)
	}
}
