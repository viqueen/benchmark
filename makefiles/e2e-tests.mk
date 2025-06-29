.PHONY: e2e-tests
e2e-tests:
	@echo "Running E2E tests..."
	@NAME=$${NAME:-go-basic} docker compose -f ./e2e-tests/docker-compose.yaml up
