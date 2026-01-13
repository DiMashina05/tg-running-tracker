package handlers

import (
	"fmt"

	"github.com/DiMashina05/tg-running-tracker/internal/service"
	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
	tg "github.com/DiMashina05/tg-running-tracker/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, store storage.Store) {
	fromID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	if update.Message.IsCommand() {
		if store.IsWaitingName(fromID) {
			tg.SendText(bot, chatID, "Сначала введи имя")
			return
		}

		if store.IsWaitingDistance(fromID) {
			tg.SendText(bot, chatID, "Сначала введи дистанцию или нажми Назад")
			return
		}

		if update.Message.Command() == "start" {
			if service.CommandStart(store, fromID) {
				tg.SendText(bot, chatID, "Ты уже зарегистрирован")

				tg.SendMenu(bot, chatID)
			} else {
				tg.SendText(bot, chatID, "Данная версия бота пока не рабочая."+
					" Ждём, когда я добавлю бд, допишу больше статистик+разновидностей тренировок и Добавлю заявки в друзья")
				tg.SendText(bot, chatID, "Введите имя")
			}
		}
		return
	}

	text := update.Message.Text

	if store.IsWaitingName(fromID) {

		name, err := service.NameInput(store, text, fromID)

		if err != nil {
			tg.SendText(bot, chatID, err.Error())
			return
		}

		tg.SendText(bot, chatID, fmt.Sprintf("%s, Поздравляю, ты зарегистрировался!", name))

		tg.SendMenu(bot, chatID)

		return
	}

	if store.IsWaitingDistance(fromID) {
		dist, err := service.DistInput(store, text, fromID)
		if err != nil {
			tg.SendText(bot, chatID, err.Error())
			return
		}
		tg.SendText(bot, chatID, fmt.Sprintf("Молодец! Ты пробежал %.2f километров\nТренировка добавлена", dist))
		tg.SendMenu(bot, chatID)
		return
	}

	if !store.IsRegistered(fromID) {
		tg.SendText(bot, chatID, "Сначала зарегистрируйся: введи команду /start")
		return
	}

	tg.SendText(bot, chatID, "Функция пока не добавлена, лучше погладь котика")
}
