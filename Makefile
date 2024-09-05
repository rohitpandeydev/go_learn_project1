.PHONY: all

build:
	@go build -o bin/ecom cmd/main.go


test:
	@go test -v ./...

run: build
	@./bin/ecom


migration:
	@echo "Creating migration: $(filter-out $@,$(MAKECMDGOALS))"
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

migrate-force:
	@go run cmd/migrate/main.go force $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
