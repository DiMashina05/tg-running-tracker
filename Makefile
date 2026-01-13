.PHONY: up down ps logs init run db

# Поднять Postgres (Docker)
up:
	docker compose up -d

# Остановить контейнеры 
down:
	docker compose down

# Показать статус контейнеров
ps:
	docker compose ps

# Логи Postgres
logs:
	docker compose logs -f db

# Инициализировать БД (создать таблицы из SQL-файла)
init:
	. .env && psql "$$DATABASE_URL" -f sql/001_Create.sql

# Запустить бота (подхватить переменные из .env)
run:
	. .env && go run ./cmd/tg

# Быстрый старт: поднять БД и применить схему
db: up init

