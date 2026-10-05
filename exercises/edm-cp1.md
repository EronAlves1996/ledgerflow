## Exercises

Total estimated time: **60 minutes** (30 × 2).

1. **Event-Driven vs. Request-Response Comparison — 10 min**  
   Choose a simple order-placement scenario. List three benefits and three drawbacks of using event-driven microservices instead of request-response microservices. Expected answer should mention asynchronous communication, decoupling, and scalability as benefits, and eventual consistency, harder debugging, and message ordering/idempotency as drawbacks.
   
   A: In a order placement scenario, we have some stock for the products, and for each order place, the stock will go down less one or the quantity the customer posted. USing an event-driven architecture instead of a request-response one have some advantages. The first one is the communication. Each order placement generates an event that can be consumed by other systems in the company. Since it's an event stream, the producer doesn't need to know who are the consumers, just need to put the events into the stream and that's it. The producer doesn't need to handle the how many consumers it have, so it's scalable. But it have some drawbacks. THe first one is that you lose some transactional guarantees. You dont know in the same moment if the request succeeded or not, and should wait for a success event or failure. By the time the success/failure event come back, some time is lost and maybe the product is unavailable. It's harder to debug, because if the events dependencies are too long, it cant form a good trace of the event and maybe its harder to know where the failure come. The payload need to be idempotent, because on event stream, retries can occur on delivery failure, and if the payload is not idempotent, or the logic, it could double the effects.

2. **Identify Bounded Contexts and Domain Events — 10 min**  
   For an online bookstore, identify at least three bounded contexts, such as Catalog, Orders, Payments, and Shipping. For each context, name one domain event it could publish. Expected answer should show clear boundaries and events that belong to the owning context.

   A: Online bookstore are such a very complex system and could be sliced in many ways. 
   First, the consumer should have a clear idea which products it can buy, so a Catalog bounded context could exist. Books can enter the catalog and can exit, so a change in the catalog could emit an event.
   Second, after the consumer choosed his books, it can add to the cart. A Shopping Cart/Orders bounded context could exist and emit events such as item added, item removed, item quantity changed.
   Third, the consumer want's to pay this cart, so, a Payments context could exist, that will manage the total value and communication with the bank provider. Payment Accepted, New Credit Card Added, New Debit Card Added are all examples of domain events that could happen for this one.

3. **Classify Communication Structures — 10 min**  
   Given a customer-onboarding system, classify examples of business communication structures, implementation communication structures, and data communication structures. Expected answer should distinguish team/organizational communication from service APIs, queues, shared databases, and event streams.

   A: Business Communication Structures: Team meetings, face-to-face conversations, slack threads, email, etc.
   Implementation communication structures: Which ports different api's open for communication. Webhooks, rest apis, producing/consuming items for queues, etc.
   Data communication structures: Mostly data schemas. API contract, open api manuals, queue item contract, database schemas, etc.

4. **Apply Conway’s Law — 10 min**  
   A company has separate teams for mobile, backend, data, and operations. Predict the likely service boundaries and explain how Conway’s Law would shape them. Then briefly describe one “inverse Conway maneuver” the company could use. Expected answer should connect team structure to architecture and explain how to reorganize teams to achieve a desired architecture.

   A: doesnt remember conway's law and probably will need to revise it

5. **Make a New Service or Add to Existing Service — 10 min**  
   A payment service already handles card payments. The team wants to add fraud scoring. Decide whether to make a new fraud-scoring service or add it to the existing payment service. Justify the decision using ownership, rate of change, scaling needs, data coupling, and failure isolation. Expected answer should include a clear decision and trade-offs.

   A: Fraud scoring is not only a payment thing, but it could be decisive for credit decisioning: maybe in the future the company want to add ways to give credit for people, and this gonna need to validate the personal identity. Fraud detection changes so much and for some times because frauders change it's method, so it changes a lot. For fraud, we are most likely involved with PII data, more PII data than payments, so the need for data is very different. But the fraud detection thing is very essential and the payment should fail if the fraud detection service fails. So, I go for creating an new service, that will be independent, but a little coupled with the payments service.

6. **Design an Event-Driven Data Flow — 10 min**  
   Design an event flow for “order shipped” notifications. List the events involved, their producers, and their consumers. Explain where eventual consistency appears and how the system could handle a consumer failure. Expected answer should include events such as OrderShipped, NotificationRequested, and NotificationSent, with clear producer/consumer roles.

   A: 
   1. Events 
     1.1 Order Shipped 
     1.2 Client Notified 
     1.3 Merchant Notified 
     1.4 Stock Downgraded 
     1.5 Delivery Finished 
  2. Producers 
     2.1 Order Shipped => Warehouse Management System 
     2.2 Client Notified => Consumer Managament System 
     2.3 Merchant Notified => Merchant Management System 
     2.4 Stock Downgraded => Warehouse Management System 
     2.5 Delivery Finished => Delivery Management System 
  3. Consumers
     3.1 Order Shipped => CMS, MMS, DMS 
     3.2 Client Notified => MMS 
     3.3 Merchant Notified => CMS, WMS 
     3.4 Stock Downgrade => MMS 
     3.5 Delivery FInished => MMS, WMS, CMS
  In this system an eventual consistency can happen in every point, but it's most substantial in the Warehouse Management System, because an order with no stock can be tried to be delivered, or the same event can be delivered twice and the stock downgraded twice as well if we dont have idempotency. Notification is a problematic part, because the client and the merchant can be notified twice too and be misguided in thinking that two orders were shipped. The failures can occur, but we must guarantee that each event is idempotently handled. That way we can investigate what happened and replay these events.

## Feature Suggestion

### Business Justification

Because LedgerFlow is being built from scratch and no code exists yet, the project needs a minimal runnable foundation before ledger logic, Kafka, SQL Server, or observability can be added. A Go `transfer-api` walking skeleton with health and readiness endpoints validates the toolchain, repository layout, local run workflow, and future Kubernetes probe strategy. It also establishes early conventions for configuration, structured logging, graceful shutdown, and HTTP testing without pulling in external infrastructure too soon.

### What Needs to Be Done

- Create the initial repository layout for the Go `transfer-api` service.
- Initialize the Go module and dependency management.
- Implement a minimal HTTP server.
- Add a health endpoint that reports process liveness.
- Add a readiness endpoint that reports whether the service is ready to accept traffic.
- Read the HTTP port and service name from environment variables with sensible defaults.
- Add structured startup and shutdown logging.
- Add graceful shutdown on interrupt signals.
- Add unit tests for the health and readiness handlers.
- Add an integration test that starts the server and calls both endpoints.
- Document how to run tests and start the service locally in the README.
- Do not add SQL Server, Kafka, the C# ledger-service, Docker, or Kubernetes yet.

### Acceptance Criteria

- The service starts locally on the configured port.
- `GET /health` returns HTTP 200 and a clear alive/ok status.
- `GET /ready` returns HTTP 200 and a clear ready status after startup.
- If the configured port is unavailable, the service logs a clear error and exits non-zero.
- The port is configurable via environment variable with a default.
- The service name is configurable via environment variable with a default.
- The service logs a structured startup message including service name and port.
- The service shuts down gracefully on SIGINT or SIGTERM within a few seconds.
- Unit tests cover the health and readiness handlers.
- An integration test starts the server, calls both endpoints, and asserts HTTP 200 and the expected response shape.
- README documents how to run tests and start the service.
- No external infrastructure is required to complete this feature.
- The feature can be completed by hand within 60 minutes.
