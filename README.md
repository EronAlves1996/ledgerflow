# LedgerFlow

> A hands-on, event-driven mini-fintech ledger built from scratch to practice Go, C#, SQL Server, Kafka, Kubernetes, AWS, and OpenTelemetry.

## Overview

LedgerFlow is a learning project that simulates a small financial ledger with double-entry accounting, idempotent transfers, auditability, and distributed tracing. It is designed to be built incrementally, starting with a minimal core and evolving into a polyglot microservices system.

The goal is not to build a production-ready banking platform, but to create a realistic playground for fintech engineering concepts: consistency, concurrency, idempotency, event-driven architecture, observability, and cloud deployment.

## Goals

- Practice Go and C# in a realistic fintech domain.
- Implement double-entry accounting correctly.
- Handle concurrency, deadlocks, and isolation levels in SQL Server.
- Use Kafka for event-driven communication and audit trails.
- Instrument services with OpenTelemetry and visualize traces in Grafana, Dynatrace, and Elastic.
- Deploy to Kubernetes locally and on AWS.
- Learn by building, testing, and observing.

## Non-Goals

- Real money movement.
- Full regulatory compliance.
- Production-grade security or scalability.
- Multi-region or high-availability setup in early phases.

## Planned Features

### Core Ledger

- [ ] Create accounts (checking, savings, internal).
- [ ] Deposit funds into an account.
- [ ] Withdraw funds from an account.
- [ ] Transfer funds between accounts.
- [ ] Double-entry accounting: every transaction has balanced debits and credits.
- [ ] Account balance calculation from ledger entries.
- [ ] Transaction history and statements.
- [ ] Transaction reversal with audit trail.
- [ ] Immutable ledger entries.

### API and Correctness

- [ ] REST API for account and transaction operations.
- [ ] Idempotency support via `Idempotency-Key` header.
- [ ] Request validation and consistent error responses.
- [ ] Pagination and filtering for transaction history.
- [ ] Optimistic and pessimistic concurrency control examples.
- [ ] Row versioning and snapshot isolation experiments.
- [ ] Deadlock detection and retry policies.
- [ ] Idempotent transfer endpoint.

### Event-Driven Architecture

- [ ] Kafka topics for transaction events.
- [ ] Publish `TransactionCreated`, `TransactionCompleted`, `TransactionReversed`.
- [ ] Outbox pattern for reliable event publishing.
- [ ] Consumer for notifications or audit.
- [ ] Dead-letter queue (DLQ) and retry handling.
- [ ] Schema evolution strategy for events.

### Observability

- [ ] OpenTelemetry instrumentation for Go services.
- [ ] OpenTelemetry instrumentation for C# services.
- [ ] Distributed tracing across API, ledger, SQL Server, and Kafka.
- [ ] Custom metrics for transfers, latency, errors, and deadlocks.
- [ ] Structured logging with correlation IDs.
- [ ] Export traces to Grafana, Dynatrace, and Elastic.
- [ ] Dashboards for key fintech flows.
- [ ] Alerts for failed transfers and DLQ growth.

### Security

- [ ] OAuth2 / JWT authentication.
- [ ] Role-based authorization.
- [ ] Rate limiting.
- [ ] PII masking in logs and traces.
- [ ] Secrets management with AWS Secrets Manager or local Vault.
- [ ] Basic audit logging for sensitive operations.

### Infrastructure and Deployment

- [ ] Local development with Docker Compose.
- [ ] Kubernetes manifests with kind or minikube.
- [ ] Helm charts or Kustomize overlays.
- [ ] Health checks, readiness probes, and liveness probes.
- [ ] Horizontal Pod Autoscaler (HPA) example.
- [ ] AWS deployment: EKS, RDS SQL Server, MSK, S3, Secrets Manager.
- [ ] CI/CD pipeline with GitHub Actions.
- [ ] Infrastructure as Code with Terraform or AWS CDK.

### Testing

- [ ] Unit tests for domain logic.
- [ ] Integration tests with SQL Server and Kafka.
- [ ] Contract tests between Go and C# services.
- [ ] Load tests for transfer concurrency.
- [ ] Chaos experiments for Kafka and SQL Server failures.
- [ ] End-to-end test for a full transfer flow.

## Architecture (Planned)

```text
+----------------+       +-------------------+       +----------------+
|  Client / CLI  | ----> |  transfer-api     | ----> |  ledger-service|
|                |       |  (Go)             |       |  (C#/.NET)     |
+----------------+       +-------------------+       +----------------+
                                |                            |
                                |                            |
                                v                            v
                         +-------------+              +-------------+
                         |   Kafka     |              | SQL Server  |
                         |  (events)   |              |  (ledger)   |
                         +-------------+              +-------------+
                                |
                                v
                         +-------------+
                         | Observability|
                         | OTel Collector|
                         | Grafana      |
                         | Dynatrace    |
                         | Elastic      |
                         +-------------+
```

## Services

