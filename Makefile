EASYJSON_PATHS = ./internal/...

# ===== RUN =====
.PHONY: build
build:
	docker compose -f docker-compose.yml build api

.PHONY: up
up:
	docker compose -f docker-compose.yml up -d --build db sessions-db api dev-frontend

.PHONY: deploy
deploy:
	git pull
	make stop
	make up

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

# ===== GENERATORS =====

.PHONY: swag
swag:
	swag init -g cmd/api/main.go

.PHONY: mocks
mocks:
	./scripts/gen_mocks.sh

.PHONY: easyjson
easyjson:
	go generate ${EASYJSON_PATHS}

# ===== TESTS =====

.PHONY: unit-test
unit-test:
	go test ./tests/unit/...