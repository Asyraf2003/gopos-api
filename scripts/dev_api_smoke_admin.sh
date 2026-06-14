#!/usr/bin/env bash
# Copyright (C) 2026 Asyraf Mubarak
#
# This file is part of gopos-api.
#
# gopos-api is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, version 3 only.
#
# gopos-api is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

ENV_HTTP_PORT="${HTTP_PORT:-}"
ENV_API_BASE_URL="${API_BASE_URL:-}"
ENV_AUTH_DEBUG_ENABLED="${AUTH_DEBUG_ENABLED:-}"

if [[ -f ".env" ]]; then
  set -a
  source .env
  set +a
fi

if [[ -n "$ENV_HTTP_PORT" ]]; then
  HTTP_PORT="$ENV_HTTP_PORT"
fi

if [[ -n "$ENV_API_BASE_URL" ]]; then
  API_BASE_URL="$ENV_API_BASE_URL"
fi

if [[ -n "$ENV_AUTH_DEBUG_ENABLED" ]]; then
  AUTH_DEBUG_ENABLED="$ENV_AUTH_DEBUG_ENABLED"
fi

HTTP_PORT="${HTTP_PORT:-8081}"
API_BASE_URL="${API_BASE_URL:-http://127.0.0.1:${HTTP_PORT}}"
AUTH_DEBUG_ENABLED="${AUTH_DEBUG_ENABLED:-false}"

ADMIN_EMAIL="${AUTH_MANUAL_ADMIN_EMAIL:-admin@example.com}"
ADMIN_PASSWORD="${AUTH_MANUAL_ADMIN_PASSWORD:-12345678}"

if [[ "$AUTH_DEBUG_ENABLED" != "true" ]]; then
  echo "[WARN] AUTH_DEBUG_ENABLED is not true."
  echo "[WARN] Manual login may be disabled."
fi

tmp_login_body="$(mktemp)"
tmp_request_body="$(mktemp)"
tmp_login_payload="$(mktemp)"
trap 'rm -f "$tmp_login_body" "$tmp_request_body" "$tmp_login_payload"' EXIT

python3 - "$tmp_login_payload" "$ADMIN_EMAIL" "$ADMIN_PASSWORD" <<'PYJSON'
import json
import sys

path, email, password = sys.argv[1], sys.argv[2], sys.argv[3]
with open(path, "w", encoding="utf-8") as f:
    json.dump({"email": email, "password": password}, f)
PYJSON

echo "== admin api smoke =="
echo "api: ${API_BASE_URL}"
echo

echo "== health =="
health_status="$(
  curl -sS     -o "$tmp_request_body"     -w '%{http_code}'     "${API_BASE_URL}/api/health"
)"
echo "health status=${health_status}"
if [[ "$health_status" -lt 200 || "$health_status" -ge 300 ]]; then
  echo "[FAIL] /api/health failed"
  cat "$tmp_request_body"
  echo
  exit 1
fi

echo "== manual admin login =="
login_status="$(
  curl -sS     -o "$tmp_login_body"     -w '%{http_code}'     -X POST "${API_BASE_URL}/api/auth/manual/login"     -H 'Content-Type: application/json'     -d "{"email":"${ADMIN_EMAIL}","password":"${ADMIN_PASSWORD}"}"
)"
echo "manual_login status=${login_status}"
if [[ "$login_status" != "200" ]]; then
  echo "[FAIL] manual admin login failed"
  cat "$tmp_login_body"
  echo
  exit 1
fi

access_token="$(
  python3 - "$tmp_login_body" <<'PYJSON'
import json
import sys

with open(sys.argv[1], "r", encoding="utf-8") as f:
    payload = json.load(f)

token = payload.get("access_token") or payload.get("accessToken") or ""
if not token:
    raise SystemExit("[FAIL] access token not found in login response")

print(token)
PYJSON
)"

echo "access_token_present=true"
echo

request_json() {
  local name="$1"
  local path="$2"

  local status
  status="$(
    curl -sS       -o "$tmp_request_body"       -w '%{http_code}'       "${API_BASE_URL}${path}"       -H "Authorization: Bearer ${access_token}"       -H 'Accept: application/json'
  )"

  echo "${name} status=${status}"

  if [[ "$status" -lt 200 || "$status" -ge 300 ]]; then
    echo "[FAIL] ${name} failed"
    cat "$tmp_request_body"
    echo
    exit 1
  fi
}

echo "== authenticated checks =="
request_json "me" "/api/me"
request_json "products_list" "/api/products"
request_json "products_lookup" "/api/products/lookup"

echo
echo "[PASS] admin api smoke passed"
