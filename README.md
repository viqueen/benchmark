# benchmark

## requirements

- Docker
- Docker Compose
- Make

## setup

### environment

```bash
make start-monitoring
```

- open grafana [http://localhost:3000](http://localhost:3000)
- open prometheus [http://localhost:9090](http://localhost:9090)

```bash
make start-stack
```

- comes with postgres

### codegen

- build protobuf-gen locally

```bash
make build-protobuf-gen-image
```

- generate code

```bash
make local-schema-codegen
```

## flow

```mermaid
sequenceDiagram
    participant client as Client
    participant server as Server
    participant database as Database

    client ->> server: /ListFolders
    server ->> database: Folders Query
    database -->> server: {folders}
    server -->> client: {folders}
```
