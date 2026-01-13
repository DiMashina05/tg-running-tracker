package handlers

import (
	"log"

	service "github.com/DiMashina05/tg-running-tracker/internal/service"
	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
	tg "github.com/DiMashina05/tg-running-tracker/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Callback string

const (
	cbMe          Callback = "me"
	cbStats       Callback = "stats"
	cbAddTraining Callback = "add_training"
	cbBack        Callback = "back"
)

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, store storage.Store) {
	cq := update.CallbackQuery

	cb := tgbotapi.NewCallback(cq.ID, "")
	_, err := bot.Request(cb)
	if err != nil {
		log.Print(err)
	}

	data := cq.Data
	fromID := cq.From.ID
	chatID := cq.Message.Chat.ID
	messageID := cq.Message.MessageID

	if !store.IsRegistered(fromID) {
		tg.SendText(bot, chatID, "Сначала зарегистрируйся: введи команду /start")
		return
	}

	switch Callback(data) {
	case cbMe:

		text := service.OpenMe(store, fromID)
		tg.EditBack(bot, chatID, messageID, text)

	case cbStats:

		text := service.OpenStats(store, fromID)
		tg.EditBack(bot, chatID, messageID, text)

	case cbAddTraining:

		text := service.OpenAddTraining(store, fromID)
		tg.EditBack(bot, chatID, messageID, text)

	case cbBack:

		service.OpenBack(store, fromID)
		tg.EditMenu(bot, chatID, messageID)

	}
}
