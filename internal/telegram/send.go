package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendText(bot *tgbotapi.BotAPI, chatId int64, str string) {
	msg := tgbotapi.NewMessage(chatId, str)
	_, err := bot.Send(msg)
	if err != nil {
		log.Print(err)
	}
}

func SendMenu(bot *tgbotapi.BotAPI, chatId int64) {

	keyboard := getKeyboardMenu()

	SendWithMarkup(bot, chatId, "Меню: ⬇", keyboard)
}

func SendBack(bot *tgbotapi.BotAPI, chatId int64, text string) {
	btnback := tgbotapi.NewInlineKeyboardButtonData("⬅ Назад", "back")
	row := tgbotapi.NewInlineKeyboardRow(btnback)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	SendWithMarkup(bot, chatId, text, keyboard)
}

func SendWithMarkup(bot *tgbotapi.BotAPI, chatId int64, text string, keyboard tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = keyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Print(err)
	}
}
