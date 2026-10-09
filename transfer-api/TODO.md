# TODO

Known gaps and planned improvements for LedgerFlow.

This file tracks code-quality issues, follow-ups, and future work that are
not yet reflected in the code or README. It is intentionally separate from
user-facing documentation.

---

## transfer-api

### Dependency hygiene

- [ ] Run `go mod tidy` and commit the result.
  - `github.com/gorilla/mux` and `go.uber.org/zap` are imported directly
    but are currently marked as `// indirect` in `go.mod`.
  - `go.sum` should be regenerated consistently after tidying.

### Configuration

- [ ] Refactor `BuildConfig` to return `(config, error)` instead of calling
      `logger.Fatalf` directly.
  - Improves testability and lets `main` decide when to exit.
- [ ] Decide whether `PORT=0` should be allowed.
  - Currently `0` passes the range check and binds to an ephemeral port.
  - Either document this behavior or reject `0` as invalid.

### Logging

- [ ] Replace formatted log calls with structured, field-based logging.
  - Example: `l.Infow("starting service", "service", name, "port", port)`.
  - Applies to the startup message and shutdown messages.
- [ ] Use the injected logger in handlers instead of `fmt.Println` when
      JSON encoding fails.
- [ ] Defer `logger.Sync()` in `main` and ignore its error explicitly
      (common practice when syncing on exit).

### Code quality

- [ ] Fix typo: `"Shuting down the server"` → `"Shutting down the server"`.
- [ ] Extract a small helper for writing JSON responses to remove the
      duplicated encoding logic in `HealthHandler` and `ReadinessHandler`.
- [ ] Improve error propagation in `startServer`.
  - The goroutine currently calls `logger.Fatalf`, which exits the process.
  - Consider sending the error on a channel and handling it in `main`.

### Testing

- [ ] Add an integration test that starts the server on an ephemeral port
      and calls `/health` and `/ready` end to end.
  - Optional for the walking skeleton, but required before any real
    infrastructure (SQL Server, Kafka) is introduced.
- [ ] Add a test for `BuildConfig` covering:
  - default values
  - valid `PORT` override
  - invalid `PORT` value
  - out-of-range `PORT` value

### Documentation

- [ ] Confirm whether the project overview should live in the root
      `README.md` and keep `transfer-api/README.md` service-focused.
- [ ] If the root README is added, link to it from `transfer-api/README.md`.
- [ ] Add a short "Known gaps" or "Future work" note only if it affects
      how someone runs or evaluates the service.

---

## Future work (not scoped yet)

- [ ] ledger-service (C#/.NET) walking skeleton.
- [ ] Docker Compose for local development.
- [ ] SQL Server integration for ledger entries.
- [ ] Kafka topics and event publishing.
- [ ] OpenTelemetry instrumentation for both services.
- [ ] Kubernetes manifests and health/readiness probes.
- [ ] GitHub Actions CI pipeline.
