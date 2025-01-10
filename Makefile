.PHONY: monitoring
monitoring:
	@echo "Starting monitoring..."
	@docker-compose -f _harness/monitoring/docker-compose.yaml up -d


.PHONY: build-protobuf-gen-image
build-protobuf-gen-image:
	@echo "Building protobuf-gen image..."
	@docker build -t protobuf-gen _tools/docker-images/protobuf-gen

.PHONY: local-schema-codgen
local-schema-codegen:
	@echo "Generating code from schema..."
	@docker run --rm \
		--volume $(PWD)/_schema:/workspace/schema \
		--volume $(PWD)/api:/workspace/api \
		--volume $(PWD)/product:/workspace/product \
		--workdir /workspace/schema \
		protobuf-gen generate --verbose

.PHONY: schema-codegen
schema-codegen:
	@echo "Generating code from schema..."
	@docker run --rm \
		--volume $(PWD)/_schema:/workspace/schema \
		--volume $(PWD)/api:/workspace/api \
		--volume $(PWD)/product:/workspace/product \
		--workdir /workspace/schema \
		ghcr.io/viqueen/benchmark-protobuf-gen:main generate --verbose

