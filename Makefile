.PHONY: up down build logs ps clean

up:
	docker compose up -d

build:
	docker compose up --build -d

down:
	docker compose down

clean:
	docker compose down -v

logs:
	docker compose logs -f

ps:
	docker compose ps

rebuild-%:
	docker compose build --no-cache $*
	docker compose up -d $*
