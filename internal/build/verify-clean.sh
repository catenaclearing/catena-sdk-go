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

# Check for uncommitted changes in tracked generated code
mapfile -t tracked_files < <(git ls-files gen/ specs/combined/)
if [ ${#tracked_files[@]} -eq 0 ]; then
    echo "No tracked generated files found in gen/ or specs/combined; skipping verification."
    exit 0
fi

if ! git diff --exit-code -- "${tracked_files[@]}"; then
    echo ""
    echo "ERROR: Generated code has uncommitted changes!"
    echo ""
    echo "Tracked generated files in gen/ or specs/combined/ don't match what's committed."
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
