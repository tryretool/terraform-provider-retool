#!/usr/bin/env bash
#
# Record acceptance-test HTTP interactions ONE TEST AT A TIME.
#
# The Retool instance enforces a rolling API rate limit, so recording the whole
# suite in one process (even serially) exhausts the budget partway through and
# the remaining tests fail with 429s. This script records each acceptance test
# in its own `go test` invocation with a pause between them so the rate limit
# can recover.
#
# It is resumable: tests whose cassette already exists are skipped, so if a run
# is interrupted (or some tests 429), just run it again to finish the rest.
#
# Requirements: go on PATH, plus RETOOL_HOST / RETOOL_SCHEME / RETOOL_ACCESS_TOKEN.
#
# Usage:
#   RETOOL_HOST=... RETOOL_SCHEME=https RETOOL_ACCESS_TOKEN=... \
#     scripts/record_acc_tests.sh [TEST_NAME_REGEX]
#
# Options (env vars):
#   SLEEP_SECONDS   Pause between tests (default 30).
#   FORCE=1         Re-record even if a cassette already exists.
#
# Examples:
#   scripts/record_acc_tests.sh                 # record everything not yet recorded
#   scripts/record_acc_tests.sh TestAccSSO      # record only matching tests
#   SLEEP_SECONDS=60 scripts/record_acc_tests.sh
set -uo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

PKGS="./internal/provider/..."
RECORDINGS_DIR="test/data/recordings"
SLEEP_SECONDS="${SLEEP_SECONDS:-30}"
NAME_REGEX="${1:-TestAcc}"

if [ -z "${RETOOL_HOST:-}" ] || [ -z "${RETOOL_ACCESS_TOKEN:-}" ]; then
  echo "ERROR: set RETOOL_HOST, RETOOL_SCHEME and RETOOL_ACCESS_TOKEN" >&2
  exit 1
fi

mkdir -p "$RECORDINGS_DIR"

# Discover acceptance tests as "package|testname" pairs so each test runs against
# only its own package (avoids the noisy "no tests to run" output from the other
# provider packages).
PAIRS=""
for pkg in $(go list $PKGS 2>/dev/null); do
  for name in $(go test -list "${NAME_REGEX}.*" "$pkg" 2>/dev/null | grep -E '^TestAcc'); do
    PAIRS="${PAIRS}${pkg}|${name}
"
  done
done
PAIRS="$(printf '%s' "$PAIRS" | grep -v '^$' | sort -u)"

if [ -z "$PAIRS" ]; then
  echo "No acceptance tests matching /${NAME_REGEX}/ found." >&2
  exit 1
fi

total="$(printf '%s\n' "$PAIRS" | wc -l | tr -d ' ')"
echo "Found ${total} acceptance test(s) matching /${NAME_REGEX}/."
echo "Sleeping ${SLEEP_SECONDS}s between tests to respect the API rate limit."
echo

idx=0
recorded=0
skipped=0
FAILED=""
while IFS= read -r pair; do
  [ -z "$pair" ] && continue
  pkg="${pair%%|*}"
  name="${pair##*|}"
  idx=$((idx + 1))
  cassette="${RECORDINGS_DIR}/${name}.yaml"

  if [ "${FORCE:-0}" != "1" ] && [ -s "$cassette" ]; then
    echo "==> [${idx}/${total}] SKIP ${name} (cassette already exists)"
    skipped=$((skipped + 1))
    continue
  fi

  echo "==> [${idx}/${total}] RECORD ${name}"
  rm -f "$cassette"
  if RETOOL_HTTP_RECORDINGS=on TF_ACC=1 \
      go test -v -p 1 -parallel 1 -run "^${name}$" -timeout 30m "$pkg"; then
    recorded=$((recorded + 1))
  else
    echo "!! ${name} failed to record (likely 429); removing partial cassette so it can be retried."
    rm -f "$cassette"
    FAILED="${FAILED} ${name}"
  fi

  if [ "$idx" -lt "$total" ]; then
    echo "    sleeping ${SLEEP_SECONDS}s..."
    sleep "$SLEEP_SECONDS"
  fi
done <<EOF
$PAIRS
EOF

echo
echo "Done. recorded=${recorded} skipped=${skipped}"
if [ -n "$FAILED" ]; then
  echo "Failed (re-run this script to retry):${FAILED}"
  exit 1
fi
echo "All acceptance tests recorded."
