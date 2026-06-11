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

MAX_LINES=100
ALLOWLIST_FILE="scripts/config/file_size_allowlist.txt"
MARKER_PREFIX="audit:allow-oversize"

if [[ ! -f "$ALLOWLIST_FILE" ]]; then
  echo "[FAIL] missing allowlist file: $ALLOWLIST_FILE"
  exit 1
fi

is_allowlisted() {
  local path="$1"
  grep -Fxq "$path" "$ALLOWLIST_FILE"
}

has_marker() {
  local path="$1"

  head -n 40 "$path" | rg -Fq "$MARKER_PREFIX"
}

count_lines() {
  local path="$1"

  awk '
    NR == 1 && $0 == "// Copyright (C) 2026 Asyraf Mubarak" {
      skipping_license_header = 1
      next
    }

    skipping_license_header == 1 && $0 ~ /^\/\// {
      next
    }

    skipping_license_header == 1 && $0 == "" {
      skipping_license_header = 0
      next
    }

    {
      count++
    }

    END {
      print count
    }
  ' "$path"
}

echo "== file size audit =="
echo "max lines: $MAX_LINES"
echo "allowlist: $ALLOWLIST_FILE"
echo "marker: $MARKER_PREFIX"
echo

fail=0

while IFS= read -r file; do
  [[ -z "$file" ]] && continue

  lines="$(count_lines "$file")"

  if (( lines <= MAX_LINES )); then
    continue
  fi

  allowlisted=false
  marked=false

  if is_allowlisted "$file"; then
    allowlisted=true
  fi

  if has_marker "$file"; then
    marked=true
  fi

  if [[ "$allowlisted" == true && "$marked" == true ]]; then
    echo "[WARN] allowlisted oversized file with marker: $file ($lines lines)"
    continue
  fi

  if [[ "$allowlisted" == true && "$marked" == false ]]; then
    echo "[FAIL] allowlisted oversized file missing marker: $file ($lines lines)"
    fail=1
    continue
  fi

  if [[ "$allowlisted" == false && "$marked" == true ]]; then
    echo "[FAIL] oversized file has marker but is not in allowlist: $file ($lines lines)"
    fail=1
    continue
  fi

  echo "[FAIL] oversized file: $file ($lines lines)"
  fail=1
done < <(fd -e go . internal | sort)

echo
if (( fail != 0 )); then
  echo "[FAIL] file size audit failed"
  exit 1
fi

echo "[PASS] file size audit passed"
