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

# Blueprints

This folder contains active implementation plans.

Rules:
- One blueprint owns one scope.
- Blueprints are not handoffs.
- Blueprints should stay readable as plans: keep ongoing logs in handoffs and progress ledgers.
- Closed blueprints may keep compact closeout summaries that link to proof-bearing handoffs or ledgers.
- Blueprints must define scope, risks, architecture impact, DB impact, API contract impact, capability keys, tests,
and step order.
- Do not implement from chat-only decisions.
