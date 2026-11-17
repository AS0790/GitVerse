#!/usr/bin/env bash
set -euo pipefail
echo "Running go vet..."
go vet ./...
echo "Running tests..."
go test ./...
echo "Running gendocs check..."
go run ./cmd/gendocs -check
echo "Running JavaScript unit tests..."
node --test "internal/webui/jstest/*.test.mjs"
echo "All checks passed successfully!"
