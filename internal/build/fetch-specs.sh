#!/usr/bin/env bash
set -euo pipefail

# fetch-specs.sh
# Downloads OpenAPI specifications from Catena Telematics API endpoints
# and stores them in the specs/ directory.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
SPECS_DIR="${REPO_ROOT}/specs"

# API spec URLs
INTEGRATIONS_URL="https://api.catenatelematics.com/v2/integrations/openapi.json"
ORGS_URL="https://api.catenatelematics.com/v2/orgs/openapi.json"
TELEMATICS_URL="https://api.catenatelematics.com/v2/telematics/openapi.json"
NOTIFICATIONS_URL="https://api.catenatelematics.com/v2/notifications/openapi.json"

echo "==> Fetching OpenAPI specifications..."

# Create specs directories if they don't exist
mkdir -p "${SPECS_DIR}/integrations"
mkdir -p "${SPECS_DIR}/orgs"
mkdir -p "${SPECS_DIR}/telematics"
mkdir -p "${SPECS_DIR}/notifications"

# Download each spec
echo "Downloading integrations spec..."
curl -fsSL "${INTEGRATIONS_URL}" -o "${SPECS_DIR}/integrations/openapi.json"

echo "Downloading orgs spec..."
curl -fsSL "${ORGS_URL}" -o "${SPECS_DIR}/orgs/openapi.json"

echo "Downloading telematics spec..."
curl -fsSL "${TELEMATICS_URL}" -o "${SPECS_DIR}/telematics/openapi.json"

echo "Downloading notifications spec..."
curl -fsSL "${NOTIFICATIONS_URL}" -o "${SPECS_DIR}/notifications/openapi.json"

echo "==> Normalizing specs (fixing duplicate tags)..."
"${SCRIPT_DIR}/fix-tags.py" "${SPECS_DIR}/integrations/openapi.json"
"${SCRIPT_DIR}/fix-tags.py" "${SPECS_DIR}/orgs/openapi.json"
"${SCRIPT_DIR}/fix-tags.py" "${SPECS_DIR}/telematics/openapi.json"
"${SCRIPT_DIR}/fix-tags.py" "${SPECS_DIR}/notifications/openapi.json"

echo "==> Successfully downloaded all OpenAPI specifications"
echo ""
echo "Downloaded specs:"
ls -lh "${SPECS_DIR}"/*/openapi.json
