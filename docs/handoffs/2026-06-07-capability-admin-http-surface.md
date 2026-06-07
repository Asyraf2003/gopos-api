# Handoff: Capability Admin HTTP Surface

## Date

2026-06-07

## Active Scope

Add admin capability HTTP surface.

## Scope In

- Add `capability.manage` permission.
- Assign `capability.manage` to `admin` role.
- Add admin capability list/show/enable/disable HTTP handlers.
- Wire admin routes behind authn, `capability.manage` authz, and capability runtime check.
- Keep POS business CRUD out of scope.

## Scope Out

- Products CRUD.
- Sale order workflow.
- Inventory.
- Payments.
- Route-to-capability audit script.

## Planned Routes

```text
GET  /api/admin/capabilities
GET  /api/admin/capabilities/:key
POST /api/admin/capabilities/:key/enable
POST /api/admin/capabilities/:key/disable
```

## Proof To Collect

- `make dev`
- Permission SQL proof
- Role permission SQL proof
- `capability.manage` SQL proof
- `go test ./internal/modules/capability/...`
- `go test ./internal/modules/capability/transport/http/...`
- `go test ./internal/app/bootstrap/...`
- `make verify`

## Open Gaps After This Step

- Route-to-capability audit script.
- Route-level disabled protected endpoint proof.
- POS domain CRUD remains blocked until capability-control foundation proof is complete.

## NEXT sekarang

Jalankan validasi docs lokal vs remote dulu:

```bash
cd /home/asyraf/Code/go/pos-go

git status --short
git log --oneline -5
git pull --ff-only

grep -n "SQL seed proof\|aggregate audit passed\|existing protected route capability records\|Admin capability HTTP surface" \
  docs/handoffs/2026-06-07-capability-route-seeds.md \
  docs/evidence/0003_laravel_to_go_transition_progress_ledger.md
```

Setelah output itu jelas, baru masuk implement migration `0008`.

Jangan loncat bikin CRUD, nanti capability foundation cuma jadi hiasan arsitektur yang mahal dan tidak berguna.
