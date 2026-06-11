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

# Security

This folder defines baseline security rules.

## Contents

- `0080_security_baseline.md`: authentication, authorization, capability, input, output, secret, audit, and abuse-control rules.

## Use This Folder When

- adding protected APIs;
- changing authn, authz, roles, permissions, or capabilities;
- handling secrets, tokens, input validation, or audit decisions;
- deciding whether a mutation is safe to expose.

Security rules are P0 constraints when they intersect with API exposure or state mutation.
