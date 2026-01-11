package main

import (
	"log"
	"os"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"

	handlers "github.com/DiMashina05/tg-running-tracker/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 300

	updates := bot.GetUpdatesChan(u)

	var store storage.Store = storage.NewState()

	for update := range updates {

		if update.CallbackQuery != nil {
			handlers.HandleCallback(bot, update, store)
			continue
		}

		if update.Message != nil {
			handlers.HandleMessage(bot, update, store)
		}
	}
}
