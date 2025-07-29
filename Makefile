test:
	docker compose exec api go test ./internal/config -v
lint:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v2.3.0 golangci-lint run
run: 
	docker compose up --build -d
stop:
	docker compose down -rmi all -v

ci: run lint test
