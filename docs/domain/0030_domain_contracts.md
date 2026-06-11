// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

// audit:allow-oversize reason=bootstrap-wiring

# Domain Contracts

## Purpose

Every database-backed domain must declare its lifecycle and operations before implementation.

## Required Domain Contract

Each domain must define:

- domain name;
- source table or tables;
- aggregate root if any;
- owned entities;
- read models;
- allowed operations;
- forbidden operations;
- lifecycle statuses;
- delete policy;
- audit policy;
- idempotency policy;
- authorization policy;
- capability keys;
- invariants;
- transaction boundary;
- expected tests.

## Default CRUD Surface

Every normal master-data domain should provide:

- create;
- update/edit;
- delete;
- show;
- list.

If an operation is not allowed, the domain contract must say why.

## Delete Rule

Delete is not automatically physical delete.

Each domain must choose one:

- physical delete;
- soft delete;
- archive;
- cancel;
- void;
- forbidden.

High-risk financial, stock, audit, payment, refund, and posted transaction records must not be physically deleted unless an explicit owner decision allows it.

## Mutation Rule

Every create, update, delete, cancel, close, refund, allocate, or reversal operation must define:

- command DTO;
- result DTO;
- validation;
- transaction boundary;
- lock target if needed;
- idempotency key if retryable;
- audit event;
- repository methods;
- API capability key;
- tests.

## UI Dynamic Rule

UI must read available operations from API/capability metadata. UI must not hard-code business authority as the source of truth.

