#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

MSG="${MSG:-}"
if [[ -z "$MSG" ]]; then
  echo "[FAIL] MSG is required"
  echo 'usage: make push MSG="your commit message"'
  exit 1
fi

BRANCH="$(git rev-parse --abbrev-ref HEAD 2>/dev/null || true)"
if [[ -z "$BRANCH" || "$BRANCH" == "HEAD" ]]; then
  echo "[FAIL] unable to detect active git branch"
  exit 1
fi

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

echo
echo "[PASS] git push completed"
