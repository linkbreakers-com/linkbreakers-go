#!/usr/bin/env bash
# Compiles every Go block in README.md against the generated package, so the
# hand-written README cannot call methods the generator does not produce.
# Each block must be a complete program (package main).
set -euo pipefail

root="$(cd "$(dirname "$0")/.." && pwd)"
work="$root/docscheck_tmp"
trap 'rm -rf "$work"' EXIT
rm -rf "$work"

awk -v dir="$work" '
  /^```go$/ { n++; inside = 1; file = sprintf("%s/block%d/main.go", dir, n); system("mkdir -p " dir "/block" n); next }
  /^```$/ && inside { inside = 0; close(file); next }
  inside { print > file }
' "$root/README.md"

count=$(find "$work" -name main.go 2>/dev/null | wc -l | tr -d ' ')
if [ "$count" = "0" ]; then
  echo "No Go blocks found in README.md" >&2
  exit 1
fi

cd "$root"
if ! go vet ./docscheck_tmp/...; then
  echo "A README.md Go block no longer compiles against the generated SDK." >&2
  exit 1
fi
echo "Compiled $count README.md Go blocks against the generated SDK."
