#!/bin/bash
set -e

# Default path to consumer sibling directory, can be overridden by argument
CONSUMER_DIR="${1:-../tayvens-stock-report}"
SDK_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SDK_MODULE="github.com/Chalupa-Tech/go-schwab-api-individual"

echo "=================================================="
echo "SDK Validation: Consumer Integration Test"
echo "=================================================="
echo "SDK Directory:      ${SDK_DIR}"
echo "Consumer Directory: ${CONSUMER_DIR}"
echo "=================================================="

# Verify consumer directory exists
if [ ! -d "${CONSUMER_DIR}" ]; then
    echo "❌ Error: Consumer directory not found at ${CONSUMER_DIR}"
    echo "Usage: ./scripts/validate_consumer.sh [path_to_consumer_repo]"
    exit 1
fi

# Normalize Consumer Dir path
CONSUMER_ABS_PATH="$(cd "${CONSUMER_DIR}" && pwd)"

echo "📍 Switching to consumer directory: ${CONSUMER_ABS_PATH}"
cd "${CONSUMER_ABS_PATH}"

echo "🔧 Replacing SDK module with local version..."
# Force replace to ensure we are testing against the current SDK state
go mod edit -replace "${SDK_MODULE}=${SDK_DIR}"

echo "🧹 Tidying modules..."
# This ensures dependencies are consistent with the replaced SDK
go mod tidy

echo "🏗️  Running build verification..."
if go build -mod=mod -v ./...; then
    echo "✅ Build SUCCESSFUL."
else
    echo "❌ Build FAILED."
    exit 1
fi

echo "=================================================="
echo "Validation Complete."
