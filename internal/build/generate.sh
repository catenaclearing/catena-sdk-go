#!/usr/bin/env bash
set -euo pipefail

# generate.sh
# Generates Go client code from OpenAPI specifications using OpenAPI Generator.
# Uses Docker for reproducible generation with pinned version.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
SPECS_DIR="${REPO_ROOT}/specs"
GEN_DIR="${REPO_ROOT}/gen"

# Pin OpenAPI Generator version for reproducibility
OPENAPI_GENERATOR_VERSION="v7.11.0"
OPENAPI_GENERATOR_IMAGE="openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION}"

# Module name for generated code
MODULE_NAME="github.com/catenaclearing/catena-sdk-go"

echo "==> Generating Go client code from OpenAPI specifications..."
echo "Using OpenAPI Generator: ${OPENAPI_GENERATOR_IMAGE}"

# Check if Docker is available
if ! command -v docker &> /dev/null; then
    echo "Error: Docker not found. Please install Docker to run the generator."
    exit 1
fi

# Pull the generator image
echo "Pulling OpenAPI Generator Docker image..."
docker pull "${OPENAPI_GENERATOR_IMAGE}"

# Function to generate client for a specific API
generate_client() {
    local api_name=$1
    local package_name="${api_name}api"
    local module_name="${MODULE_NAME}/gen/${api_name}"
    local input_spec="${SPECS_DIR}/combined/${api_name}.json"
    local output_dir="${GEN_DIR}/${api_name}"
    
    echo ""
    echo "==> Generating ${api_name} client..."
    
    # Check if bundled spec exists, if not use original
    if [ ! -f "${input_spec}" ]; then
        input_spec="${SPECS_DIR}/${api_name}/openapi.json"
        echo "Using original spec: ${input_spec}"
    fi
    
    # Clean output directory
    rm -rf "${output_dir}"
    mkdir -p "${output_dir}"
    
    # Create .openapi-generator-ignore file to reduce noise
    cat > "${output_dir}/.openapi-generator-ignore" <<EOF
# Don't generate these files
.travis.yml
git_push.sh
.gitlab-ci.yml
.gitignore
go.mod
go.sum
EOF
    
    # Run OpenAPI Generator
    docker run --rm \
        --user "$(id -u):$(id -g)" \
        -v "${REPO_ROOT}:/local" \
        "${OPENAPI_GENERATOR_IMAGE}" generate \
        --skip-validate-spec \
        -i "/local/specs/${api_name}/openapi.json" \
        -g go \
        -o "/local/gen/${api_name}" \
        --additional-properties=packageName="${package_name}" \
        --additional-properties=moduleName="${module_name}" \
        --additional-properties=generateGoMod=false \
        --additional-properties=enumClassPrefix=true \
        --additional-properties=generateInterfaces=true \
        --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
        --git-user-id=catenaclearing \
        --git-repo-id=catena-sdk-go \
        --package-name="${package_name}"

    # Ensure submodule metadata isn't created in generated code.
    rm -f "${output_dir}/go.mod" "${output_dir}/go.sum"

    # Normalize imports in generated docs/tests to use the gen/<api> path.
    python - <<PY
from pathlib import Path

root = Path("${output_dir}")
old = "github.com/catenaclearing/catena-sdk-go"
new = "${MODULE_NAME}/gen/${api_name}"

for path in root.rglob("*"):
    if path.suffix not in {".go", ".md"}:
        continue
    try:
        text = path.read_text(encoding="utf-8")
    except Exception:
        continue
    if old not in text:
        continue
    path.write_text(text.replace(old, new), encoding="utf-8")
PY
    
    # Add generated marker comment to main files
    local api_file="${output_dir}/api_*.go"
    if ls ${api_file} 1> /dev/null 2>&1; then
        echo "Generated ${api_name} client successfully"
    else
        echo "Warning: No API files generated for ${api_name}"
    fi
}

# Generate clients for each API
for api in integrations orgs telematics notifications authentication; do
    generate_client "${api}"
done

echo ""
echo "==> Code generation complete!"
echo ""
echo "Generated packages:"
find "${GEN_DIR}" -name "go.mod" -exec dirname {} \;

echo ""
echo "==> Running gofmt on generated code..."
find "${GEN_DIR}" -name "*.go" -exec gofmt -w {} \;

echo ""
echo "==> Generating pagination wrappers..."
# Clean up old root-level pagination files
rm -f "${REPO_ROOT}/pagination_"*.go
go run ./internal/codegen/pagination
echo "==> Pagination wrappers generated."

echo ""
echo "==> Generating webhooks..."
go run ./internal/codegen/webhooks --spec "${SPECS_DIR}/notifications/openapi.json" --out "${REPO_ROOT}/webhooks_gen/notifications_events.go"
echo "==> Webhooks generated."

echo ""
echo "==> Running goimports on generated code..."
goimports_bin=""
if command -v goimports &> /dev/null; then
    goimports_bin="$(command -v goimports)"
else
    echo "goimports not found, installing..."
    go install golang.org/x/tools/cmd/goimports@latest
    gobin="$(go env GOBIN)"
    if [ -z "${gobin}" ]; then
        gopath="$(go env GOPATH)"
        gobin="${gopath}/bin"
    fi
    if [ -n "${gobin}" ] && [ -x "${gobin}/goimports" ]; then
        goimports_bin="${gobin}/goimports"
    fi
fi

if [ -n "${goimports_bin}" ]; then
    find "${GEN_DIR}" -name "*.go" -exec "${goimports_bin}" -w -local "${MODULE_NAME}" {} \;
    find "${REPO_ROOT}/webhooks_gen" -name "*.go" -exec "${goimports_bin}" -w -local "${MODULE_NAME}" {} \;
    find "${REPO_ROOT}/pagination" -name "*.go" -exec "${goimports_bin}" -w -local "${MODULE_NAME}" {} \;
else
    echo "Warning: goimports still not available, skipping import formatting"
fi

echo ""
echo "==> Generation complete! Generated code is in: ${GEN_DIR}"
