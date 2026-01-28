#!/bin/bash
set -e

echo "Running SDK Reference Consumer Validation..."
cd examples/reference_consumer
go run main.go
echo "SDK Reference Consumer Validation Passed."
