SHELL := bash

.PHONY: up down ps logs init run db run-api init-test db-test logs-test test test-cover test-db

# Поднять Postgres (Docker): dev + test
up:
	docker compose up -d

# Остановить контейнеры
down:
	docker compose down

# Показать статус контейнеров
ps:
	docker compose ps

# Логи dev БД
logs:
	docker compose logs -f db

# Инициализировать dev БД (создать таблицы из SQL-файла)
init:
	. .env && psql "$$DATABASE_URL" -f sql/001_Create.sql

# Запустить бота (подхватить переменные из .env)
run:
	. .env && go run ./cmd/tg

# Быстрый старт dev: поднять БД и применить схему
db: up init

run-api:
	. .env && go run ./cmd/api

# Инициализировать test БД (порт 5434, база running_tracker_test)
init-test:
	psql "postgres://rt_user:rt_pass@127.0.0.1:5434/running_tracker_test" -f sql/001_Create.sql

# Быстрый старт test: поднять БД и применить схему
db-test: up init-test

# Логи test БД
logs-test:
	docker compose logs -f db_test

# Прогнать все тесты
test:
	go test ./... -count=1

# Прогнать все тесты с покрытием
test-cover:
	go test ./... -count=1 -cover

# Поднять test-БД, применить схему и прогнать тесты
test-db: db-test
	go test ./... -count=1
