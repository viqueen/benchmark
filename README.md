# benchmark

## requirements

- Docker
- Docker Compose
- Make

## setup

### observability

```bash
make monitoring
```

- open grafana [http://localhost:3000](http://localhost:3000)
- open prometheus [http://localhost:9090](http://localhost:9090)

### codegen

```bash
make scheme-codegen
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
