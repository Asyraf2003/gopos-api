#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

ENV_HTTP_PORT="${HTTP_PORT:-}"
ENV_API_BASE_URL="${API_BASE_URL:-}"

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

HTTP_PORT="${HTTP_PORT:-8081}"
API_BASE_URL="${API_BASE_URL:-http://127.0.0.1:${HTTP_PORT}}"

echo "== dev smoke =="
echo "api: ${API_BASE_URL}"
echo

echo "== /api/health =="
curl -i "${API_BASE_URL}/api/health"
echo
echo

echo "== /api/me without token =="
curl -i "${API_BASE_URL}/api/me"
echo
