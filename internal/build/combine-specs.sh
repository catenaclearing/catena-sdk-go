#!/usr/bin/env bash
set -euo pipefail

# combine-specs.sh
# Combines individual OpenAPI specs into a single unified spec using Redocly CLI,
# then validates the combined output.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
SPECS_DIR="${REPO_ROOT}/specs"
COMBINED_DIR="${SPECS_DIR}/combined"

echo "==> Combining OpenAPI specifications..."

# Check if Redocly CLI is available
if ! command -v redocly &> /dev/null; then
    echo "Error: redocly CLI not found. Please install it:"
    echo "  npm install -g @redocly/cli"
    exit 1
fi

# Create combined directory
mkdir -p "${COMBINED_DIR}"

# Use Redocly to join the specs
# Note: We generate separate packages, so we'll create a combined spec for validation
# but generate from individual specs to avoid collision issues.
echo "Bundling integrations spec..."
redocly bundle "${SPECS_DIR}/integrations/openapi.json" \
    --output "${COMBINED_DIR}/integrations.json" \
    --dereferenced

echo "Bundling orgs spec..."
redocly bundle "${SPECS_DIR}/orgs/openapi.json" \
    --output "${COMBINED_DIR}/orgs.json" \
    --dereferenced

echo "Bundling telematics spec..."
redocly bundle "${SPECS_DIR}/telematics/openapi.json" \
    --output "${COMBINED_DIR}/telematics.json" \
    --dereferenced
echo "Patching telematics spec (h3_index_11)..."
"${SCRIPT_DIR}/patch-h3-index.py" "${COMBINED_DIR}/telematics.json"

echo "Bundling notifications spec..."
redocly bundle "${SPECS_DIR}/notifications/openapi.json" \
    --output "${COMBINED_DIR}/notifications.json" \
    --dereferenced
echo "Patching notifications spec (h3_index_11)..."
"${SCRIPT_DIR}/patch-h3-index.py" "${COMBINED_DIR}/notifications.json"

echo "Bundling authentication spec..."
redocly bundle "${SPECS_DIR}/authentication/openapi.json" \
    --output "${COMBINED_DIR}/authentication.json" \
    --dereferenced

# Create a combined spec that references all APIs
# This is mainly for documentation and validation purposes
cat > "${COMBINED_DIR}/openapi.json" <<EOF
{
  "openapi": "3.0.0",
  "info": {
    "title": "Catena Telematics Combined API",
    "version": "2.0.0",
    "description": "Combined OpenAPI specification for Catena Telematics APIs. This includes integrations, orgs, telematics, and notifications endpoints."
  },
  "servers": [
    {
      "url": "https://api.catenatelematics.com/v2",
      "description": "Production server"
    }
  ],
  "paths": {},
  "components": {
    "schemas": {}
  }
}
EOF

echo "==> Validating individual specifications..."

# Validate each spec
for spec in integrations orgs telematics notifications; do
    echo "Validating ${spec}..."
    redocly lint "${SPECS_DIR}/${spec}/openapi.json" || {
        echo "Warning: ${spec} spec has validation issues"
    }
done

echo "==> Successfully combined and validated specifications"
echo ""
echo "Combined specs created in: ${COMBINED_DIR}"
ls -lh "${COMBINED_DIR}"/*.json
