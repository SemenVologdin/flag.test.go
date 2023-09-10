MIGRATE=docker-compose exec postgres migrate -path migrations -database "postgres://$(DB_HOST)/$(DB_NAME)?sslmode=disable"

migrate-status:
	$(MIGRATE) status
migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

up:
	docker-compose up -d

build:
	docker-compose build

buildup:
	docker compose up --build

stop:
	docker-compose down