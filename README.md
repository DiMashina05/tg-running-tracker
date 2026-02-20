# tg-running-tracker

- Telegram-бот для учёта беговых тренировок и просмотра статистики.
- Бот в телеграме: @DM_Running_bot

## Возможности
- REST API для работы с пользователями и статистикой
- Регистрация пользователя
- Добавление беговых тренировок (дистанция)
- Просмотр статистики
- Интерактивное управление через Telegram-кнопки

## Технологии
- Go
- REST API
- Telegram Bot API
- PostgreSQL
- Docker / Docker Compose
- CI (GitHub Actions)
- pgx (PostgreSQL driver)
- go test (юнит/интеграционные тесты)

## Запуск проекта

### Требования
- Go 1.24+
- Docker + Docker Compose
- Telegram Bot Token

### Переменные окружения
Для запуска используются переменные окружения из файла `.env`:

- `BOT_TOKEN` — токен Telegram-бота
- `DATABASE_URL` — строка подключения к dev PostgreSQL

Пример:

BOT_TOKEN=your_telegram_bot_token
DATABASE_URL=postgres://rt_user:rt_pass@127.0.0.1:5432/running_tracker?sslmode=disable
SERVER_PORT=:8080

### Быстрый старт (dev)
Поднять dev-БД, применить SQL-схему и запустить бота:

make db
make run

### Запуск REST API

make run-api

## Тесты

### Быстрый старт (test DB + тесты)
Поднять отдельную test-БД, применить схему и прогнать тесты:

make test-db

### Запуск тестов без БД
Прогнать все тесты:

make test

Прогнать все тесты с покрытием:

make test-cover

## Структура проекта
cmd/tg/              — точка входа Telegram-бота
cmd/api/             — точка входа REST API

internal/handlers/   — обработчики сообщений и callback-кнопок (Telegram)
internal/httpapi/    — REST API (handlers, DTO, server)
internal/service/    — бизнес-логика
internal/storage/    — интерфейсы хранилища данных
internal/storage/postgres/ — реализация хранилища на PostgreSQL
internal/telegram/   — обёртки над Telegram Bot API

sql/                 — SQL-скрипты (создание таблиц)
.github/workflows/   — CI (GitHub Actions)


## Статус проекта
В активной разработке.

Планируется:
- OpenAPI/Swagger
- Prometheus/Grafana