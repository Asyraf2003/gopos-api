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

# API

This folder defines public HTTP API contract rules.

## Contents

- `0050_echo_http_contract.md`: Echo transport rules, request/response ownership, envelope expectations, route behavior, and forbidden HTTP-layer shortcuts.

## Use This Folder When

- adding or changing an HTTP endpoint;
- defining response envelopes or public DTO behavior;
- mapping application errors to public API errors;
- checking whether transport code is leaking domain or persistence logic.

API contracts are public contracts. Changes here usually need blueprint coverage, focused API tests, and sometimes an ADR.
