.PHONY: monitoring
start-monitoring:
	@echo "Starting monitoring..."
	@docker-compose -f _harness/monitoring/docker-compose.yaml up -d

.PHONY: stop-monitoring
stop-monitoring:
	@echo "Stopping monitoring..."
	@docker-compose -f _harness/monitoring/docker-compose.yaml down

.PHONY: stack
start-stack:
	@echo "Starting stack..."
	@docker-compose -f _harness/stack/docker-compose.yaml up -d

.PHONY: stop-stack
stop-stack:
	@echo "Stopping stack..."
	@docker-compose -f _harness/stack/docker-compose.yaml down

.PHONY: build-protobuf-gen-image
build-protobuf-gen-image:
	@echo "Building protobuf-gen image..."
	@docker build -t protobuf-gen _tools/docker-images/protobuf-gen

.PHONY: local-schema-codgen
local-go-codegen:
	@echo "Generating code from schema..."
	@docker run --rm \
		--volume $(PWD)/_schema:/workspace/schema \
		--volume $(PWD)/api:/workspace/api \
		--volume $(PWD)/product:/workspace/product \
		--volume $(PWD)/_schema/buf.gen.go.yaml:/workspace/schema/buf.gen.yaml \
		--workdir /workspace/schema \
		protobuf-gen  generate  --verbose

	@echo "Running SQLC..."
	@docker run --rm \
		--volume $(PWD)/product/backend-connect-go:/src \
		--workdir /src \
		sqlc/sqlc:1.27.0 \
		generate

	@echo "Running Flyway migrations..."
	@docker run --rm \
		--network="benchmark_stack" \
		--volume $(PWD)/product/backend-connect-go/data/schema:/flyway/backend/data/schema \
		flyway/flyway:7.2.0-alpine \
			-url="jdbc:postgresql://database:5432/benchmark" \
			-user="benchmark" -password="benchmark" \
            -schemas="public" -defaultSchema="public" \
            -locations=filesystem:/flyway/backend/data/schema \
            migrate info

.PHONY: go-codegen
go-codegen:
	@echo "Generating code from schema..."
	@docker run --rm \
		--volume $(PWD)/_schema:/workspace/schema \
		--volume $(PWD)/api:/workspace/api \
		--volume $(PWD)/product:/workspace/product \
		--volume $(PWD)/_schema/buf.gen.go.yaml:/workspace/schema/buf.gen.yaml \
		--workdir /workspace/schema \
		ghcr.io/viqueen/benchmark-protobuf-gen:main generate --config buf.gen.go.yaml --verbose

