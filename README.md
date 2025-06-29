# benchmark

## requirements

- Docker
- Make

## getting started

### codegen

The following command generates the SDKs for the API 
endpoints defined in the ProtoBuf files available in `_schema/protos/`.

It creates an SDK for each technology stack under the `api/*-sdk` directories.

```bash
make api-codegen
```

### e2e tests

The following command runs all end-to-end tests defined in the `e2e-tests` directory against a given implementation that is available in the `packages/` directory.

It expects the API server to be running on port `8080` by default.

#### golang

```bash
make e2e-tests NAME=go-basic
```

#### node

```bash
make e2e-tests NAME=node-basic
```

#### java

```bash
make e2e-tests NAME=java-basic
```

### load tests

The following command runs all load tests defined in the `load-tests` directory against a given implementation that is available in the `packages/` directory.

It expects the API server to be running on port `8080` by default.

#### monitoring stack

```bash
make start-monitoring
```

#### golang

```bash
make load-tests NAME=go-basic
```