# ===== RUN =====
.PHONY: build
build:
	docker compose -f docker-compose.yml build api

.PHONY: up
up:
	docker compose -f docker-compose.yml up -d --build db sessions-db api

.PHONY: stop
stop:
	docker compose -f docker-compose.yml stop

.PHONY: down
down:
	docker compose -f docker-compose.yml down -v

# ===== LOGS =====
service = api
.PHONY: logs
logs:
	docker compose logs -f $(service)