- [`transfer-api`](./transfer-api/README.md) – Go service that exposes the HTTP API for transfers and accounts.

Each service has its own README with build, test, and run instructions.

---

## 2. Study Recommendations

To avoid repeating the flaws found in the review, focus on these topics:

### Conway’s Law & Team Topologies

- **Book:** _Team Topologies_ by Matthew Skelton & Manuel Pais – explains Conway’s Law and the inverse Conway maneuver.
- **Article:** Martin Fowler’s “Conway’s Law” – concise explanation with examples.
- **Video:** “Inverse Conway Maneuver” talks on YouTube (e.g., from DevOps conferences).

### Go Modules & Dependency Management

- **Official Go Blog:** “Using Go Modules” – covers `go mod tidy`, direct vs indirect dependencies.
- **Documentation:** `go help mod tidy` and `go help modules`.

### Testing in Go

- **Package `net/http/httptest`** – for handler tests and integration tests with `httptest.NewServer`.
- **Article:** “Testing HTTP handlers in Go” by various authors – shows table-driven tests and integration patterns.
- **Book:** _The Go Programming Language_ (Donovan & Kernighan) – chapter on testing.

### Configuration & Error Handling

- **Article:** “Configuration in Go” – prefer returning errors from config builders instead of calling `log.Fatal` inside libraries.
- **Pattern:** Use `func BuildConfig() (config, error)` and handle the error in `main`.

### Structured Logging

- **Zap documentation:** “Getting Started” and “Logger” – learn about `zap.SugaredLogger` vs `zap.Logger`, and field‑based logging (`Infow`, `Errorw`).
- **Best practices:** Always log with key‑value pairs, not formatted strings.

### Graceful Shutdown

- **Go Blog:** “Graceful Shutdown in Go” – using `context.WithTimeout` and `server.Shutdown`.

### Event-Driven Microservices

- **Book:** _Building Event-Driven Microservices_ by Adam Bellemare – covers events, idempotency, outbox pattern, DLQ.
- **Book:** _Microservices Patterns_ by Chris Richardson – chapters on event‑driven architecture, sagas, and idempotent consumers.

### Domain-Driven Design

- **Book:** _Domain-Driven Design_ by Eric Evans – bounded contexts, ubiquitous language.
- **Book:** _Implementing Domain-Driven Design_ by Vaughn Vernon – practical DDD with events.

### Fintech & Accounting

- **Article:** “Double-Entry Accounting for Developers” – explains debits, credits, and immutable ledgers.
- **Pattern:** Idempotency keys and exactly‑once processing in payment systems.

---

## 3. Exercise Revisions

Based on the review, these sections need revision:

### Exercise 1 – Event-Driven vs. Request-Response

- **Correction:** The payload does **not** need to be idempotent; the **consumer’s handling logic** must be idempotent. Events are immutable facts.
- **Add:** Mention that idempotency is achieved by deduplication keys or state checks on the consumer side.

### Exercise 2 – Identify Bounded Contexts and Domain Events

- **Improvement:** Use past‑tense event names consistently:
  - `BookAdded`, `BookRemoved`, `BookPriceChanged`
  - `CartItemAdded`, `CartItemRemoved`, `CartItemQuantityChanged`
  - `PaymentAccepted`, `CreditCardAdded`, `DebitCardAdded`

### Exercise 4 – Apply Conway’s Law

- **Rewrite completely.** You stated you didn’t remember Conway’s Law.
- **Required content:**
  - Define Conway’s Law: “Organizations design systems that mirror their own communication structure.”
  - Predict boundaries: mobile team → mobile/BFF layer; backend team → backend services/APIs; data team → analytics pipelines; operations team → infrastructure tooling.
  - Explain how Conway’s Law shapes these boundaries (teams build what they can communicate about).
  - Describe an inverse Conway maneuver: reorganize into stream‑aligned teams around business domains (e.g., payments, accounts, notifications) plus a platform team for shared infrastructure, so the architecture follows the desired service boundaries.

### Exercise 5 – Make a New Service or Add to Existing Service

- **Add:** Discuss fail‑open vs. fail‑closed behavior.
  - If fraud scoring fails, should payment fail (fail‑closed) or proceed with risk (fail‑open)?
- **Add:** Circuit breakers, timeouts, and latency budgets for the fraud service call.
- **Add:** Explicit mention of data coupling (PII vs. payment data) and regulatory boundaries.

### Exercise 6 – Design an Event-Driven Data Flow

- **Improvement:** Use consistent past‑tense event names:
  - `OrderShipped`, `ClientNotified`, `MerchantNotified`, `StockDecremented`, `DeliveryCompleted`.
- **Add:** DLQ and retry handling for failed consumers.
- **Add:** Outbox pattern for reliable event publishing from the producer side.
- **Clarify:** Eventual consistency is inherent; ensure idempotent consumers and replay capability.

After revising these points, your answers will be complete and aligned with the expected solutions.

```

```
