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
- [ ] Replace `zap.NewProduction()` in `main_test.go` with `zap.NewNop()` or a `zaptest` logger to avoid noisy tests.
- [ ] [ ] Add an integration test in `main_test.go` that hits `/health` and `/ready` through `routes(l)` with a `zaptest/observer` logger attached, asserting both the response (status + body) and the `request completed` log fields (`method`, `path`, `statusCode`, `duration`). This closes AC4 for the audit feature.
- [ ] Change audit duration calculation from `time.Now().UnixMilli() - start.UnixMilli()` to `time.Since(start)` and log as `duration_ms` or `zap.Duration`.
- [ ] Consider wrapping the router globally instead of wrapping `/health` and `/ready` individually, so future routes are audited automatically.

### Code quality

- [ ] Fix typo: `"Shuting down the server"` → `"Shutting down the server"`.
- [ ] Extract a small helper for writing JSON responses to remove the
      duplicated encoding logic in `HealthHandler` and `ReadinessHandler`.
- [ ] Improve error propagation in `startServer`.
  - The goroutine currently calls `logger.Fatalf`, which exits the process.
  - Consider sending the error on a channel and handling it in `main`.
- [ ] Ensure `inspectingWriter` supports optional `http.Flusher`/`http.Hijacker` if the audit middleware is used on streaming endpoints.
- [ ] Default or fallback the audit status code to `http.StatusOK` when a handler does not call `WriteHeader`.

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
- [ ] Add an audit middleware test that verifies log output, not only response behavior.

### Documentation

- [ ] Confirm whether the project overview should live in the root
      `README.md` and keep `transfer-api/README.md` service-focused.
- [ ] If the root README is added, link to it from `transfer-api/README.md`.
- [ ] Add a short "Known gaps" or "Future work" note only if it affects
      how someone runs or evaluates the service.
- [ ] Keep `transfer-api/README.md` aligned with the audit log fields and per-route wrapping behavior.

## exercises

- [ ] Remove accidental `exercises/ledgerflow-tdd/integers/diff`; it appears to be a committed diff artifact.
- [ ] Fix Go doc typo in `integers.Add`: `it's sum` → `its sum`.
- [ ] Clean up `exercises/lgwt-1.md` code fences and duplicated diff blocks.
- [ ] Align exercise 3's test data with the actual `integers_test.go` cases, or vice versa. The spec in `lgwt-1.md` says `2+2=4`, `-1+1=0`, `0+0=0`, `-3+-2=-5`; the code tests `2+2=4`, `3+(-4)=-1`, `5+(-5)=0`, `0+0=0`. Same categories, different numbers — pick one source of truth.
- [ ] Note the closure capture pattern in `greeting_test.go`. `for _, v := range table { t.Run(..., func(t *testing.T){ ... v ... }) }` is safe on Go 1.22+, but if the exercise is meant to teach closures, call this out explicitly (or use `v := v`).
- [ ] Decide where the feature suggestion belongs. It is currently duplicated inside `exercises/lgwt-1.md`; that file should be the exercise log only.

---

## Future work (not scoped yet)

- [ ] ledger-service (C#/.NET) walking skeleton.
- [ ] Docker Compose for local development.
- [ ] SQL Server integration for ledger entries.
- [ ] Kafka topics and event publishing.
- [ ] OpenTelemetry instrumentation for both services.
- [ ] Kubernetes manifests and health/readiness probes.
- [ ] GitHub Actions CI pipeline.
