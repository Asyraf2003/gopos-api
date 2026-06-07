# ServiceCatalog Implementation Slice 1 Blueprint

## Status

Closed / implemented with proof.

## Date

2026-06-08

## Active Scope

ServiceCatalog implementation slice 1 after accepted domain contract.

## Accepted Contract

```text
docs/blueprints/0024_servicecatalog_domain_contract.md
```

## FACT

- ServiceCatalog domain contract is accepted.
- Capability-control foundation is closed.
- ServiceCatalog slice 1 is implemented.
- Implementation exists only in the allowed slice 1 packages:
  - `internal/modules/servicecatalog/domain`
  - `internal/modules/servicecatalog/ports`
  - `internal/modules/servicecatalog/usecase`
- No HTTP transport was added in this slice.
- No PostgreSQL adapter was added in this slice.
- No PostgreSQL migration was added in this slice.
- No route registration was added in this slice.
- No capability seed migration was added in this slice.
- Business Phase 1 is partial because ServiceCatalog slice 1 exists, but ServiceCatalog HTTP/PostgreSQL integration and ProductCatalog are not implemented.

## DECISION

Slice 1 implemented only ServiceCatalog domain and usecase contracts with tests.

HTTP transport, PostgreSQL adapter, migrations, route registration, and capability seeds remain deferred to the next accepted blueprint.

## IMPLEMENTED SCOPE

- `internal/modules/servicecatalog/domain`
- `internal/modules/servicecatalog/ports`
- `internal/modules/servicecatalog/usecase`
- Unit tests for domain and usecase behavior

## IMPLEMENTED FILES

```text
internal/modules/servicecatalog/domain/errors.go
internal/modules/servicecatalog/domain/normalizer.go
internal/modules/servicecatalog/domain/service_catalog_item_behavior.go
internal/modules/servicecatalog/domain/service_catalog_item.go
internal/modules/servicecatalog/domain/service_catalog_item_test.go
internal/modules/servicecatalog/domain/validation.go
internal/modules/servicecatalog/ports/service_catalog_repository.go
internal/modules/servicecatalog/usecase/activate_item.go
internal/modules/servicecatalog/usecase/create_item.go
internal/modules/servicecatalog/usecase/create_item_test.go
internal/modules/servicecatalog/usecase/deactivate_item.go
internal/modules/servicecatalog/usecase/errors.go
internal/modules/servicecatalog/usecase/fake_repository_helpers_test.go
internal/modules/servicecatalog/usecase/fake_repository_query_test.go
internal/modules/servicecatalog/usecase/fake_repository_state_test.go
internal/modules/servicecatalog/usecase/fake_repository_test.go
internal/modules/servicecatalog/usecase/lifecycle_item_test.go
internal/modules/servicecatalog/usecase/list_items.go
internal/modules/servicecatalog/usecase/list_item_test.go
internal/modules/servicecatalog/usecase/lookup_items.go
internal/modules/servicecatalog/usecase/lookup_item_test.go
internal/modules/servicecatalog/usecase/show_item.go
internal/modules/servicecatalog/usecase/types.go
internal/modules/servicecatalog/usecase/update_item.go
internal/modules/servicecatalog/usecase/update_item_test.go
```

## DOMAIN TYPES

- `ServiceCatalogItem`
- `ServiceCatalogItemStatus`
- `ServiceCatalogItemID`
- `NormalizedName`
- `MoneyRupiah`

Minimum fields:

- `id`
- `name`
- `normalized_name`
- `default_price_rupiah`
- `is_active`
- `created_at`
- `updated_at`

## DOMAIN RULES PROVEN

- Name is required after trim.
- Repeated internal whitespace is compacted.
- Normalized name is lowercase.
- Default price must be greater than zero.
- New item is active by default.
- Deactivate keeps item stored.
- Activate restores active status.
- Physical delete is forbidden.

## PORTS IMPLEMENTED

Repository port supports:

- `Create`
- `Update`
- `FindByID`
- `FindByNormalizedName`
- `List`
- `Lookup`
- `SetActive`

Transactor port remains deferred.

## USECASES IMPLEMENTED

- `CreateServiceCatalogItem`
- `UpdateServiceCatalogItem`
- `ActivateServiceCatalogItem`
- `DeactivateServiceCatalogItem`
- `ShowServiceCatalogItem`
- `ListServiceCatalogItems`
- `LookupServiceCatalogItems`

## TEST MATRIX PROVEN

Domain tests:

- Normalize trims whitespace.
- Normalize compacts repeated whitespace.
- Normalize lowercases.
- Reject blank name.
- Reject zero default price.
- Reject negative default price.
- Create active item by default.

Usecase tests:

- Create stores item.
- Create rejects duplicate normalized name.
- Update changes name and price.
- Update rejects duplicate normalized name.
- Activate marks inactive item active.
- Deactivate marks active item inactive.
- Show missing item returns not found.
- List filters active/inactive/all.
- Lookup excludes inactive by default.
- Lookup enforces max limit.

## PROOF COLLECTED

```text
go test ./internal/modules/servicecatalog/...
```

Result:

```text
ok  	pos-go/internal/modules/servicecatalog/domain
?   	pos-go/internal/modules/servicecatalog/ports	[no test files]
ok  	pos-go/internal/modules/servicecatalog/usecase
```

```text
make verify
```

Result:

```text
[PASS] go test ./...
[PASS] go vet audit
[PASS] format audit
[PASS] AI rules audit
[PASS] file size audit
[PASS] hexagonal import audit
[PASS] route capability audit
[PASS] security gosec audit
[PASS] aggregate audit passed
```

Security proof:

```text
Gosec  : dev
Files  : 112
Lines  : 4659
Nosec  : 0
Issues : 0
```

## SCOPE-OUT CONFIRMED

Still not implemented in this slice:

- HTTP transport
- PostgreSQL adapter
- PostgreSQL migrations
- Route registration
- Capability seed migrations
- ProductCatalog
- Inventory

## ACCEPTANCE

ServiceCatalog implementation slice 1 is accepted and closed with proof on 2026-06-08.

## NEXT ACTIVE STEP

Plan the next ServiceCatalog implementation slice.

Candidate scope:

ServiceCatalog HTTP transport, PostgreSQL adapter, migrations, route registration, and capability seeds.

Do not implement the next scope until a new accepted blueprint defines:

- Exact files
- Route and capability mapping
- Persistence schema
- Transaction behavior
- Audit behavior
- Authorization behavior
- Proof commands
