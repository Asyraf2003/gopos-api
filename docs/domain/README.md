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

# Domain

This folder defines domain contract rules.

## Contents

- `0030_domain_contracts.md`: required create, update, delete, show, list, authorization, capability, audit, transaction, and test declarations for database-backed domains.

## Use This Folder When

- creating a new business domain;
- adding CRUD behavior;
- forbidding or constraining a domain operation;
- mapping domain lifecycle rules before API or database implementation.

Do not implement domain CRUD without an explicit domain contract and capability keys.
