.PHONY: monitoring
monitoring:
	@echo "Starting monitoring..."
	@docker-compose -f _harness/monitoring/docker-compose.yaml up -d


.PHONY: schema-codegen
schema-codegen:
	@echo "Generating code from schema..."
	@docker run --rm \
		--volume $(PWD)/_schema:/workspace/schema \
		--volume $(PWD)/api:/workspace/api \
		--workdir /workspace/schema \
		bufbuild/buf:1.49.0 generate --verbose