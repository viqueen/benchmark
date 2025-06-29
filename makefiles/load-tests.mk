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
	@NAME=$${NAME:-go-basic} docker compose -f ./load-tests/docker-compose.yaml --profile test up