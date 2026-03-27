include .env
export

.PHONY: db-up db-down migrate-diff migrate-apply sqlc lint test run

# Docker
db-up:
	docker compose up -d

db-down:
	docker compose down

# Atlas
migrate-diff:
	atlas migrate diff $(name) \
		--dir "file://internal/infrastructure/postgres/migration" \
		--to "file://internal/infrastructure/postgres/schema/schema.hcl" \
		--dev-url "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

migrate-apply:
	atlas migrate apply \
		--dir "file://internal/infrastructure/postgres/migration" \
		--url "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

# sqlc
sqlc:
	sqlc generate

# Go
run:
	go run cmd/playtics/main.go

lint:
	golangci-lint run ./...

test:
	go test ./...