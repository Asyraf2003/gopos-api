# ADR 0013: ProductCatalog Versioning And Transactional Audit Outbox

## Status

accepted

## Context

ProductCatalog is being migrated from the Laravel POS system into the Go Echo API with PostgreSQL and hexagonal architecture.

ProductCatalog write operations must support:

- create product;
- update product;
- soft delete product;
- restore product.

These operations need durable business history and audit evidence without slowing down the API request hot path.

The API mission is to keep normal request handling bounded and compatible with a sub-1-second target. Full audit enrichment, telemetry export, reporting projection, dashboard updates, notification dispatch, or external logging must not block ProductCatalog write responses.

ProductCatalog also needs product version history so the system can explain product changes over time.

Audit must answer:

- who performed the action;
- which role or authority performed the action;
- when the action happened;
- what action happened;
- what the relevant data looked like before the action;
- what the relevant data looked like after the action;
- why the action happened when a reason is supplied;
- request or correlation metadata when available.

## Decision

ProductCatalog write operations must use transactional product versioning and a transactional audit intent/outbox.

For create, update, soft delete, and restore, the usecase transaction must write:

```text
products row change
product_versions row
audit_outbox or audit_intents minimal row
```

These writes must be committed together.

The request hot path must not wait for full audit materialization, full telemetry export, reporting projection, dashboard projection, notification dispatch, external logging, or slow audit enrichment.

Full audit materialization may be handled later by a worker/cold path that reads audit outbox/intent records.

## Hot Path Contract

Required hot-path flow:

```text
validate request
load required state
begin transaction
write products
write product_versions
write audit_outbox or audit_intents minimal record
commit transaction
return API response
```

The audit outbox/intent write must be local, bounded, and durable.

Untracked goroutine audit writes are forbidden because they can fail invisibly after the product write succeeds.

## Audit Intent / Outbox Minimal Fields

The minimal audit intent/outbox record must include enough information to reconstruct or materialize the audit event later.

Required fields:

```text
id
bounded_context = product_catalog
aggregate_type = product
aggregate_id
event_name
operation
actor_id
actor_role
occurred_at
reason nullable
source_channel
revision_no
before_snapshot nullable
after_snapshot nullable
metadata_json
status = pending
created_at
```

`occurred_at` and `created_at` are server-owned timestamps.

`metadata_json` may include request id, correlation id, product version id, route/capability context, and compact domain-specific context.

## Snapshot Policy

Create:

```text
before_snapshot = null
after_snapshot = created product snapshot
```

Update:

```text
before_snapshot = product snapshot before update
after_snapshot = product snapshot after update
```

Soft delete:

```text
before_snapshot = active product snapshot
after_snapshot = deleted product snapshot
```

Restore:

```text
before_snapshot = deleted product snapshot
after_snapshot = restored active product snapshot
```

A future implementation may store before/after version references instead of full snapshots only if the audit worker can resolve those references durably and deterministically.

## Product Version Policy

Each ProductCatalog write operation must append one product version row.

Version events:

```text
product_created
product_updated
product_soft_deleted
product_restored
```

Each version snapshot must include product state and revision number.

`product_versions.changed_at` is server-owned.

## Forbidden Hot-Path Behavior

ProductCatalog write requests must not synchronously perform:

```text
external network logging
full telemetry export
dashboard or report generation
notification dispatch
slow audit enrichment
untracked goroutine audit writes
```

These belong outside the request hot path and must run through a worker/cold path when needed.

## Options Considered

### Option A: Full Synchronous Audit In The Request

The API writes product data and waits for full audit, telemetry, reporting, and enrichment before returning.

Rejected.

This gives strong immediate audit materialization but risks slowing normal API requests and violating the sub-1-second mission.

### Option B: Fire-And-Forget Async Audit

The API writes product data and starts asynchronous audit work without a durable outbox record in the same transaction.

Rejected.

This keeps the API fast but can lose audit evidence if the process crashes or the async task fails after the product write succeeds.

### Option C: Transactional Audit Intent / Outbox

The API writes product data, product version history, and a minimal audit intent/outbox record inside one transaction. Full audit materialization runs later.

Accepted.

This preserves durable audit evidence while keeping the hot path bounded.

## Consequences

### Positive

- Product writes and audit intent cannot diverge silently.
- API responses do not wait for full telemetry or report generation.
- Product history is available through `product_versions`.
- Audit materialization can evolve without changing ProductCatalog write semantics.
- The design supports a sub-1-second API mission.

### Negative

- Requires an audit outbox/intent table and worker later.
- Requires careful transaction boundaries in usecases.
- Requires tests proving product row, product version row, and audit intent rollback together.
- Adds more implementation steps before ProductCatalog is fully complete.

## Trade-off

The project accepts slightly more write-path structure in exchange for durable audit evidence and bounded request latency.

## Follow-up

- Add ProductCatalog ports/usecase contracts for transactional versioning and audit intent recording.
- Add usecase tests for create/update/soft-delete/restore version and audit intent behavior.
- Add PostgreSQL schema for products and product_versions in a later slice.
- Add audit_outbox or audit_intents schema in the storage slice where audit persistence is introduced.
- Add rollback tests proving product write, version write, and audit intent write commit or rollback together.
- Add background audit materialization worker in a later scope.
- Keep full telemetry, reporting projection, dashboard projection, and notifications outside ProductCatalog request hot path.
