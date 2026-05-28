.PHONY: dev build test lint up down clean
up:
	docker-compose up -d
down:
	docker-compose down
build:
	cd backend && go build -o bin/api cmd/api/main.go
	cd frontend && npm run build
test:
	cd backend && go test ./...
clean:
	cd backend && rm -rf bin
	cd frontend && rm -rf .next
