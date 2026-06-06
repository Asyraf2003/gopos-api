#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

BRANCH="$(git rev-parse --abbrev-ref HEAD 2>/dev/null || true)"
if [[ -z "$BRANCH" || "$BRANCH" == "HEAD" ]]; then
  echo "[FAIL] unable to detect active git branch"
  exit 1
fi

COUNTER_FILE=".git/.auto_commit_counter"

if [[ -f "$COUNTER_FILE" ]]; then
  LAST_NUMBER="$(cat "$COUNTER_FILE")"
else
  LAST_NUMBER="0"
fi

if ! [[ "$LAST_NUMBER" =~ ^[0-9]+$ ]]; then
  echo "[FAIL] invalid auto commit counter: $LAST_NUMBER"
  exit 1
fi

NEXT_NUMBER=$((LAST_NUMBER + 1))
MSG="commit $NEXT_NUMBER"

echo "== git push =="
echo "branch: $BRANCH"
echo "message: $MSG"
echo

echo "-- git status --"
git status --short
echo

git add .

if git diff --cached --quiet; then
  echo "[FAIL] no staged changes to commit"
  exit 1
fi

git commit -m "$MSG"
git push origin "$BRANCH"

echo "$NEXT_NUMBER" > "$COUNTER_FILE"

echo
echo "[PASS] git push completed"
echo "[PASS] commit message: $MSG"
