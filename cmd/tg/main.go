package main

import (
	"context"
	"log"
	"os"
	"time"

	storage "github.com/DiMashina05/tg-running-tracker/internal/storage"
	postgres "github.com/DiMashina05/tg-running-tracker/internal/storage/postgres"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	handlers "github.com/DiMashina05/tg-running-tracker/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	//подключаюсь к бд
	dsn := os.Getenv("DATABASE_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Panic(err)

	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Panic(err)

	}
	//подключаюсь к боту
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 300

	updates := bot.GetUpdatesChan(u)

	var store storage.Store = postgres.New(pool)

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
