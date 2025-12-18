#!/usr/bin/env bash
set -euo pipefail

# verify-clean.sh
# Verifies that generated code matches what's committed in git.
# Used in CI to ensure the repository isn't drifting from the specs.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

cd "${REPO_ROOT}"

echo "==> Verifying generated code is up-to-date..."

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo "Not in a git repository, skipping verification"
    exit 0
fi

# Check for uncommitted changes in generated code
if ! git diff --exit-code gen/ specs/combined/; then
    echo ""
    echo "ERROR: Generated code has uncommitted changes!"
    echo ""
    echo "The generated code in gen/ or specs/combined/ doesn't match what's committed."
    echo "This usually means:"
    echo "  1. You modified generated code manually (don't do this), or"
    echo "  2. You need to run 'make generate' and commit the changes, or"
    echo "  3. The specs have changed and generation needs to be re-run"
    echo ""
    echo "To fix this:"
    echo "  make generate"
    echo "  git add gen/ specs/"
    echo "  git commit -m 'Regenerate client code'"
    echo ""
    exit 1
fi

echo "==> Verification passed! Generated code is clean."
