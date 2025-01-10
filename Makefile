.PHONY: monitoring
monitoring:
	@echo "Starting monitoring..."
	@docker-compose -f harness/monitoring/docker-compose.yaml up -d