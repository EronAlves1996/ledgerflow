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
