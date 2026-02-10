# tg-running-tracker

Telegram-бот для учёта беговых тренировок и просмотра статистики.
Бот в телеграме: @DM_Running_bot

## Возможности
- Регистрация пользователя
- Добавление беговых тренировок (дистанция)
- Просмотр статистики
- Интерактивное управление через Telegram-кнопки

## Технологии
- Go
- Telegram Bot API
- PostgreSQL
- Docker / Docker Compose
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

### Быстрый старт (dev)
Поднять dev-БД, применить SQL-схему и запустить бота:

make db
make run

## Тесты

### Быстрый старт (test DB + тесты)
Поднять отдельную test-БД, применить схему и прогнать тесты:

make test-db

### Запуск тестов без БД
Прогнать все тесты:

make test

Прогнать все тесты с покрытием:

make test-cover

## Полезные команды (Makefile)
- `make up` / `make down` — поднять/остановить dev+test БД (docker compose)
- `make logs` / `make logs-test` — логи dev/test БД
- `make init` — применить схему к dev БД
- `make db` — dev: up + init
- `make init-test` — применить схему к test БД (порт 5434, база `running_tracker_test`)
- `make db-test` — test: up + init-test

## Структура проекта
cmd/tg/              — точка входа приложения
internal/handlers/   — обработчики сообщений и callback-кнопок
internal/service/    — бизнес-логика
internal/storage/    — работа с хранилищем данных
internal/telegram/   — обёртки над Telegram Bot API
sql/                 — SQL-скрипты (создание таблиц)

## Статус проекта
В активной разработке.

Планируется:
- рефакторинг бизнес-логики
- добавление REST API
- CI (GitHub Actions)