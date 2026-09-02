.PHONY: build-api run build-migrate migrate-up migrate-down clean

build-api:
	@go build -o bin/api ./cmd/api

run: build-api
	@./bin/api

build-migrate:
	@go build -o bin/migrate ./cmd/migrate

migrate-up: build-migrate
	@./bin/migrate up

migrate-down: build-migrate
	@./bin/migrate down

clean:
	@rm -rf bin/