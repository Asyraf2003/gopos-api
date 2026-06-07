#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

MANIFEST="scripts/config/route_capabilities.tsv"
SEED_FILES=(
  "migrations/0007_seed_existing_protected_capabilities.up.sql"
  "migrations/0008_seed_capability_manage_permission.up.sql"
)
SOURCE_FILES=(
  "internal/app/bootstrap/app.go"
  "internal/modules/system/transport/http/me_handler.go"
  "internal/modules/auth/transport/http/logout_handler.go"
  "internal/modules/auth/transport/http/account_role_handler.go"
  "internal/modules/capability/transport/http/capability_handler.go"
)

echo "== route capability audit =="
echo "manifest: $MANIFEST"
echo

fail=0

fail_msg() {
  echo "[FAIL] $*"
  fail=1
}

pass_msg() {
  echo "[PASS] $*"
}

require_file() {
  local file="$1"
  if [[ ! -f "$file" ]]; then
    fail_msg "missing file: $file"
  fi
}

contains_in_files() {
  local needle="$1"
  shift
  local file
  for file in "$@"; do
    if [[ -f "$file" ]] && grep -Fq "$needle" "$file"; then
      return 0
    fi
  done
  return 1
}

require_file "$MANIFEST"
for file in "${SEED_FILES[@]}"; do
  require_file "$file"
done
for file in "${SOURCE_FILES[@]}"; do
  require_file "$file"
done

if (( fail != 0 )); then
  echo
  echo "[FAIL] route capability audit failed"
  exit 1
fi

line_no=0
checked=0

while IFS=$'\t' read -r method path capability_key required_permission match source_pattern; do
  line_no=$((line_no + 1))

  [[ -z "${method:-}" ]] && continue
  [[ "$method" == \#* ]] && continue

  checked=$((checked + 1))

  case "$match" in
    exact|prefix) ;;
    *)
      fail_msg "$MANIFEST:$line_no invalid match value: $match"
      continue
      ;;
  esac

  if [[ "$method" != "*" ]]; then
    contains_in_files "'$method'" "${SEED_FILES[@]}" \
      || fail_msg "$capability_key missing seeded method: $method"
  fi

  contains_in_files "'$path'" "${SEED_FILES[@]}" \
    || fail_msg "$capability_key missing seeded path: $path"

  contains_in_files "'$capability_key'" "${SEED_FILES[@]}" \
    || fail_msg "$path missing seeded capability key: $capability_key"

  contains_in_files "'$required_permission'" "${SEED_FILES[@]}" \
    || fail_msg "$capability_key missing seeded permission: $required_permission"

  contains_in_files "$source_pattern" "${SOURCE_FILES[@]}" \
    || fail_msg "$capability_key missing route source pattern: $source_pattern"

  if [[ "$match" == "prefix" ]]; then
    contains_in_files "'*'" "${SEED_FILES[@]}" \
      || fail_msg "$capability_key prefix route must use wildcard method seed"
  fi

done < "$MANIFEST"

if (( checked == 0 )); then
  fail_msg "manifest has no route capability rows"
fi

echo
echo "checked route capability rows: $checked"

if (( fail != 0 )); then
  echo "[FAIL] route capability audit failed"
  exit 1
fi

echo "[PASS] route capability audit passed"
