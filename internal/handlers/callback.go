package handlers

import (
	"log"

	service "github.com/DiMashina05/tg-running-tracker/internal/service"
	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
	tg "github.com/DiMashina05/tg-running-tracker/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cbMe          = "me"
	cbStats       = "stats"
	cbAddTraining = "add_training"
	cbBack        = "back"
)

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *storage.State) {
	cq := update.CallbackQuery

	cb := tgbotapi.NewCallback(cq.ID, "")
	_, err := bot.Request(cb)
	if err != nil {
		log.Print(err)
	}

	data := cq.Data
	fromId := cq.From.ID
	chatId := cq.Message.Chat.ID
	messageID := cq.Message.MessageID

	if !state.IsRegistered(fromId) {
		tg.SendText(bot, chatId, "Сначала зарегистрируйся: введи команду /start")
		return
	}

	switch data {
	case cbMe:

		text := service.OpenMe(state, fromId)
		tg.EditBack(bot, chatId, messageID, text)

	case cbStats:

		text := service.OpenStats(state, fromId)
		tg.EditBack(bot, chatId, messageID, text)

	case cbAddTraining:

		text := service.OpenAddTraining(state, fromId)
		tg.EditBack(bot, chatId, messageID, text)

	case cbBack:

		service.OpenBack(state, fromId)
		tg.EditMenu(bot, chatId, messageID)

	}
}
