#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

fail=0

check_no_imports() {
  local label="$1"
  local path_glob="$2"
  local pattern="$3"

  echo "-- $label --"
  local matches
  matches="$(rg -n "$pattern" internal/modules -g "$path_glob" -g '!**/*_test.go' || true)"

  if [[ -n "$matches" ]]; then
    echo "$matches"
    fail=1
    echo "[FAIL] $label"
    echo
    return
  fi

  echo "[OK] $label"
  echo
}

echo "== hexagonal import audit =="
echo

check_no_imports \
  "domain must not import transport, platform, SQL, HTTP, config, or app" \
  "**/domain/*.go" \
  '"(github.com/labstack/echo|net/http|database/sql|github.com/jackc/pgx|pos-go/internal/(platform|transport|config|app)|pos-go/internal/modules/.*/(transport|store))'

check_no_imports \
  "ports must not import transport, platform, SQL, HTTP, config, or app" \
  "**/ports/*.go" \
  '"(github.com/labstack/echo|net/http|database/sql|github.com/jackc/pgx|pos-go/internal/(platform|transport|config|app)|pos-go/internal/modules/.*/(transport|store))'

check_no_imports \
  "usecase must not import transport, SQL adapters, HTTP, config, or app" \
  "**/usecase/*.go" \
  '"(github.com/labstack/echo|net/http|database/sql|github.com/jackc/pgx|pos-go/internal/(config|app)|pos-go/internal/transport|pos-go/internal/modules/.*/(transport|store))'

check_no_imports \
  "http transport must not import SQL drivers or persistence adapters" \
  "**/transport/http/*.go" \
  '"(database/sql|github.com/jackc/pgx|pos-go/internal/platform/postgres|pos-go/internal/modules/.*/store)'

echo "-- one package per folder --"
while IFS= read -r dir; do
  packages="$(awk '/^package / { print $2 }' "$dir"/*.go 2>/dev/null | sort -u | wc -l | tr -d ' ')"
  if [[ "$packages" != "1" ]]; then
    echo "[FAIL] multiple packages in $dir"
    fail=1
  fi
done < <(fd -e go . internal -x dirname | sort -u)

if (( fail != 0 )); then
  echo
  echo "[FAIL] hexagonal import audit failed"
  exit 1
fi

echo "[PASS] hexagonal import audit passed"
