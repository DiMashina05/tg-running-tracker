package handlers

import (
	"fmt"

	"github.com/DiMashina05/tg-running-tracker/internal/service"
	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
	tg "github.com/DiMashina05/tg-running-tracker/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, store storage.Store) {
	fromId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	if update.Message.IsCommand() {
		if store.IsWaitingName(fromId) {
			tg.SendText(bot, chatId, "Сначала введи имя")
			return
		}

		if store.IsWaitingDistance(fromId) {
			tg.SendText(bot, chatId, "Сначала введи дистанцию или нажми Назад")
			return
		}

		if update.Message.Command() == "start" {
			if service.CommandStart(store, fromId) {
				tg.SendText(bot, chatId, "Ты уже зарегистрирован")

				tg.SendMenu(bot, chatId)
			} else {
				tg.SendText(bot, chatId, "Данная версия бота пока не рабочая."+
					" Ждём, когда я добавлю бд, допишу больше статистик+разновидностей тренировок и Добавлю заявки в друзья")
				tg.SendText(bot, chatId, "Введите имя")
			}
		}
		return
	}

	text := update.Message.Text

	if store.IsWaitingName(fromId) {

		name, err := service.NameInput(store, text, fromId)

		if err != nil {
			tg.SendText(bot, chatId, err.Error())
			return
		}

		tg.SendText(bot, chatId, fmt.Sprintf("%s, Поздравляю, ты зарегистрировался!", name))

		tg.SendMenu(bot, chatId)

		return
	}

	if store.IsWaitingDistance(fromId) {
		dist, err := service.DistInput(store, text, fromId)
		if err != nil {
			tg.SendText(bot, chatId, err.Error())
			return
		}
		tg.SendText(bot, chatId, fmt.Sprintf("Молодец! Ты пробежал %.2f километров\nТренировка добавлена", dist))
		tg.SendMenu(bot, chatId)
		return
	}

	if !store.IsRegistered(fromId) {
		tg.SendText(bot, chatId, "Сначала зарегистрируйся: введи команду /start")
		return
	}

	tg.SendText(bot, chatId, "Функция пока не добавлена, лучше погладь котика")
}
