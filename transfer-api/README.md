# transfer-api

Go service that exposes the HTTP API for LedgerFlow transfers and accounts.

## Endpoints

- `GET /health` – liveness probe. Returns `{"Status":"Ok"}`.
- `GET /ready` – readiness probe. Returns `{"Status":"Ready"}`.

## Configuration

The service reads configuration from environment variables:

| Variable       | Default        | Description                |
| -------------- | -------------- | -------------------------- |
| `PORT`         | `8080`         | HTTP port to listen on.    |
| `SERVICE_NAME` | `transfer-api` | Service name used in logs. |

## Running the service

```sh
go run .
```

Override defaults with environment variables:

```sh
PORT=4545 SERVICE_NAME=other_service go run .
```

## Running tests

### Unit tests

```sh
go test .
```

## Development notes

- Dependencies are managed with Go modules. Run `go mod tidy` before committing to keep `go.mod` and `go.sum` clean.
- Use structured logging (zap) with field‑based messages for startup and shutdown.
- Graceful shutdown is implemented using `http.Server.Shutdown` with a 10‑second timeout.
