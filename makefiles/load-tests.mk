.PHONY: start-monitoring
start-monitoring:
	@echo "Starting monitoring services..."
	@docker compose -f ./load-tests/docker-compose.yaml --profile monitoring up -d

.PHONY: stop-monitoring
stop-monitoring:
	@echo "Stopping monitoring services..."
	@docker compose -f ./load-tests/docker-compose.yaml --profile monitoring down


.PHONY: load-tests
load-tests:
	@echo "Running load tests..."
	@NAME=$${NAME:-go-basic} TARGET="--vus 10 --duration 5m" docker compose -f ./load-tests/docker-compose.yaml --profile test up

.PHONY: smoke-tests
smoke-tests:
	@echo "Running smoke tests..."
	@NAME=$${NAME:-go-basic} TARGET="--vus 1 --duration 30s" docker compose -f ./load-tests/docker-compose.yaml --profile test up

.PHONY: stress-tests
stress-tests:
	@echo "Running stress tests..."
	@NAME=$${NAME:-go-basic} TARGET="--vus 50 --duration 10m" docker compose -f ./load-tests/docker-compose.yaml --profile test up