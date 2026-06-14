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
ROLE="${1:-admin}"

case "$ROLE" in
  admin)
    EMAIL="${AUTH_MANUAL_ADMIN_EMAIL:-admin@example.com}"
    PASSWORD="${AUTH_MANUAL_ADMIN_PASSWORD:-12345678}"
    ;;
  cashier|kasir)
    EMAIL="${AUTH_MANUAL_CASHIER_EMAIL:-kasir@example.com}"
    PASSWORD="${AUTH_MANUAL_CASHIER_PASSWORD:-12345678}"
    ;;
  *)
    echo "[FAIL] unsupported role: ${ROLE}"
    echo "Usage: scripts/dev_manual_login.sh admin|cashier"
    exit 1
    ;;
esac

if [[ "${AUTH_DEBUG_ENABLED:-false}" != "true" ]]; then
  echo "[WARN] AUTH_DEBUG_ENABLED is not true in environment/.env"
  echo "Manual login route may be disabled."
  echo
fi

tmp_body="$(mktemp)"
tmp_headers="$(mktemp)"
tmp_payload="$(mktemp)"
trap 'rm -f "$tmp_body" "$tmp_headers" "$tmp_payload"' EXIT

python3 - "$tmp_payload" "$EMAIL" "$PASSWORD" <<'PYJSON'
import json
import sys

path, email, password = sys.argv[1], sys.argv[2], sys.argv[3]
with open(path, "w", encoding="utf-8") as f:
    json.dump({"email": email, "password": password}, f)
PYJSON

echo "== manual login =="
echo "api: ${API_BASE_URL}"
echo "email: ${EMAIL}"
echo

status="$(
  curl -sS \
    -o "$tmp_body" \
    -D "$tmp_headers" \
    -w '%{http_code}' \
    -X POST "${API_BASE_URL}/api/auth/manual/login" \
    -H 'Content-Type: application/json' \
    --data-binary "@${tmp_payload}"
)"

echo "HTTP_STATUS=${status}"
echo
cat "$tmp_headers"
cat "$tmp_body"
echo
echo

if [[ "$status" != "200" ]]; then
  echo "[FAIL] manual login failed"
  exit 1
fi

access_token="$(
  python3 - "$tmp_body" <<'PY'
import json
import sys

path = sys.argv[1]
with open(path, "r", encoding="utf-8") as f:
    data = json.load(f)

token = data.get("access_token") or data.get("accessToken") or ""
if not token:
    raise SystemExit("[FAIL] access token not found in response")

print(token)
PY
)"

echo "ACCESS_TOKEN=${access_token}"
echo
echo "== /api/me with bearer token =="
curl -i \
  "${API_BASE_URL}/api/me" \
  -H "Authorization: Bearer ${access_token}"
echo
