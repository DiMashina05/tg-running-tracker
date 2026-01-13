# tg-running-tracker

Telegram-бот для учёта беговых тренировок и просмотра статистики.
Проект написан на Go и используется как pet-проект для практики backend-разработки.

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

## Запуск проекта

### Требования
- Go 1.24+
- Docker
- Docker Compose
- Telegram Bot Token

### Переменные окружения
Проект использует файл `.env` со следующими переменными:

- BOT_TOKEN — токен Telegram-бота
- DATABASE_URL — строка подключения к PostgreSQL

Пример:

BOT_TOKEN=your_telegram_bot_token
DATABASE_URL=postgres://rt_user:rt_pass@127.0.0.1:5432/running_tracker?sslmode=disable

### Быстрый старт
1) Поднять PostgreSQL в Docker
2) Применить SQL-схему
3) Запустить бота

Командой:

make db
make run

### Ручной запуск (без Makefile)
Поднять базу данных:

docker compose up -d

Применить SQL-скрипт:

psql "$DATABASE_URL" -f sql/001_Create.sql

Запустить бота:

go run ./cmd/tg

## Структура проекта
cmd/tg/              — точка входа приложения
internal/handlers/   — обработчики сообщений и callback-кнопок
internal/service/    — бизнес-логика
internal/storage/    — работа с хранилищем данных
internal/telegram/   — обёртки над Telegram Bot API
sql/                 — SQL-скрипты (создание таблиц)
docker-compose.yml   — конфигурация PostgreSQL
Makefile             — команды для локальной разработки

## Статус проекта
В активной разработке.

Планируется:
- покрытие тестами
- расширение бизнес-логики
- CI (GitHub Actions)