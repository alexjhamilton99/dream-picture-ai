run: build
	@./bin/dream-picture-ai

install:
	# @go install github.com/alexjhamilton99/dream-picture-ai
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

css:
	@tailwindcss -i view/css/app.css -o public/styles.css --watch

templ:
	@templ generate --watch --proxy=http://localhost:3000

build:
	@npx tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@go build -tags dev -o bin/dream-picture-ai main.go

up: ## database migration up
	@go run cmd/migrate/main.go up

reset:
	@go run cmd/reset/main.go up

down: ## database migration down
	@go run cmd/migrate/main.go down

migration: ## migrations against the database
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go
