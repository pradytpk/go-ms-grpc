.PHONY: graph
graph:
	@go run github.com/99designs/gqlgen generate

.PHONY: up
up:
	@docker-compose up -d --build